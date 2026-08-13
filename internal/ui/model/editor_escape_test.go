package model

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/stretchr/testify/require"
)

// TestEscapeClearsTextarea pins that a plain esc in the editor clears the
// draft contents (previously esc double-pressed to cancel; now esc clears
// input and ctrl+esc stops the task).
func TestEscapeClearsTextarea(t *testing.T) {
	ws := &countingWorkspace{ready: true}
	m := newBusyUI(ws)

	m.textarea.Focus()
	m.textarea.SetValue("hello world")

	_, _ = m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})

	require.Empty(t, m.textarea.Value(), "esc must clear the draft contents")
}

// TestEscapeOnEmptyTextareaIsNoop pins that esc on an already-empty draft
// does not panic and leaves the state intact.
func TestEscapeOnEmptyTextareaIsNoop(t *testing.T) {
	ws := &countingWorkspace{ready: true}
	m := newBusyUI(ws)

	m.textarea.Focus()
	m.textarea.SetValue("")

	_, _ = m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})

	require.Empty(t, m.textarea.Value())
}
