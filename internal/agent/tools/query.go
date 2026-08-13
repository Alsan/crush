package tools

import (
	"context"
	_ "embed"

	"charm.land/fantasy"
)

const QueryToolName = "query"

//go:embed query.md
var queryDescription string

type QueryParams struct {
	Pattern  string `json:"pattern" description:"callers_of, callees_of, imports_of, importers_of, children_of, tests_for, inheritors_of, file_summary"`
	Target   string `json:"target" description:"Node name, qualified name, or file path to query"`
	RepoRoot string `json:"repo_root,omitempty" description:"Repository root (auto-detected if omitted)"`
}

func NewQueryTool(workingDir string) fantasy.AgentTool {
	return fantasy.NewAgentTool(
		QueryToolName,
		queryDescription,
		func(ctx context.Context, params QueryParams, _ fantasy.ToolCall) (fantasy.ToolResponse, error) {
			root := params.RepoRoot
			if root == "" {
				root = workingDir
			}
			// pattern/target/repo_root passed as argv (sys.argv[1..3]) so they
			// cannot break out of the Python -c string.
			expr := `
import json, sys
from better_code_review_graph.tools import query_graph
r = query_graph(pattern=sys.argv[1], target=sys.argv[2], repo_root=sys.argv[3])
sys.stdout.write(json.dumps(r))
`
			out, err := runBCRG(ctx, expr, params.Pattern, params.Target, root)
			if err != nil {
				return fantasy.NewTextErrorResponse(out + "\n" + err.Error()), nil
			}
			return fantasy.NewTextResponse(out), nil
		},
	)
}
