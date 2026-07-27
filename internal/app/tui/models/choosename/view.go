package choosename

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)


func (m model) View() tea.View {
	title := lipgloss.NewStyle().
		Bold(true).
		Render("Create character")

	input := m.Input.View()

	return tea.NewView(
		lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			"",
			"Name:",
			input,
		),
	)
}
