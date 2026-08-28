package model

import (
	"testing"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/crush/internal/ui/attachments"
	"github.com/charmbracelet/crush/internal/ui/dialog"
	"github.com/stretchr/testify/require"
)

func newEditorHelpTestUI(t *testing.T) *UI {
	t.Helper()
	u := newTestUI()
	u.dialog = dialog.NewOverlay()
	sty := u.com.Styles.Attachments
	u.attachments = attachments.New(
		attachments.NewRenderer(sty.Normal, sty.Deleting, sty.Image, sty.Text, sty.Skill, sty.Remove, sty.Remove),
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

	u := newEditorHelpTestUI(t)

	_, _ = u.Update(tea.KeyPressMsg{Code: '/', Mod: tea.ModCtrl})

	require.True(t, u.dialog.ContainsDialog(dialog.EditorHelpID), "ctrl+/ must open the editor help dialog")
}

// TestEditorHelp_ReopenDoesNotStack pins that pressing ctrl+/ twice keeps a
// single dialog instance: closing once removes it entirely,
// rather than leaving a duplicate behind.
func TestEditorHelp_ReopenDoesNotStack(t *testing.T) {
	t.Parallel()

	u := newEditorHelpTestUI(t)

	for range 2 {
		_, _ = u.Update(tea.KeyPressMsg{Code: '/', Mod: tea.ModCtrl})
	}
	require.True(t, u.dialog.ContainsDialog(dialog.EditorHelpID))

	u.dialog.CloseDialog(dialog.EditorHelpID)
	require.False(t, u.dialog.ContainsDialog(dialog.EditorHelpID),
		"a single close must remove all help dialog instances")
}

// TestEditorHelp_SlashKeyIsolation pins that ctrl+/ belongs to the editor
// help dialog while bare '/' stays with the commands binding.
func TestEditorHelp_SlashKeyIsolation(t *testing.T) {
	t.Parallel()

	km := DefaultKeyMap()
	ctrl := key.Matches(tea.KeyPressMsg{Code: '/', Mod: tea.ModCtrl}, km.Editor.HelpDialog)
	plainHelp := key.Matches(tea.KeyPressMsg{Code: '/'}, km.Editor.HelpDialog)
	plainCommands := key.Matches(tea.KeyPressMsg{Code: '/'}, km.Editor.Commands)
	require.True(t, ctrl, "ctrl+/ must match HelpDialog")
	require.False(t, plainHelp, "bare / must not match HelpDialog")
	require.True(t, plainCommands, "bare / must match Commands")
}
