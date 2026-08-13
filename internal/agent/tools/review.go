package tools

import (
	"context"
	_ "embed"

	"charm.land/fantasy"
)

const ReviewToolName = "review"

//go:embed review.md
var reviewDescription string

type ReviewParams struct {
	RepoRoot string `json:"repo_root,omitempty" description:"Repository root (auto-detected if omitted)"`
	FromSHA  string `json:"from_sha,omitempty" description:"Starting commit SHA"`
	ToSHA    string `json:"to_sha,omitempty" description:"Ending commit SHA"`
}

func NewReviewTool(workingDir string) fantasy.AgentTool {
	return fantasy.NewAgentTool(
		ReviewToolName,
		reviewDescription,
		func(ctx context.Context, params ReviewParams, _ fantasy.ToolCall) (fantasy.ToolResponse, error) {
			root := params.RepoRoot
			if root == "" {
				root = workingDir
			}
			// repo_root/from_sha/to_sha passed as argv (sys.argv[1..3]) so they
			// cannot break out of the Python -c string.
			expr := `
import json, sys
from better_code_review_graph.tools import review_delta
r = review_delta(repo_root=sys.argv[1], from_sha=sys.argv[2], to_sha=sys.argv[3])
sys.stdout.write(json.dumps(r))
`
			out, err := runBCRG(ctx, expr, root, params.FromSHA, params.ToSHA)
			if err != nil {
				return fantasy.NewTextErrorResponse(out + "\n" + err.Error()), nil
			}
			return fantasy.NewTextResponse(out), nil
		},
	)
}
