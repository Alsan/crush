package permission

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/charmbracelet/crush/internal/pubsub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestPermissionService_LocalYoloAutoApprovesWithinSubtree verifies that
// local yolo mode auto-grants requests whose target path resolves inside the
// working directory subtree, while pathless requests and requests outside
// the subtree still require confirmation.
func TestPermissionService_LocalYoloAutoApprovesWithinSubtree(t *testing.T) {
	t.Parallel()

	workDir := t.TempDir()
	subDir := filepath.Join(workDir, "sub")
	require.NoError(t, os.MkdirAll(subDir, 0o755))
	outside := t.TempDir()

	t.Run("file inside subtree is auto-approved", func(t *testing.T) {
		t.Parallel()
		service := NewPermissionServiceWithExclusions(workDir, false, nil, []string{})
		service.SetLocalSkipRequests(true)

		granted, err := service.Request(t.Context(), CreatePermissionRequest{
			SessionID:  "s1",
			ToolCallID: "call-1",
			ToolName:   "edit",
			Action:     "write",
			Path:       filepath.Join(subDir, "file.go"),
		})
		require.NoError(t, err)
		assert.True(t, granted)
	})

	t.Run("working directory itself is auto-approved", func(t *testing.T) {
		t.Parallel()
		service := NewPermissionServiceWithExclusions(workDir, false, nil, []string{})
		service.SetLocalSkipRequests(true)

		granted, err := service.Request(t.Context(), CreatePermissionRequest{
			SessionID:  "s2",
			ToolCallID: "call-2",
			ToolName:   "edit",
			Action:     "write",
			Path:       workDir,
		})
		require.NoError(t, err)
		assert.True(t, granted)
	})

	t.Run("path escaping via .. is not auto-approved", func(t *testing.T) {
		t.Parallel()
		service := NewPermissionServiceWithExclusions(workDir, false, nil, []string{})
		service.SetLocalSkipRequests(true)

		granted := requestAndDeny(t, service, "call-4", "edit", "write", filepath.Join(workDir, "..", "escape.go"))
		assert.False(t, granted, "path escaping the subtree must still prompt")
	})

	t.Run("path outside subtree is not auto-approved", func(t *testing.T) {
		t.Parallel()
		service := NewPermissionServiceWithExclusions(workDir, false, nil, []string{})
		service.SetLocalSkipRequests(true)

		granted := requestAndDeny(t, service, "call-5", "edit", "write", filepath.Join(outside, "file.go"))
		assert.False(t, granted, "path outside the subtree must still prompt")
	})

	t.Run("pathless request is not auto-approved", func(t *testing.T) {
		t.Parallel()
		service := NewPermissionServiceWithExclusions(workDir, false, nil, []string{})
		service.SetLocalSkipRequests(true)

		granted := requestAndDeny(t, service, "call-6", "bash", "execute", "")
		assert.False(t, granted, "pathless requests (e.g. bash) must still prompt")
	})

	t.Run("symlink pointing outside subtree is not auto-approved", func(t *testing.T) {
		if !symlinksSupported(t) {
			t.Skip("symlinks not supported on this platform")
		}
		t.Parallel()
		service := NewPermissionServiceWithExclusions(workDir, false, nil, []string{})
		service.SetLocalSkipRequests(true)

		link := filepath.Join(workDir, "escape-link")
		require.NoError(t, os.Symlink(outside, link))

		granted := requestAndDeny(t, service, "call-7", "edit", "write", filepath.Join(link, "file.go"))
		assert.False(t, granted, "a symlink resolving outside the subtree must still prompt")
	})
}

// TestPermissionService_LocalYoloDisabledPrompts verifies that with local
// yolo mode off, even in-subtree requests require confirmation.
func TestPermissionService_LocalYoloDisabledPrompts(t *testing.T) {
	t.Parallel()

	workDir := t.TempDir()
	service := NewPermissionService(workDir, false, nil)

	granted := requestAndDeny(t, service, "call-1", "edit", "write", filepath.Join(workDir, "file.go"))
	assert.False(t, granted, "without local yolo the request must prompt")
}

// TestPermissionService_LocalYoloPublishesGrantedNotification verifies that
// a local-yolo auto-grant publishes a Granted notification so the UI and
// audit subscribers see the outcome.
func TestPermissionService_LocalYoloPublishesGrantedNotification(t *testing.T) {
	t.Parallel()

	workDir := t.TempDir()
	service := NewPermissionServiceWithExclusions(workDir, false, nil, []string{})
	service.SetLocalSkipRequests(true)

	notifs := service.SubscribeNotifications(t.Context())
	granted, err := service.Request(t.Context(), CreatePermissionRequest{
		SessionID:  "s1",
		ToolCallID: "call-1",
		ToolName:   "edit",
		Action:     "write",
		Path:       filepath.Join(workDir, "file.go"),
	})
	require.NoError(t, err)
	assert.True(t, granted)

	select {
	case ev := <-notifs:
		assert.Equal(t, pubsub.CreatedEvent, ev.Type)
		assert.True(t, ev.Payload.Granted)
		assert.False(t, ev.Payload.Denied)
		assert.Equal(t, "call-1", ev.Payload.ToolCallID)
	default:
		t.Fatal("expected a granted notification from the local yolo auto-grant")
	}
}

// TestPermissionService_LocalYoloExtraPathsAutoApproved verifies that local
// yolo mode auto-approves requests under a configured extra root even when it
// sits outside the working directory subtree, so trusted directories like
// ~/.config/crush do not prompt.
func TestPermissionService_LocalYoloExtraPathsAutoApproved(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	// Workdir lives in a separate temp dir so the extra root (sibling of
	// HOME, outside the subtree) is only approved because it is listed.
	workDir := t.TempDir()
	extra := filepath.Join(home, ".config", "crush")
	require.NoError(t, os.MkdirAll(extra, 0o755))

	service := NewPermissionServiceWithExclusions(workDir, false, nil, []string{extra})
	service.SetLocalSkipRequests(true)

	for _, rel := range []string{
		".config/crush",
		".config/crush/AGENTS.md",
		".config/crush/agent.json",
	} {
		rel := rel
		t.Run(rel, func(t *testing.T) {
			granted, err := service.Request(t.Context(), CreatePermissionRequest{
				SessionID:  "s1",
				ToolCallID: "call-x",
				ToolName:   "edit",
				Action:     "write",
				Path:       filepath.Join(home, rel),
			})
			require.NoError(t, err)
			assert.True(t, granted, "listed extra root %q must be auto-approved", rel)
		})
	}
}

// TestPermissionService_LocalYoloNoExtraConfigPrompts verifies that without
// a configured extra list, local yolo auto-approves only paths inside the
// working directory subtree; an out-of-subtree path still prompts.
func TestPermissionService_LocalYoloNoExtraConfigPrompts(t *testing.T) {
	t.Parallel()

	workDir := t.TempDir()
	outside := t.TempDir()

	service := NewPermissionService(workDir, false, nil)
	service.SetLocalSkipRequests(true)

	granted := requestAndDeny(t, service, "call-1", "edit", "write", filepath.Join(outside, "file.go"))
	assert.False(t, granted, "an unlisted out-of-subtree path must still prompt")
}

// TestPermissionService_LocalYoloConfigExtraRoots verifies that a configured
// extra root is auto-approved even outside the subtree, while a path beneath
// the working directory that is not covered by an extra root stays approved
// via the subtree and a path outside both still prompts.
func TestPermissionService_LocalYoloConfigExtraRoots(t *testing.T) {
	t.Parallel()

	workDir := t.TempDir()
	extra := t.TempDir()
	outside := t.TempDir()

	service := NewPermissionServiceWithExclusions(workDir, false, nil, []string{extra})
	service.SetLocalSkipRequests(true)

	// A path inside the configured extra root is auto-approved.
	granted, err := service.Request(t.Context(), CreatePermissionRequest{
		SessionID:  "s1",
		ToolCallID: "call-1",
		ToolName:   "edit",
		Action:     "write",
		Path:       filepath.Join(extra, "cfg.json"),
	})
	require.NoError(t, err)
	assert.True(t, granted, "a path inside a configured extra root must be auto-approved")

	// A path outside both the subtree and any extra root still prompts.
	granted = requestAndDeny(t, service, "call-2", "edit", "write", filepath.Join(outside, "file.go"))
	assert.False(t, granted, "a path outside the subtree and extra roots must still prompt")
}

// TestPermissionService_GlobalYoloSupersedesLocal verifies that global yolo
// (SkipRequests) auto-approves everything, including pathless requests and
// out-of-subtree paths, even when local yolo is on.
func TestPermissionService_GlobalYoloSupersedesLocal(t *testing.T) {
	t.Parallel()

	outside := t.TempDir()
	service := NewPermissionService(t.TempDir(), true, nil)
	service.SetLocalSkipRequests(true)

	for _, tt := range []struct {
		name string
		path string
	}{
		{name: "pathless", path: ""},
		{name: "outside", path: outside},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			granted, err := service.Request(t.Context(), CreatePermissionRequest{
				SessionID:  "s1",
				ToolCallID: "call-1",
				ToolName:   "bash",
				Action:     "execute",
				Path:       tt.path,
			})
			require.NoError(t, err)
			assert.True(t, granted, "global yolo must auto-approve regardless of path")
		})
	}
}

// requestAndDeny issues a permission request that falls through to the
// prompt, consumes the pending request event, and denies it, returning the
// outcome the caller received. It fails the test if the request was
// auto-approved (no event) or the flow misbehaved.
func requestAndDeny(t *testing.T, service Service, toolCallID, tool, action, path string) bool {
	t.Helper()
	events := service.Subscribe(t.Context())

	type outcome struct {
		granted bool
		err     error
	}
	resCh := make(chan outcome, 1)
	go func() {
		granted, err := service.Request(t.Context(), CreatePermissionRequest{
			SessionID:  "s",
			ToolCallID: toolCallID,
			ToolName:   tool,
			Action:     action,
			Path:       path,
		})
		resCh <- outcome{granted: granted, err: err}
	}()

	select {
	case ev := <-events:
		service.Deny(ev.Payload)
	case <-time.After(5 * time.Second):
		t.Fatal("expected the request to reach the confirmation flow (pending event)")
	}

	select {
	case res := <-resCh:
		require.NoError(t, res.err)
		return res.granted
	case <-time.After(5 * time.Second):
		t.Fatal("request did not resolve after deny")
		return false
	}
}

// symlinksSupported reports whether os.Symlink works on this platform.
func symlinksSupported(t *testing.T) bool {
	t.Helper()
	dir := t.TempDir()
	link := filepath.Join(dir, "link")
	if err := os.Symlink(dir, link); err != nil {
		return false
	}
	return true
}
