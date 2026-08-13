package tools

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

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

// runBCRG runs a Python expression in the bcrg venv and returns combined
// stdout+stderr. The expression is fixed Crush-side; user-supplied arguments
// are passed as separate argv entries (not interpolated into the script) so
// they cannot break out of the -c string or inject code.
func runBCRG(ctx context.Context, expr string, args ...string) (string, error) {
	py := bcrgPythonPath()
	if py == "" {
		return "", errors.New("cannot resolve bcrg python interpreter")
	}
	cmd := exec.CommandContext(ctx, py, append([]string{"-c", expr}, args...)...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), fmt.Errorf("bcrg: %w", err)
	}
	return string(out), nil
}
