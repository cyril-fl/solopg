package onboardforgecharacter

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	content := strings.Builder{}

	content.WriteString("Stats: ")
	buildsChoiceStr := make([]string, len(m.buildsChoice))
	for i, modifier := range m.buildsChoice {
		buildsChoiceStr[i] = modifier.String()
	}
	content.WriteString(strings.Join(buildsChoiceStr, ", "))
	content.WriteString(" Attempt: ")
	content.WriteString(fmt.Sprintf("%d", m.attempt))

	content.WriteString("\n\n")
	content.WriteString(m.choiceList.View())

	return tea.NewView(content.String())
}
