package onboardforgecharacter

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	content := strings.Builder{}

	content.WriteString("Stats: ")
	buildsChoiceStr := make([]string, len(m.reroll.Value))
	for i, modifier := range m.reroll.Value {
		buildsChoiceStr[i] = modifier.String()
	}
	content.WriteString(strings.Join(buildsChoiceStr, ", "))
	content.WriteString(" Attempt: ")
	content.WriteString(fmt.Sprintf("%d", m.reroll.Attempt))

	content.WriteString("\n\n")
	content.WriteString(m.reroll.Options.View())

	return tea.NewView(content.String())
}
