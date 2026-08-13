package tools

import (
	"bytes"
	"context"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"charm.land/fantasy"
)

// CodegraphToolName exposes the native CodeGraph CLI knowledge-graph index.
const CodegraphToolName = "codegraph"

//go:embed codegraph.md
var codegraphDescription string

// codegraphBin returns the path to the CodeGraph CLI, overridable via
// CRUSH_CODEGRAPH_BIN. Defaults to the Homebrew install.
func codegraphBin() string {
	if p := os.Getenv("CRUSH_CODEGRAPH_BIN"); p != "" {
		return p
	}
	return "/opt/homebrew/bin/codegraph"
}

// codegraphTimeout bounds a CodeGraph CLI run; overridable via
// CRUSH_CODEGRAPH_TIMEOUT (seconds).
func codegraphTimeout() time.Duration {
	if v := os.Getenv("CRUSH_CODEGRAPH_TIMEOUT"); v != "" {
		if secs, err := strconv.Atoi(v); err == nil && secs > 0 {
			return time.Duration(secs) * time.Second
		}
	}
	return defaultBCRGTimeout
}

// runCodegraph runs the CodeGraph CLI and returns its combined output, capped
// to keep memory bounded.
func runCodegraph(ctx context.Context, args ...string) (string, error) {
	bin := codegraphBin()
	if bin == "" {
		return "", errors.New("cannot resolve codegraph binary")
	}
	cmdCtx, cancel := context.WithTimeout(ctx, codegraphTimeout())
	defer cancel()
	cmd := exec.CommandContext(cmdCtx, bin, args...)
	var out bytes.Buffer
	out.Grow(4096)
	cmd.Stdout = &boundedWriter{w: &out, limit: maxBCRGOutput}
	cmd.Stderr = &boundedWriter{w: &out, limit: maxBCRGOutput}
	if err := cmd.Run(); err != nil {
		return out.String(), fmt.Errorf("codegraph: %w", err)
	}
	return out.String(), nil
}

// codegraphOps are the CLI subcommands the tool accepts.
var codegraphOps = []string{
	"query", "callers", "callees", "context", "files",
	"impact", "affected", "index", "sync", "status",
}

type codegraphParams struct {
	Op     string   `json:"op,omitempty" description:"CodeGraph CLI subcommand"`
	Path   string   `json:"path,omitempty" description:"Project path (defaults to working dir)"`
	Query  string   `json:"query,omitempty" description:"Search term for query/context"`
	Symbol string   `json:"symbol,omitempty" description:"Symbol for callers/callees/impact"`
	Limit  int      `json:"limit,omitempty" description:"Max results"`
	Files  []string `json:"files,omitempty" description:"Changed files for affected"`
}

// buildCodegraphArgs maps op params to CLI argv. All user values are passed
// as separate argv entries (never interpolated into a shell).
func buildCodegraphArgs(op string, p codegraphParams) []string {
	args := []string{op}
	switch op {
	case "query", "callers", "callees", "impact":
		target := p.Symbol
		if op == "query" {
			target = p.Query
		}
		if target != "" {
			args = append(args, target)
		}
		if p.Limit > 0 {
			args = append(args, "-l", strconv.Itoa(p.Limit))
		}
	case "context":
		if p.Query != "" {
			args = append(args, p.Query)
		}
	case "affected":
		args = append(args, p.Files...)
	}
	if p.Path != "" {
		args = append(args, "-p", p.Path)
	}
	switch op {
	case "query", "callers", "callees", "impact", "affected", "files":
		args = append(args, "-j")
	}
	return args
}

func NewCodegraphTool(workingDir string) fantasy.AgentTool {
	return fantasy.NewAgentTool(
		CodegraphToolName,
		codegraphDescription,
		func(ctx context.Context, params codegraphParams, _ fantasy.ToolCall) (fantasy.ToolResponse, error) {
			if params.Op == "" {
				return fantasy.NewTextErrorResponse("op is required"), nil
			}
			valid := false
			for _, op := range codegraphOps {
				if params.Op == op {
					valid = true
					break
				}
			}
			if !valid {
				return fantasy.NewTextErrorResponse(fmt.Sprintf("unknown op %q", params.Op)), nil
			}
			repo := params.Path
			if repo == "" {
				repo = workingDir
			}
			params.Path = repo
			args := buildCodegraphArgs(params.Op, params)
			out, err := runCodegraph(ctx, args...)
			if err != nil {
				return fantasy.NewTextErrorResponse(out + "\n" + err.Error()), nil
			}
			return fantasy.NewTextResponse(out), nil
		},
	)
}
