package codexform

import (
	"solopg/internal/infrastructure/t"

	"charm.land/lipgloss/v2"
)

func (m Model) View() string {
	parts := []string{lipgloss.NewStyle().Bold(true).Render(t.Localize("codex.add", map[string]any{"Kind": string(m.kind)})), ""}
	for i, item := range m.fields {
		label := item.label
		if i == m.index {
			label = "▸ " + label
		}
		parts = append(parts, label)
		if len(item.choices) > 0 {
			for choiceIndex, choice := range item.choices {
				marker := "  "
				if choiceIndex == item.choiceIndex {
					marker = "› "
				}
				parts = append(parts, marker+choice)
			}
		} else {
			parts = append(parts, item.input.View())
		}
	}
	if m.err != "" {
		parts = append(parts, "", lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(m.err))
	}
	return lipgloss.JoinVertical(lipgloss.Left, parts...)
}
