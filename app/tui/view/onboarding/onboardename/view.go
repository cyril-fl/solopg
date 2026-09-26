package onboardename

import (
	"solopg/app/services/i18n"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
	title := lipgloss.NewStyle().
		Bold(true).
		Render(i18n.Localize("onboarding"))

	input := m.Input.View()

	return tea.NewView(
		lipgloss.JoinVertical(
			lipgloss.Left,
			/*
				TODO LOW verrifier pourquoi CREATE CHARACTER apparais pas sur les autre view
			*/
			title,
			"",
			i18n.Localize("name"),
			input,
		),
	)
}
