package dialog

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/crush/internal/ui/common"
	"github.com/charmbracelet/crush/internal/ui/styles"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/stretchr/testify/require"
)

func newTestEditorHelp() *EditorHelp {
	s := styles.CharmtonePantera()
	com := &common.Common{Styles: &s}
	return NewEditorHelp(com, []HelpSection{
		{Title: "Test", Rows: []HelpRow{{Keys: "ctrl+a", Desc: "Select all"}}},
	})
}

// TestEditorHelp_ID returns the registered dialog identifier.
func TestEditorHelp_ID(t *testing.T) {
	t.Parallel()
	require.Equal(t, "editor_help", newTestEditorHelp().ID())
}

// TestEditorHelp_CloseKey pins that esc closes the dialog.
func TestEditorHelp_CloseKey(t *testing.T) {
	t.Parallel()
	e := newTestEditorHelp()
	action := e.HandleMsg(tea.KeyPressMsg{Code: tea.KeyEscape})
	require.IsType(t, ActionClose{}, action, "esc must close the editor help dialog")
}

// TestEditorHelp_AltEscapeCloses pins that alt+esc also closes.
func TestEditorHelp_AltEscapeCloses(t *testing.T) {
	t.Parallel()
	e := newTestEditorHelp()
	action := e.HandleMsg(tea.KeyPressMsg{Code: tea.KeyEscape, Mod: tea.ModAlt})
	require.IsType(t, ActionClose{}, action)
}

// TestEditorHelp_OtherKeysNoop verifies unrelated keys do nothing.
func TestEditorHelp_OtherKeysNoop(t *testing.T) {
	t.Parallel()
	e := newTestEditorHelp()
	require.Nil(t, e.HandleMsg(tea.KeyPressMsg{Code: tea.KeyEnter}))
	require.Nil(t, e.HandleMsg(tea.KeyPressMsg{Code: tea.KeySpace}))
	require.Nil(t, e.HandleMsg(tea.MouseWheelMsg{}))
}

// TestEditorHelp_DrawNoPanic renders the dialog at a range of sizes and
// ensures it never panics, covering narrow and small terminals.
func TestEditorHelp_DrawNoPanic(t *testing.T) {
	t.Parallel()
	e := newTestEditorHelp()
	scr := uv.NewScreenBuffer(80, 24)
	for _, size := range [][2]int{
		{80, 24}, {40, 12}, {20, 6}, {5, 3}, {0, 0},
	} {
		_ = e.Draw(scr, uv.Rect(0, 0, size[0], size[1]))
	}
}
