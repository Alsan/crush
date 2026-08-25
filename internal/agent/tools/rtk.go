package tools

import (
	"log/slog"
	"os/exec"
	"sync"
	"testing"

	"github.com/charmbracelet/crush/internal/log"
)

// getRtk returns the path to the rtk binary, or "" when not available.
// rtk filters and compresses command output to save tokens.
var getRtk = sync.OnceValue(func() string {
	if testing.Testing() {
		return ""
	}
	path, err := exec.LookPath("rtk")
	if err != nil {
		if log.Initialized() {
			slog.Warn("rtk not found in $PATH. Command output won't be compressed.")
		}
		return ""
	}
	return path
})