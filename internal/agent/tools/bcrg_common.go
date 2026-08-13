package tools

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// bcrgPython returns the Python interpreter inside the better-code-review-graph
// source venv used by Crush's native integration. The repo lives under the
// user's data/oss directory alongside Crush (personal setup).
func bcrgPython() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, "data", "oss", "better-code-review-graph", ".venv", "bin", "python")
}

// runBCRG runs a Python expression via the bcrg venv interpreter and returns
// its combined stdout/stderr. The expr must print the JSON result on stdout.
func runBCRG(expr string) (string, error) {
	py := bcrgPython()
	if py == "" {
		return "", fmt.Errorf("cannot resolve bcrg python interpreter")
	}
	out, err := exec.Command(py, "-c", expr).CombinedOutput()
	if err != nil {
		return string(out), fmt.Errorf("bcrg: %w", err)
	}
	return string(out), nil
}
