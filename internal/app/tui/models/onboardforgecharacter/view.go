package onboardforgecharacter

import (
	"fmt"
	"solopg/internal/infrastructure/t"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	content := strings.Builder{}

	content.WriteString(t.Localize("stats"))
	buildsChoiceStr := make([]string, len(m.reroll.Value))
	for i, modifier := range m.reroll.Value {
		key := "stat." + string(modifier.Stat)
		label := t.Localize(key)
		buildsChoiceStr[i] = fmt.Sprintf("%s: %d", label, modifier.Value)
	}
	content.WriteString(strings.Join(buildsChoiceStr, ", "))
	content.WriteString(" ")
	content.WriteString(t.Localize("attempt"))
	content.WriteString(": ")
	content.WriteString(fmt.Sprintf("%d", m.reroll.Attempt))

	content.WriteString("\n\n")
	content.WriteString(m.reroll.Options.View())

	return tea.NewView(content.String())
}
