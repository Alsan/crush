package tools

import (
	"context"
	_ "embed"
	"fmt"
	"strings"

	"charm.land/fantasy"
)

const MemContextToolName = "mem_context"

//go:embed mem_context.md
var memContextDescription string

type MemContextParams struct {
	Scope string `json:"scope,omitempty" description:"Scope filter: project (default) or personal"`
	Limit int    `json:"limit,omitempty" description:"Max recent memories"`
}

func NewMemContextTool(workingDir, dataDir string) fantasy.AgentTool {
	return fantasy.NewAgentTool(
		MemContextToolName,
		memContextDescription,
		func(ctx context.Context, params MemContextParams, _ fantasy.ToolCall) (fantasy.ToolResponse, error) {
			h, err := openMemStore(dataDir)
			if err != nil {
				return fantasy.NewTextResponse("failed to open engram store: " + err.Error()), nil
			}
			defer h.Close()

			// RecentObservations filters by project when project is non-empty;
			// use the working dir as the project so this surfaces the current
			// project's context. For "personal", pass an empty project to
			// include all projects.
			project := workingDir
			if params.Scope == "personal" {
				project = ""
			}

			obs, err := h.RecentObservations(project, params.Scope, params.Limit)
			if err != nil {
				return fantasy.NewTextResponse("context load failed: " + err.Error()), nil
			}
			if len(obs) == 0 {
				return fantasy.NewTextResponse("no memories found"), nil
			}

			var b strings.Builder
			for _, o := range obs {
				fmt.Fprintf(&b, "[id=%d %s] %s\n%s\n\n", o.ID, o.Type, o.Title, o.Content)
			}
			return fantasy.NewTextResponse(b.String()), nil
		},
	)
}
