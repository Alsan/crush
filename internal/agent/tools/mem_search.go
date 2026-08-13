package tools

import (
	"context"
	_ "embed"
	"fmt"
	"strings"

	"charm.land/fantasy"
	mem "github.com/Gentleman-Programming/engram/mem"
)

const MemSearchToolName = "mem_search"

//go:embed mem_search.md
var memSearchDescription string

type MemSearchParams struct {
	Query     string `json:"query" description:"Search text, natural language or keywords"`
	Type      string `json:"type,omitempty" description:"Filter by observation type: decision, architecture, bugfix, pattern, config, discovery, learning, manual"`
	MatchMode string `json:"match_mode,omitempty" description:"Token matching: all (default) or any"`
	Limit     int    `json:"limit,omitempty" description:"Max results (default 10)"`
}

func NewMemSearchTool(workingDir, dataDir string) fantasy.AgentTool {
	return fantasy.NewAgentTool(
		MemSearchToolName,
		memSearchDescription,
		func(ctx context.Context, params MemSearchParams, _ fantasy.ToolCall) (fantasy.ToolResponse, error) {
			h, err := openMemStore(dataDir)
			if err != nil {
				return fantasy.NewTextErrorResponse("failed to open engram store: " + err.Error()), nil
			}

			results, err := h.Search(params.Query, mem.SearchOptions{
				Type:      params.Type,
				Project:   workingDir,
				MatchMode: params.MatchMode,
				Limit:     params.Limit,
			})
			if err != nil {
				return fantasy.NewTextErrorResponse("search failed: " + err.Error()), nil
			}
			if len(results) == 0 {
				return fantasy.NewTextResponse("no memories found"), nil
			}

			var b strings.Builder
			for _, r := range results {
				fmt.Fprintf(&b, "[id=%d rank=%.2f %s] %s\n%s\n\n",
					r.ID, r.Rank, r.Type, r.Title, r.Content)
			}
			return fantasy.NewTextResponse(b.String()), nil
		},
	)
}
