package codexmenu

import (
	"solopg/app/services/t"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// - View - //
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
		return t.Localize("error.page_not_found")
	}

	return page.GetItemView()
}

func (m *codexMenuItem) GetItemView() string {
	page := m.GetPageHeader()

	if m.showForm {
		return page + m.form.View().Content
	} else {
		return page + strings.Join(m.table.Summaries(), "\n\n")
	}
}

func (m *codexMenuItem) GetPageHeader() string {
	return "Page: " + t.Localize(m.id) + "\n\n"
}

// Footer
func (m *codexMenu) GetFooter() []string {
	footer := []string{}

	if m.IsOpen() {
		footer = append(footer, t.Localize("shift-enter:add"))
	} else {
		footer = append(footer, t.Localize("shift-enter:open"))
	}

	return footer
}

// - Handlers - //
func (m *codexMenu) HandleWindowResize(msg tea.WindowSizeMsg) tea.Cmd {
	if form := m.GetFormFromCurrentPage(); form != nil {
		form.Update(msg)
	}

	return nil
}
