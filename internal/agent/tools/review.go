package tools

import (
	"context"
	_ "embed"
	"fmt"

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
			expr := fmt.Sprintf(`
import json, sys
from better_code_review_graph.tools import review_delta
r = review_delta(repo_root=%q, from_sha=%q, to_sha=%q)
sys.stdout.write(json.dumps(r))
`, root, params.FromSHA, params.ToSHA)
			out, err := runBCRG(expr)
			if err != nil {
				return fantasy.NewTextResponse(out + "\n" + err.Error()), nil
			}
			return fantasy.NewTextResponse(out), nil
		},
	)
}
