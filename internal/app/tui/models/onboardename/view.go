package onboardename

import (
	"solopg/internal/infrastructure/t"

	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
	title := lipgloss.NewStyle().
		Bold(true).
		Render(t.Local.MustLocalize(&goi18n.LocalizeConfig{MessageID: "create_character"}))

	input := m.Input.View()

	return tea.NewView(
		lipgloss.JoinVertical(
			lipgloss.Left,
			// TODO verrifier pourquoi CREATE CHARACTER apparais pas sur les autre view
			title,
			"",
			t.Local.MustLocalize(&goi18n.LocalizeConfig{MessageID: "name"}),
			input,
		),
	)
}
