package dicemenu

import (
	"solopg/internal/infrastructure/t"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// - View - //
func (m *DiceMenu) GetMenuView() string {
	title := lipgloss.NewStyle().Bold(true).Render(t.Localize(m.id))
	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		m.list.View(),
	)
}

func (m *DiceMenu) GetView() string {
	return ""
}

func (m *DiceMenu) GetFooter() []string {
	return []string{
		t.Localize("shift-enter:roll"),
	}
}

// - Handlers - //
func (m *DiceMenu) HandleWindowResize(msg tea.WindowSizeMsg) tea.Cmd {
	// m.list.SetSize(msg.Width, msg.Height)
	return nil
}
