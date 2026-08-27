package onboardlocation

import (
	"fmt"
	"solopg/internal/app/tui"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	content := strings.Builder{}

	content.WriteString("Location: ")
	content.WriteString(m.reroll.Value.Name)
	content.WriteString(" Attempt: ")
	content.WriteString(fmt.Sprintf("%d", m.reroll.Attempt))

	if m.reroll.Value.Description.Description != tui.EmptyKey {
		content.WriteString("\n")
		content.WriteString(m.reroll.Value.Description.Description)
	}
	content.WriteString("\n\n")
	content.WriteString(m.reroll.Options.View())

	return tea.NewView(content.String())
}
