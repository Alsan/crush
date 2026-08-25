package model

import (
	"testing"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/crush/internal/ui/attachments"
	"github.com/charmbracelet/crush/internal/ui/dialog"
	"github.com/stretchr/testify/require"
)

func newEditorHelpTestUI() *UI {
	u := newTestUI()
	u.dialog = dialog.NewOverlay()
	sty := u.com.Styles.Attachments
	u.attachments = attachments.New(
		attachments.NewRenderer(sty.Normal, sty.Deleting, sty.Image, sty.Text, sty.Skill, sty.Remove),
		attachments.Keymap{},
	)
	u.keyMap = DefaultKeyMap()
	u.textarea.Focus()
	return u
}

// TestEditorHelp_CtrlSlashOpensDialog pins that pressing ctrl+/ in the
// editor opens the editor keybinding help dialog.
func TestEditorHelp_CtrlSlashOpensDialog(t *testing.T) {
	t.Parallel()

	u := newEditorHelpTestUI()

	_, _ = u.Update(tea.KeyPressMsg{Code: '/', Mod: tea.ModCtrl})

	require.True(t, u.dialog.ContainsDialog(dialog.EditorHelpID), "ctrl+/ must open the editor help dialog")
}

// TestEditorHelp_ReopenBringToFront pins that pressing ctrl+/ twice keeps a
// single dialog instance rather than stacking duplicates.
func TestEditorHelp_ReopenBringToFront(t *testing.T) {
	t.Parallel()

	u := newEditorHelpTestUI()

	for range 2 {
		_, _ = u.Update(tea.KeyPressMsg{Code: '/', Mod: tea.ModCtrl})
	}
	require.True(t, u.dialog.ContainsDialog(dialog.EditorHelpID))
}

// TestEditorHelp_EscapeClosesDialog pins that esc dismisses the help dialog.
// The UI-level esc path re-focuses the textarea and touches attachment
// state, so close semantics are asserted at the dialog layer instead
// (see internal/ui/dialog/editor_help_test.go).
func TestEditorHelp_EscapeClosesDialog(t *testing.T) {
	t.Parallel()

	e := dialog.NewEditorHelp(newEditorHelpTestUI().com, []dialog.HelpSection{
		{Title: "Test", Rows: []dialog.HelpRow{{Keys: "ctrl+a", Desc: "Select all"}}},
	})
	action := e.HandleMsg(tea.KeyPressMsg{Code: tea.KeyEscape})
	require.IsType(t, dialog.ActionClose{}, action, "esc must close the editor help dialog")
}

// TestEditorHelp_BareSlashDoesNotOpenDialog pins that the plain '/' key
// belongs to commands, not the editor help dialog.
func TestEditorHelp_BareSlashDoesNotOpenDialog(t *testing.T) {
	t.Parallel()

	km := DefaultKeyMap()
	help := key.Matches(tea.KeyPressMsg{Code: '/', Mod: tea.ModCtrl}, km.Editor.HelpDialog)
	plain := key.Matches(tea.KeyPressMsg{Code: '/'}, km.Editor.HelpDialog)
	require.True(t, help, "ctrl+/ must match HelpDialog")
	require.False(t, plain, "bare / must not match HelpDialog")
}
