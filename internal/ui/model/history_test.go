package model

import (
	"testing"

	"github.com/charmbracelet/crush/internal/config"
	"github.com/charmbracelet/crush/internal/workspace"
	"github.com/stretchr/testify/require"
)

type historyWorkspace struct {
	workspace.Workspace
}

func (historyWorkspace) Config() *config.Config {
	return &config.Config{}
}

func (historyWorkspace) PermissionSkipRequests() bool {
	return false
}

func TestHistoryBangCommandStripsPrefixWhileAlreadyInBangMode(t *testing.T) {
	t.Parallel()

	u := newTestUI()
	u.com.Workspace = historyWorkspace{}
	u.promptHistory.messages = []string{"!echo one", "!echo two"}
	u.promptHistory.index = -1

	require.True(t, u.historyPrev())
	require.True(t, u.bangMode)
	require.Equal(t, "echo one", u.textarea.Value())

	require.True(t, u.historyPrev())
	require.True(t, u.bangMode)
	require.Equal(t, "echo two", u.textarea.Value())
}

// TestRecordPromptHistoryPrepends pins that a just-submitted prompt lands at
// the newest history slot so up-arrow recalls it even before the async DB
// write lands.
func TestRecordPromptHistoryPrepends(t *testing.T) {
	t.Parallel()

	u := newTestUI()
	u.promptHistory.messages = []string{"older"}

	u.recordPromptHistory("latest")
	require.Equal(t, []string{"latest", "older"}, u.promptHistory.messages)
	require.Equal(t, -1, u.promptHistory.index, "recording must reset the cursor to the draft")
}

// TestRecordPromptHistorySkipsDuplicate pins that re-recording the current
// newest prompt does not duplicate it.
func TestRecordPromptHistorySkipsDuplicate(t *testing.T) {
	t.Parallel()

	u := newTestUI()
	u.recordPromptHistory("same")
	u.recordPromptHistory("same")
	require.Equal(t, []string{"same"}, u.promptHistory.messages)
}

// TestRecordPromptHistoryIgnoresEmpty pins that an empty prompt is not added.
func TestRecordPromptHistoryIgnoresEmpty(t *testing.T) {
	t.Parallel()

	u := newTestUI()
	u.recordPromptHistory("")
	require.Empty(t, u.promptHistory.messages)
}
