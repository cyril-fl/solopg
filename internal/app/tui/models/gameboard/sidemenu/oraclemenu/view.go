package oraclemenu

import (
	"solopg/internal/infrastructure/t"

	"charm.land/lipgloss/v2"
)

func (m *OracleMenu) GetMenuView() string {
	title := lipgloss.NewStyle().Bold(true).Render(t.Localize(m.id))
	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		m.list.View(),
	)
}

func (m *OracleMenu) GetView() string {
	return ""
}

func (m *OracleMenu) GetFooter() []string {
	return []string{}
}
