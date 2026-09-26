package onboardlocation

import (
	"fmt"
	"solopg/app/services/i18n"
	"solopg/app/tui"
	"solopg/app/utils/transform"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	content := strings.Builder{}

	fmt.Fprintf(&content, "%s: %s", transform.Capitalize(i18n.Localize("location")), i18n.Localize(m.reroll.Value.Name))
	content.WriteString(" | ")
	fmt.Fprintf(&content, "%s: %d", transform.Capitalize(i18n.Localize("attempt")), m.reroll.Attempt)

	if m.reroll.Value.Description.Description != tui.EmptyKey {
		content.WriteString("\n")
		content.WriteString(i18n.Localize(m.reroll.Value.Description.Description))
	}

	content.WriteString("\n\n")
	content.WriteString(m.reroll.Options.View())

	return tea.NewView(content.String())
}
