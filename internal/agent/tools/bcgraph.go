package tools

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"

	"charm.land/fantasy"
)

const BcgraphToolName = "bcgraph"

//go:embed bcgraph.md
var bcgraphDescription string

type BcgraphParams struct {
	Tool     string         `json:"tool" description:"better-code-review-graph function to call"`
	RepoRoot string         `json:"repo_root,omitempty" description:"Repository root (defaults to working dir)"`
	Params   map[string]any `json:"params,omitempty" description:"Keyword arguments for the tool"`
}

// bcgraphToolWhitelist mirrors the importable functions exposed by the
// better_code_review_graph.tools module to the Python shim.
var bcgraphToolWhitelist = []string{
	"build_or_update_graph", "query_graph", "review_delta", "diff_graph",
	"get_review_context", "get_docs_section", "get_impact_radius",
	"list_graph_stats", "semantic_search_nodes", "find_large_functions",
	"embed_graph", "export_graph_dispatch", "import_graph_dispatch", "renamed_in_diff",
}

func NewBcgraphTool(workingDir string) fantasy.AgentTool {
	return fantasy.NewAgentTool(
		BcgraphToolName,
		bcgraphDescription,
		func(ctx context.Context, params BcgraphParams, _ fantasy.ToolCall) (fantasy.ToolResponse, error) {
			tool := params.Tool
			valid := false
			for _, allowed := range bcgraphToolWhitelist {
				if tool == allowed {
					valid = true
					break
				}
			}
			if !valid {
				msg := fmt.Sprintf("unknown tool %q (allowed: %s)", tool, strings.Join(bcgraphToolWhitelist, ", "))
				return fantasy.NewTextErrorResponse(msg), nil
			}
			root := params.RepoRoot
			if root == "" {
				root = workingDir
			}
			raw, err := json.Marshal(params.Params)
			if err != nil {
				return fantasy.NewTextErrorResponse("bad params: " + err.Error()), nil
			}
			// Tool name, repo root, and JSON params travel as argv so nothing is
			// interpolated into the -c script.
			expr := `
import json, sys
import better_code_review_graph.tools as tools
tool = sys.argv[1]
root = sys.argv[2]
kwargs = json.loads(sys.argv[3])
if 'repo_root' in kwargs:
    root = kwargs.pop('repo_root')
fn = getattr(tools, tool)
r = fn(repo_root=root, **kwargs) if root else fn(**kwargs)
sys.stdout.write(json.dumps(r, default=str))
`
			out, err := runBCRG(ctx, expr, tool, root, string(raw))
			if err != nil {
				return fantasy.NewTextErrorResponse(out + "\n" + err.Error()), nil
			}
			return fantasy.NewTextResponse(out), nil
		},
	)
}
