package codexmenu

import (
	"fmt"
	"solopg/app/services/i18n"
	"solopg/app/utils/transform"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// - View - //
func (m *codexMenu) GetMenuView() string {
	title := transform.Uppercase(i18n.Localize(m.id))
	title = lipgloss.NewStyle().Bold(true).Render(title)

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
		return i18n.Localize("error.page_not_found")
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
	return fmt.Sprintf("%s\n\n", i18n.Localize("sidemenu.title:page", map[string]any{
		"Submenu": transform.Capitalize(i18n.Localize(m.id)),
	}))
}

// Footer
func (m *codexMenu) GetFooter() []string {
	footer := []string{}

	if m.IsOpen() {
		footer = append(footer, i18n.Localize("shift-enter:add"))
	} else {
		footer = append(footer, i18n.Localize("shift-enter:open"))
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
