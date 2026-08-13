package tools

import (
	"context"
	_ "embed"
	"strconv"

	"charm.land/fantasy"
	mem "github.com/Gentleman-Programming/engram/mem"
)

const MemSaveToolName = "mem_save"

//go:embed mem_save.md
var memSaveDescription string

type MemSaveParams struct {
	Title    string `json:"title" description:"Short, searchable title for the memory (e.g. 'JWT auth middleware')"`
	Content  string `json:"content" description:"Structured content (What/Why/Where/Learned)"`
	Type     string `json:"type,omitempty" description:"Category: decision, architecture, bugfix, pattern, config, discovery, learning, manual"`
	TopicKey string `json:"topic_key,omitempty" description:"Optional stable topic key for upserts"`
}

func NewMemSaveTool(workingDir, dataDir string) fantasy.AgentTool {
	return fantasy.NewAgentTool(
		MemSaveToolName,
		memSaveDescription,
		func(ctx context.Context, params MemSaveParams, _ fantasy.ToolCall) (fantasy.ToolResponse, error) {
			h, err := openMemStore(dataDir)
			if err != nil {
				return fantasy.NewTextResponse("failed to open engram store: " + err.Error()), nil
			}
			defer h.Close()

			id, err := h.SaveObservation(mem.AddObservationParams{
				Type:      params.Type,
				Title:     params.Title,
				Content:   params.Content,
				Project:   workingDir,
				TopicKey:  params.TopicKey,
				SessionID: "",
			})
			if err != nil {
				return fantasy.NewTextResponse("failed to save memory: " + err.Error()), nil
			}
			return fantasy.NewTextResponse("saved memory id " + strconv.FormatInt(id, 10)), nil
		},
	)
}
