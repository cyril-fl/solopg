package codexmenu

import (
	"solopg/internal/infrastructure/t"

	"charm.land/lipgloss/v2"
)

func (m *CodexMenu) GetMenuView() string {
	title := lipgloss.NewStyle().Bold(true).Render(t.Localize(m.id))
	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		m.list.View(),
	)
}

func (m *CodexMenu) GetView() string {
	page := m.getCurrentPage()
	if page == nil {
		// i18N
		return "page not found"
	}

	return page.View()
}

func (m *CodexMenu) GetFooter() []string {
	return []string{
		t.Localize("shift-open"),
	}
}
