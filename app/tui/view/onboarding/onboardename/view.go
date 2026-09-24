package onboardename

import (
	"solopg/app/components/page"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	page := page.NewPage(page.Template{
		Title:    "create_character",
		Subtitle: "name",
		Body:     m.Input.View(),
	})

	return page.View()
}
