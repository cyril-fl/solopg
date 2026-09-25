package hintmenu

import (
	"solopg/app/services/t"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// - View - //
func (m *HintMenu) GetMenuView() string {
	title := lipgloss.NewStyle().Bold(true).Render(t.Localize(m.id))
	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		m.list.View(),
	)
}

func (m *HintMenu) GetView() string {
	return ""
}

func (m *HintMenu) GetFooter() []string {
	return []string{
		t.Localize("shift-enter:roll"),
	}
}

// - Handlers --//
func (m *HintMenu) HandleWindowResize(msg tea.WindowSizeMsg) tea.Cmd {
	// m.list.SetSize(msg.Width, msg.Height)
	return nil
}
