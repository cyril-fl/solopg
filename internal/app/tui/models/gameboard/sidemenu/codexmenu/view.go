package codexmenu

import (
	"solopg/internal/infrastructure/t"
	"strings"

	"charm.land/lipgloss/v2"
)

func (m *codexMenu) GetMenuView() string {
	title := lipgloss.NewStyle().Bold(true).Render(t.Localize(m.id))
	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		m.list.View(),
	)
}

func (m *codexMenu) GetView() string {
	page := m.getCurrentPage()
	if page == nil {
		// i18N
		return "page not found"
	}

	return page.GetItemView()
}

func (m *codexMenuItem) GetItemView() string {
	page := "Page: " + t.Localize(m.id) + "\n\n"

	if m.showForm {
		newform := m.form()
		return page + newform.View()
	} else {
		return page + strings.Join(m.table.Summaries(), "\n\n")
	}
}

func (m *codexMenu) GetFooter() []string {
	footer := []string{}

	if m.IsOpen() {
		footer = append(footer, t.Localize("shift-enter:add"))
	} else {
		footer = append(footer, t.Localize("shift-enter:open"))
	}

	return footer
}
