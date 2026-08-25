package dialog

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/crush/internal/ui/common"
	"github.com/charmbracelet/crush/internal/ui/styles"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/stretchr/testify/require"
)

func newTestEditorHelp(t *testing.T) *EditorHelp {
	t.Helper()
	s := styles.CharmtonePantera()
	com := &common.Common{Styles: &s}
	return NewEditorHelp(com, []HelpSection{
		{Title: "Test", Rows: []HelpRow{{Keys: "ctrl+a", Desc: "Select all"}}},
	})
}

// TestEditorHelp_ID returns the registered dialog identifier.
func TestEditorHelp_ID(t *testing.T) {
	t.Parallel()
	require.Equal(t, "editor_help", newTestEditorHelp(t).ID())
}

// TestEditorHelp_CloseKey pins that esc (with or without alt) closes the
// dialog.
func TestEditorHelp_CloseKey(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		mod  tea.KeyMod
	}{
		{"esc", 0},
		{"alt+esc", tea.ModAlt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e := newTestEditorHelp(t)
			action := e.HandleMsg(tea.KeyPressMsg{Code: tea.KeyEscape, Mod: tt.mod})
			require.IsType(t, ActionClose{}, action, "escaping must close the editor help dialog")
		})
	}
}

// TestEditorHelp_OtherKeysNoop verifies unrelated keys do nothing.
func TestEditorHelp_OtherKeysNoop(t *testing.T) {
	t.Parallel()
	e := newTestEditorHelp(t)
	require.Nil(t, e.HandleMsg(tea.KeyPressMsg{Code: tea.KeyEnter}))
	require.Nil(t, e.HandleMsg(tea.KeyPressMsg{Code: tea.KeySpace}))
	require.Nil(t, e.HandleMsg(tea.MouseWheelMsg{}))
}

// TestEditorHelp_DrawNoPanic renders the dialog at a range of sizes and
// ensures it never panics, covering narrow and small terminals.
func TestEditorHelp_DrawNoPanic(t *testing.T) {
	t.Parallel()
	e := newTestEditorHelp(t)
	for _, size := range [][2]int{
		{80, 24}, {40, 12}, {20, 6}, {5, 3}, {0, 0},
	} {
		scr := uv.NewScreenBuffer(size[0], size[1])
		_ = e.Draw(scr, uv.Rect(0, 0, size[0], size[1]))
	}
}
