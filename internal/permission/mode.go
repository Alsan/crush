package permission

// Mode describes the permission approval mode.
//
// The mode is a three-state machine: off, local, and global.
//
//   - ModeOff: every tool request goes through the normal confirmation flow.
//   - ModeLocal: requests whose target path lies inside the working directory
//     subtree are auto-approved. Pathless requests (e.g. bash) and requests
//     outside the subtree still prompt. This is the default mode.
//   - ModeGlobal: all requests are auto-approved, regardless of path. This is
//     the historical "yolo" mode (--yolo flag / Ctrl+Y).
//
// ModeGlobal supersedes ModeLocal: when global mode is active the local flag
// may still be set, but the service short-circuits on the global check first.
type Mode int

const (
	// ModeOff disables auto-approval entirely.
	ModeOff Mode = iota
	// ModeLocal auto-approves requests inside the working directory subtree.
	ModeLocal
	// ModeGlobal auto-approves every request.
	ModeGlobal
)

// String returns a stable lowercase name for the mode.
func (m Mode) String() string {
	switch m {
	case ModeLocal:
		return "local"
	case ModeGlobal:
		return "global"
	default:
		return "off"
	}
}

// ModeFromFlags derives the effective mode from the global and local skip
// flags. Global wins; otherwise local; otherwise off.
func ModeFromFlags(global, local bool) Mode {
	switch {
	case global:
		return ModeGlobal
	case local:
		return ModeLocal
	default:
		return ModeOff
	}
}

// Flags returns the global and local skip flags that encode the mode.
// ModeGlobal sets both so the mode survives a round trip through the
// two-boolean wire format and a later mode computation still yields Global.
func (m Mode) Flags() (global, local bool) {
	switch m {
	case ModeGlobal:
		return true, true
	case ModeLocal:
		return false, true
	default:
		return false, false
	}
}
