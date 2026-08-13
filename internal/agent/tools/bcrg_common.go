package tools

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

// defaultBCRGTimeout bounds how long a bcrg subprocess may run before it is
// killed, so a hung graph build cannot block a tool call forever. Overridable
// via CRUSH_BCRG_TIMEOUT (in seconds).
const defaultBCRGTimeout = 5 * time.Minute

// maxBCRGOutput caps how much combined stdout+stderr a bcrg subprocess may
// produce, so a chatty build cannot grow memory without bound.
const maxBCRGOutput = 4 << 20 // 4 MiB

// bcrgVenvPython holds the path to the better-code-review-graph venv
// interpreter. Set via the CRUSH_BCRG_PYTHON env var to override it; it
// defaults to the user's data/oss checkout.
func bcrgPythonPath() string {
	if p := os.Getenv("CRUSH_BCRG_PYTHON"); p != "" {
		return p
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, "data", "oss", "better-code-review-graph", ".venv", "bin", "python")
}

// bcrgTimeout returns the per-call subprocess timeout, honoring
// CRUSH_BCRG_TIMEOUT (seconds) or falling back to the default.
func bcrgTimeout() time.Duration {
	if v := os.Getenv("CRUSH_BCRG_TIMEOUT"); v != "" {
		if secs, err := strconv.Atoi(v); err == nil && secs > 0 {
			return time.Duration(secs) * time.Second
		}
	}
	return defaultBCRGTimeout
}

// runBCRG runs a Python expression in the bcrg venv and returns combined
// stdout+stderr. The expression is fixed Crush-side; user-supplied arguments
// are passed as separate argv entries (not interpolated into the script) so
// they cannot break out of the -c string or inject code. The subprocess is
// killed if it exceeds the configured timeout or the caller's context is
// cancelled; its output is capped to avoid unbounded buffering.
func runBCRG(ctx context.Context, expr string, args ...string) (string, error) {
	py := bcrgPythonPath()
	if py == "" {
		return "", errors.New("cannot resolve bcrg python interpreter")
	}
	cmdCtx, cancel := context.WithTimeout(ctx, bcrgTimeout())
	defer cancel()
	cmd := exec.CommandContext(cmdCtx, py, append([]string{"-c", expr}, args...)...)

	var out bytes.Buffer
	out.Grow(4096)
	cmd.Stdout = &boundedWriter{w: &out, limit: maxBCRGOutput}
	cmd.Stderr = &boundedWriter{w: &out, limit: maxBCRGOutput}
	if err := cmd.Run(); err != nil {
		return out.String(), fmt.Errorf("bcrg: %w", err)
	}
	return out.String(), nil
}

// boundedWriter discards writes beyond limit so a chatty process cannot grow
// memory without bound.
type boundedWriter struct {
	w     *bytes.Buffer
	limit int
}

func (b *boundedWriter) Write(p []byte) (int, error) {
	remaining := b.limit - b.w.Len()
	if remaining <= 0 {
		return len(p), nil
	}
	if len(p) > remaining {
		p = p[:remaining]
	}
	n, _ := b.w.Write(p)
	return n, nil
}
