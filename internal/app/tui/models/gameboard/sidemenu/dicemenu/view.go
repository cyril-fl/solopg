package dicemenu

import (
	"solopg/internal/infrastructure/t"

	"charm.land/lipgloss/v2"
)

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
