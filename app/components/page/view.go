package page

import (
	"solopg/app/services/t"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Page struct {
	Title    string
	Subtitle string
	Body     string
}

type Template = Page

func NewPage(params Template) Page {
	return Page{
		Title:    params.Title,
		Subtitle: params.Subtitle,
		Body:     params.Body,
	}
}

func (p Page) View() tea.View {
	title := lipgloss.NewStyle().
		Bold(true).
		Render(t.Localize(p.Title))

	return tea.NewView(
		lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			"",
			t.Localize(p.Subtitle),
			p.Body,
		),
	)
}
