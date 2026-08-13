package tools

import (
	"context"
	_ "embed"
	"fmt"
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
			expr := fmt.Sprintf(`
import json, sys
from better_code_review_graph.tools import build_or_update_graph
r = build_or_update_graph(full_rebuild=%t, repo_root=%q, base=%q)
sys.stdout.write(json.dumps(r))
`, params.FullRebuild, root, base)
			out, err := runBCRG(expr)
			if err != nil {
				return fantasy.NewTextResponse(out + "\n" + err.Error()), nil
			}
			if strings.TrimSpace(out) == "" {
				return fantasy.NewTextResponse("graph built (no output)"), nil
			}
			return fantasy.NewTextResponse(out), nil
		},
	)
}
