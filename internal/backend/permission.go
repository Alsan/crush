package backend

import (
	"github.com/charmbracelet/crush/internal/permission"
	"github.com/charmbracelet/crush/internal/proto"
)

// GrantPermission grants, denies, or persistently grants a permission
// request. The returned bool reports whether this call resolved the
// pending request (true) or found it already resolved by a previous
// caller (false). A false return is not an error.
func (b *Backend) GrantPermission(workspaceID string, req proto.PermissionGrant) (bool, error) {
	ws, err := b.GetWorkspace(workspaceID)
	if err != nil {
		return false, err
	}

	perm := permission.PermissionRequest{
		ID:          req.Permission.ID,
		SessionID:   req.Permission.SessionID,
		ToolCallID:  req.Permission.ToolCallID,
		ToolName:    req.Permission.ToolName,
		Description: req.Permission.Description,
		Action:      req.Permission.Action,
		Params:      req.Permission.Params,
		Path:        req.Permission.Path,
	}

	switch req.Action {
	case proto.PermissionAllow:
		return ws.Permissions.Grant(perm), nil
	case proto.PermissionAllowForSession:
		return ws.Permissions.GrantPersistent(perm), nil
	case proto.PermissionDeny:
		return ws.Permissions.Deny(perm), nil
	default:
		return false, ErrInvalidPermissionAction
	}
}

// SetPermissionsSkip sets the permission approval mode flags for a
// workspace: Skip is the global yolo flag, Local enables local yolo mode.
func (b *Backend) SetPermissionsSkip(workspaceID string, req proto.PermissionSkipRequest) error {
	ws, err := b.GetWorkspace(workspaceID)
	if err != nil {
		return err
	}

	ws.Permissions.SetSkipRequests(req.Skip)
	ws.Permissions.SetLocalSkipRequests(req.Local)
	return nil
}

// GetPermissionsSkip returns the permission approval mode flags for a
// workspace.
func (b *Backend) GetPermissionsSkip(workspaceID string) (proto.PermissionSkipRequest, error) {
	ws, err := b.GetWorkspace(workspaceID)
	if err != nil {
		return proto.PermissionSkipRequest{}, err
	}

	return proto.PermissionSkipRequest{
		Skip:  ws.Permissions.SkipRequests(),
		Local: ws.Permissions.LocalSkipRequests(),
	}, nil
}
