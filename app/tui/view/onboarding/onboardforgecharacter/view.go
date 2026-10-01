package onboardforgecharacter

import (
	"fmt"
	"solopg/app/services/i19n"
	"solopg/app/utils/transform"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	content := strings.Builder{}

	fmt.Fprintf(&content, "%s: ", transform.Capitalize(i19n.Localize("stats")))

	for _, modifier := range m.reroll.Value {
		label := i19n.Localize("stat." + string(modifier.Stat))
		fmt.Fprintf(&content, "%s: %d ", label, modifier.Value)
	}

	content.WriteString(" | ")
	fmt.Fprintf(&content, "%s: %d ", transform.Capitalize(i19n.Localize("attempt")), m.reroll.Attempt)

	content.WriteString("\n\n")
	content.WriteString(m.reroll.Options.View())

	return tea.NewView(content.String())
}
