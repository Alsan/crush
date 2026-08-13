package permission

import (
	"context"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/charmbracelet/crush/internal/csync"
	"github.com/charmbracelet/crush/internal/pubsub"
	"github.com/google/uuid"
)

// hookApprovalKey is the unexported context key used to mark a tool call as
// pre-approved by a PreToolUse hook. The value is the tool call ID so an
// approval can't be reused across calls that happen to share a context.
type hookApprovalKey struct{}

// WithHookApproval returns a context that marks the given tool call ID as
// pre-approved by a hook. When the permission service sees a matching
// request it short-circuits the normal prompt and grants immediately.
func WithHookApproval(ctx context.Context, toolCallID string) context.Context {
	return context.WithValue(ctx, hookApprovalKey{}, toolCallID)
}

// hookApproved reports whether the context carries a hook approval for the
// given tool call ID.
func hookApproved(ctx context.Context, toolCallID string) bool {
	if toolCallID == "" {
		return false
	}
	v, _ := ctx.Value(hookApprovalKey{}).(string)
	return v == toolCallID
}

type CreatePermissionRequest struct {
	SessionID   string `json:"session_id"`
	ToolCallID  string `json:"tool_call_id"`
	ToolName    string `json:"tool_name"`
	Description string `json:"description"`
	Action      string `json:"action"`
	Params      any    `json:"params"`
	Path        string `json:"path"`
}

type PermissionNotification struct {
	ToolCallID string `json:"tool_call_id"`
	Granted    bool   `json:"granted"`
	Denied     bool   `json:"denied"`
}

type PermissionRequest struct {
	ID          string `json:"id"`
	SessionID   string `json:"session_id"`
	ToolCallID  string `json:"tool_call_id"`
	ToolName    string `json:"tool_name"`
	Description string `json:"description"`
	Action      string `json:"action"`
	Params      any    `json:"params"`
	Path        string `json:"path"`
}

type Service interface {
	pubsub.Subscriber[PermissionRequest]
	// GrantPersistent grants a permission request and remembers the grant
	// for the session. It returns true if this call actually resolved the
	// pending request; false if the request had already been resolved
	// (e.g., by another concurrent caller) or is unknown.
	GrantPersistent(permission PermissionRequest) bool
	// Grant grants a permission request. It returns true if this call
	// actually resolved the pending request; false if the request had
	// already been resolved or is unknown.
	Grant(permission PermissionRequest) bool
	// Deny denies a permission request. It returns true if this call
	// actually resolved the pending request; false if the request had
	// already been resolved or is unknown.
	Deny(permission PermissionRequest) bool
	Request(ctx context.Context, opts CreatePermissionRequest) (bool, error)
	AutoApproveSession(sessionID string)
	SetSkipRequests(skip bool)
	SkipRequests() bool
	// SetLocalSkipRequests enables or disables local yolo mode: requests
	// whose target path lies inside the working directory subtree are
	// auto-approved.
	SetLocalSkipRequests(local bool)
	// LocalSkipRequests reports whether local yolo mode is enabled.
	LocalSkipRequests() bool
	SubscribeNotifications(ctx context.Context) <-chan pubsub.Event[PermissionNotification]
}

// PermissionKey is a composite key for session permission lookups.
type PermissionKey struct {
	SessionID string
	ToolName  string
	Action    string
	Path      string
}

type permissionService struct {
	*pubsub.Broker[PermissionRequest]

	notificationBroker    *pubsub.Broker[PermissionNotification]
	workingDir            string
	sessionPermissions    *csync.Map[PermissionKey, bool]
	pendingRequests       *csync.Map[string, chan bool]
	autoApproveSessions   map[string]bool
	autoApproveSessionsMu sync.RWMutex
	skip                  atomic.Bool
	localSkip             atomic.Bool
	allowedTools          []string
	// excludedPaths are absolute, symlink-resolved directories that local
	// yolo mode never auto-approves, even when they sit inside the working
	// directory subtree. Defaults are seeded from the built-in list; callers
	// may override/extend them via config.
	excludedPaths []string

	// used to make sure we only process one request at a time
	requestMu       sync.Mutex
	activeRequest   *PermissionRequest
	activeRequestMu sync.Mutex
}

// resolve atomically removes the pending request entry for the given
// permission and, if it was still pending, publishes exactly one
// PermissionNotification and forwards the outcome to the waiter on
// respCh. It returns true if this call resolved the request, false if
// it had already been resolved (e.g., by another concurrent caller) or
// the request ID is unknown.
//
// If onResolve is non-nil it runs after the pending entry has been
// taken but before the notification is published or the waiter is
// unblocked. This lets GrantPersistent record the session permission
// only when it actually wins the race, so a losing GrantPersistent
// that lost to a Deny does not leak an auto-approve entry.
//
// All three public resolution methods (Grant, GrantPersistent, Deny)
// route through this helper so multi-subscriber UIs can race safely:
// the first caller wins, the rest become no-ops.
func (s *permissionService) resolve(permission PermissionRequest, granted, denied bool, onResolve func()) bool {
	respCh, ok := s.pendingRequests.Take(permission.ID)
	if !ok {
		return false
	}

	if onResolve != nil {
		onResolve()
	}

	s.notificationBroker.Publish(pubsub.CreatedEvent, PermissionNotification{
		ToolCallID: permission.ToolCallID,
		Granted:    granted,
		Denied:     denied,
	})

	// respCh is buffered (cap 1) and only ever has at most one sender
	// per request because Take removes the entry under the map lock,
	// so this send never blocks.
	respCh <- granted

	s.activeRequestMu.Lock()
	if s.activeRequest != nil && s.activeRequest.ID == permission.ID {
		s.activeRequest = nil
	}
	s.activeRequestMu.Unlock()
	return true
}

func (s *permissionService) GrantPersistent(permission PermissionRequest) bool {
	// Record the persistent grant only if this call wins the
	// pending-request race. Otherwise a losing GrantPersistent that
	// lost to a Deny would still leave an auto-approve entry behind,
	// silently flipping later denied calls to allowed.
	return s.resolve(permission, true, false, func() {
		s.sessionPermissions.Set(PermissionKey{
			SessionID: permission.SessionID,
			ToolName:  permission.ToolName,
			Action:    permission.Action,
			Path:      permission.Path,
		}, true)
	})
}

func (s *permissionService) Grant(permission PermissionRequest) bool {
	return s.resolve(permission, true, false, nil)
}

func (s *permissionService) Deny(permission PermissionRequest) bool {
	return s.resolve(permission, false, true, nil)
}

func (s *permissionService) Request(ctx context.Context, opts CreatePermissionRequest) (bool, error) {
	if s.skip.Load() {
		return true, nil
	}

	// Check if the tool/action combination is in the allowlist
	commandKey := opts.ToolName + ":" + opts.Action
	if slices.Contains(s.allowedTools, commandKey) || slices.Contains(s.allowedTools, opts.ToolName) {
		return true, nil
	}

	// Local yolo mode: auto-approve requests whose target path resolves
	// inside the working directory subtree. Pathless requests (e.g. bash),
	// requests outside the subtree, and requests under an excluded path
	// fall through to the normal flow. The absolute path is resolved once
	// and shared by the subtree and exclusion checks so the (syscall-heavy)
	// symlink walk does not run twice per request. A granted notification
	// is published so the UI and audit subscribers see the outcome,
	// matching the other auto-approval paths.
	if s.localSkip.Load() {
		if abs := s.resolvePath(opts.Path); abs != "" && s.withinSubtree(abs) && !s.withinExcluded(abs) {
			s.notificationBroker.Publish(pubsub.CreatedEvent, PermissionNotification{
				ToolCallID: opts.ToolCallID,
				Granted:    true,
			})
			return true, nil
		}
	}

	// A PreToolUse hook that returned decision=allow stamps the context
	// with the tool call ID. Treat that as a pre-approval and skip the
	// prompt entirely. We still publish a granted notification so the UI
	// and audit subscribers see the outcome.
	if hookApproved(ctx, opts.ToolCallID) {
		s.notificationBroker.Publish(pubsub.CreatedEvent, PermissionNotification{
			ToolCallID: opts.ToolCallID,
			Granted:    true,
		})
		return true, nil
	}

	s.requestMu.Lock()
	defer s.requestMu.Unlock()

	// tell the UI that a permission was requested
	s.notificationBroker.Publish(pubsub.CreatedEvent, PermissionNotification{
		ToolCallID: opts.ToolCallID,
	})

	s.autoApproveSessionsMu.RLock()
	autoApprove := s.autoApproveSessions[opts.SessionID]
	s.autoApproveSessionsMu.RUnlock()

	if autoApprove {
		s.notificationBroker.Publish(pubsub.CreatedEvent, PermissionNotification{
			ToolCallID: opts.ToolCallID,
			Granted:    true,
		})
		return true, nil
	}

	fileInfo, err := os.Stat(opts.Path)
	dir := opts.Path
	if err == nil {
		if fileInfo.IsDir() {
			dir = opts.Path
		} else {
			dir = filepath.Dir(opts.Path)
		}
	}

	if dir == "." {
		dir = s.workingDir
	}
	permission := PermissionRequest{
		ID:          uuid.New().String(),
		Path:        dir,
		SessionID:   opts.SessionID,
		ToolCallID:  opts.ToolCallID,
		ToolName:    opts.ToolName,
		Description: opts.Description,
		Action:      opts.Action,
		Params:      opts.Params,
	}

	if _, ok := s.sessionPermissions.Get(PermissionKey{
		SessionID: permission.SessionID,
		ToolName:  permission.ToolName,
		Action:    permission.Action,
		Path:      permission.Path,
	}); ok {
		s.notificationBroker.Publish(pubsub.CreatedEvent, PermissionNotification{
			ToolCallID: opts.ToolCallID,
			Granted:    true,
		})
		return true, nil
	}

	s.activeRequestMu.Lock()
	s.activeRequest = &permission
	s.activeRequestMu.Unlock()

	respCh := make(chan bool, 1)
	s.pendingRequests.Set(permission.ID, respCh)
	defer s.pendingRequests.Del(permission.ID)

	// Publish the request
	s.Publish(pubsub.CreatedEvent, permission)

	select {
	case <-ctx.Done():
		return false, ctx.Err()
	case granted := <-respCh:
		return granted, nil
	}
}

func (s *permissionService) AutoApproveSession(sessionID string) {
	s.autoApproveSessionsMu.Lock()
	s.autoApproveSessions[sessionID] = true
	s.autoApproveSessionsMu.Unlock()
}

func (s *permissionService) SubscribeNotifications(ctx context.Context) <-chan pubsub.Event[PermissionNotification] {
	return s.notificationBroker.Subscribe(ctx)
}

func (s *permissionService) SetSkipRequests(skip bool) {
	s.skip.Store(skip)
}

func (s *permissionService) SkipRequests() bool {
	return s.skip.Load()
}

func (s *permissionService) SetLocalSkipRequests(local bool) {
	s.localSkip.Store(local)
}

func (s *permissionService) LocalSkipRequests() bool {
	return s.localSkip.Load()
}

// resolvePath returns the absolute, symlink-resolved form of path, or ""
// when the path is empty or cannot be resolved. Callers resolve once and pass
// the result to withinSubtree/withinExcluded so the syscall-heavy symlink
// walk runs at most once per request.
func (s *permissionService) resolvePath(path string) string {
	if path == "" {
		return ""
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return ""
	}
	return resolveExisting(abs)
}

// withinSubtree reports whether path (already absolute and symlink-resolved
// via resolvePath) refers to a location at or under the service's working
// directory. Symlinks are resolved along the deepest existing ancestor so a
// target that does not exist yet (e.g. a file about to be created) is still
// checked against the real directory hierarchy, and paths that escape the
// subtree via ".." are rejected.
func (s *permissionService) withinSubtree(abs string) bool {
	if abs == "" {
		return false
	}
	rel, err := filepath.Rel(s.workingDir, abs)
	if err != nil {
		return false
	}
	if rel == "." {
		return true
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return false
	}
	return true
}

// withinExcluded reports whether path (already absolute and symlink-resolved
// via resolvePath) resolves at or under one of the configured exclusion
// roots. Symlink resolution is applied the same way as withinSubtree so an
// excluded location cannot be aliased past the check.
func (s *permissionService) withinExcluded(abs string) bool {
	if abs == "" || len(s.excludedPaths) == 0 {
		return false
	}
	for _, excluded := range s.excludedPaths {
		rel, err := filepath.Rel(excluded, abs)
		if err != nil {
			continue
		}
		if rel == "." || (rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator))) {
			return true
		}
	}
	return false
}

// resolveExisting resolves symlinks for the deepest existing ancestor of
// path and re-appends the (possibly non-existent) remainder, so the result
// is an absolute path whose real on-disk prefix is symlink-free. It returns
// path unchanged when nothing on the path exists or symlink resolution
// fails.
func resolveExisting(path string) string {
	existing := path
	var suffix []string
	for {
		if _, err := os.Lstat(existing); err == nil {
			break
		}
		parent := filepath.Dir(existing)
		if parent == existing {
			return path
		}
		suffix = append([]string{filepath.Base(existing)}, suffix...)
		existing = parent
	}
	resolved, err := filepath.EvalSymlinks(existing)
	if err != nil {
		return path
	}
	parts := append([]string{resolved}, suffix...)
	return filepath.Join(parts...)
}

func NewPermissionService(workingDir string, skip bool, allowedTools []string) Service {
	return NewPermissionServiceWithExclusions(workingDir, skip, allowedTools, nil)
}

// NewPermissionServiceWithExclusions constructs a permission service with an
// explicit set of exclusion roots for local yolo mode. When exclusions is
// nil the built-in default set (~/.agents, ~/.claude, ~/go, ~/.config/crush,
// ~/.cache, /tmp) is used; a non-nil slice replaces it entirely.
func NewPermissionServiceWithExclusions(workingDir string, skip bool, allowedTools []string, exclusions []string) Service {
	// Normalize the working directory once at construction so subtree
	// checks are stable: resolve to an absolute, symlink-free path. On
	// macOS /var is a symlink to /private/var; without this a target
	// resolved via EvalSymlinks would never appear inside the raw
	// workingDir and every local-yolo check would conservatively fail.
	wd := workingDir
	if wd != "" {
		if abs, err := filepath.Abs(wd); err == nil {
			wd = abs
		}
		if resolved, err := filepath.EvalSymlinks(wd); err == nil {
			wd = resolved
		}
	}
	svc := &permissionService{
		Broker:              pubsub.NewBroker[PermissionRequest](),
		notificationBroker:  pubsub.NewBroker[PermissionNotification](),
		workingDir:          wd,
		sessionPermissions:  csync.NewMap[PermissionKey, bool](),
		autoApproveSessions: make(map[string]bool),
		allowedTools:        allowedTools,
		pendingRequests:     csync.NewMap[string, chan bool](),
		excludedPaths:       normalizeExclusions(exclusions),
	}
	svc.skip.Store(skip)
	return svc
}

// normalizeExclusions expands each exclusion root to an absolute,
// symlink-resolved directory. A nil input selects the built-in defaults.
func normalizeExclusions(exclusions []string) []string {
	if exclusions == nil {
		exclusions = defaultExcludedPaths()
	}
	var normalized []string
	for _, p := range exclusions {
		expanded, err := expandHome(p)
		if err != nil {
			continue
		}
		if abs, err := filepath.Abs(expanded); err == nil {
			expanded = abs
		}
		normalized = append(normalized, resolveExisting(expanded))
	}
	return normalized
}

// defaultExcludedPaths returns the absolute, symlink-resolved directories
// that local yolo mode never auto-approves. These guard user-owned
// configuration, tooling locations, and system temp space that a project
// should not be able to modify silently, even when the working directory
// sits inside one of them.
func defaultExcludedPaths() []string {
	return []string{
		"~/.agents",
		"~/.claude",
		"~/go",
		"~/.config/crush",
		"~/.cache",
		"/tmp",
	}
}

// expandHome expands a leading "~/" to the current user's home directory.
func expandHome(p string) (string, error) {
	if !strings.HasPrefix(p, "~/") {
		return p, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, strings.TrimPrefix(p, "~/")), nil
}
