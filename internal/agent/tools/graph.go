package tools

import (
	"context"
	_ "embed"
	"strings"

	"charm.land/fantasy"
)

const GraphToolName = "graph"

//go:embed graph.md
var graphDescription string

type GraphParams struct {
	RepoRoot    string `json:"repo_root,omitempty" description:"Repository root (auto-detected if omitted)"`
	FullRebuild bool   `json:"full_rebuild,omitempty" description:"Re-parse every file (default false = incremental)"`
	Base        string `json:"base,omitempty" description:"Git ref for incremental diff (default HEAD~1)"`
}

func NewGraphTool(workingDir string) fantasy.AgentTool {
	return fantasy.NewAgentTool(
		GraphToolName,
		graphDescription,
		func(ctx context.Context, params GraphParams, _ fantasy.ToolCall) (fantasy.ToolResponse, error) {
			root := params.RepoRoot
			if root == "" {
				root = workingDir
			}
			base := params.Base
			if base == "" {
				base = "HEAD~1"
			}
			// Repo root/base are passed as argv (sys.argv[1..3]) so a path
			// containing quotes or newlines cannot break out of the Python -c
			// string or inject code.
			expr := `
import json, sys
from better_code_review_graph.tools import build_or_update_graph
r = build_or_update_graph(full_rebuild=sys.argv[1] == "1", repo_root=sys.argv[2], base=sys.argv[3])
sys.stdout.write(json.dumps(r))
`
			full := "0"
			if params.FullRebuild {
				full = "1"
			}
			out, err := runBCRG(ctx, expr, full, root, base)
			if err != nil {
				return fantasy.NewTextErrorResponse(out + "\n" + err.Error()), nil
			}
			if strings.TrimSpace(out) == "" {
				return fantasy.NewTextResponse("graph built (no output)"), nil
			}
			return fantasy.NewTextResponse(out), nil
		},
	)
}
