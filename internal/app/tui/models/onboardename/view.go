package onboardename

import (
	"solopg/internal/infrastructure/t"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
	title := lipgloss.NewStyle().
		Bold(true).
		Render(t.Localize("create_character"))

	input := m.Input.View()

	return tea.NewView(
		lipgloss.JoinVertical(
			lipgloss.Left,
			/*
				TODO LOW verrifier pourquoi CREATE CHARACTER apparais pas sur les autre view
			*/
			title,
			"",
			t.Localize("name"),
			input,
		),
	)
}
