package dialog

import (
	"strings"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/crush/internal/ui/common"
	uv "github.com/charmbracelet/ultraviolet"
)

// EditorHelpID is the identifier for the editor help dialog.
const EditorHelpID = "editor_help"

// HelpRow is a single human-readable keybinding row.
type HelpRow struct {
	Keys string // display string, e.g. "ctrl+a"
	Desc string // e.g. "Select all"
}

// HelpSection is a titled group of keybinding rows. Pairs are laid side by
// side in the dialog to keep the panel compact.
type HelpSection struct {
	Title string
	Rows  []HelpRow
}

// EditorHelp displays a human-readable, grouped summary of editor keybindings
// in a centered dialog overlay.
type EditorHelp struct {
	com    *common.Common
	frame  []HelpSection
	keyMap struct {
		Close key.Binding
	}
}

var _ Dialog = (*EditorHelp)(nil)

// NewEditorHelp creates a new editor keybinding help dialog.
func NewEditorHelp(com *common.Common, sections []HelpSection) *EditorHelp {
	e := &EditorHelp{
		com:   com,
		frame: sections,
	}
	e.keyMap.Close = CloseKey
	return e
}

// ID implements [Dialog].
func (*EditorHelp) ID() string {
	return EditorHelpID
}

// HandleMsg implements [Dialog].
func (e *EditorHelp) HandleMsg(msg tea.Msg) Action {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if key.Matches(msg, e.keyMap.Close) {
			return ActionClose{}
		}
	}
	return nil
}

// Draw implements [Dialog].
func (e *EditorHelp) Draw(scr uv.Screen, area uv.Rectangle) *tea.Cursor {
	const title = "Editor Keys"

	innerWidth := min(area.Dx()-4, 110)
	contentWidth := max(0, innerWidth-2)
	maxHeight := max(0, area.Dy()-10)

	body := e.renderSections(contentWidth, maxHeight)

	hint := e.com.Styles.Dialog.SecondaryText.Render("esc close")
	uhint := lipgloss.NewStyle().Width(contentWidth).Align(lipgloss.Right).Render(hint)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		body,
		"",
		uhint,
	)

	content = lipgloss.NewStyle().Width(contentWidth).Render(content)
	view := e.com.Styles.Dialog.View.Render(content)
	DrawCenter(scr, area, view)

	return nil
}

// renderSections lays sections out in a two-column grid: every block shares
// the same key-column width and the same block width, so all columns line up
// vertically across all rows. Truncated to maxHeight rows.
func (e *EditorHelp) renderSections(contentWidth, maxHeight int) string {
	// Global max key width across ALL sections so every block's description
	// column starts at the same x position.
	maxKey := 0
	for _, sec := range e.frame {
		for _, r := range sec.Rows {
			if w := lipgloss.Width(r.Keys); w > maxKey {
				maxKey = w
			}
		}
	}

	// Two equal columns with a 4-space gutter between them.
	colWidth := max(0, (contentWidth-4)/2)
	blk := func(idx int) string {
		if idx >= len(e.frame) {
			return ""
		}
		return e.renderSection(e.frame[idx], maxKey, colWidth)
	}

	var rows []string
	for i := 0; i < len(e.frame); i += 2 {
		left := blk(i)
		right := blk(i + 1)
		row := left
		if right != "" {
			row = lipgloss.JoinHorizontal(lipgloss.Top, left, strings.Repeat(" ", 4), right)
		}
		rows = append(rows, row)
	}
	joined := strings.Join(rows, "\n")
	if maxHeight > 0 && lipgloss.Height(joined) > maxHeight {
		joined = lipgloss.NewStyle().Height(maxHeight).Render(joined)
	}
	return joined
}

// renderSection renders one titled block of key rows, padded to width so both
// grid columns are the same width.
func (e *EditorHelp) renderSection(sec HelpSection, maxKey, width int) string {
	s := e.com.Styles
	kStyle := s.Dialog.Help.FullKey.Inline(true)
	dStyle := s.Dialog.Help.FullDesc.Inline(true)

	var b strings.Builder
	b.WriteString(s.Dialog.TitleText.Render(sec.Title))
	for _, r := range sec.Rows {
		pad := maxKey - lipgloss.Width(r.Keys)
		b.WriteString("\n")
		b.WriteString(kStyle.Render(r.Keys + strings.Repeat(" ", pad)))
		b.WriteString(" ")
		b.WriteString(dStyle.Render(r.Desc))
	}
	return lipgloss.NewStyle().Width(width).Render(b.String())
}

// ShortHelp implements [help.KeyMap].
func (e *EditorHelp) ShortHelp() []key.Binding {
	return []key.Binding{e.keyMap.Close}
}

// FullHelp implements [help.KeyMap].
func (e *EditorHelp) FullHelp() [][]key.Binding {
	return [][]key.Binding{{e.keyMap.Close}}
}
