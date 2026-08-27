package onboardlocation

import (
	"fmt"
	"solopg/internal/app/tui"
	"solopg/internal/infrastructure/t"
	"strings"

	tea "charm.land/bubbletea/v2"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

func (m model) View() tea.View {
	content := strings.Builder{}

	content.WriteString(t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "location"}))
	content.WriteString(m.reroll.Value.Name)
	content.WriteString(" " + t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "attempt"}) + ": ")
	content.WriteString(fmt.Sprintf("%d", m.reroll.Attempt))

	if m.reroll.Value.Description.Description != tui.EmptyKey {
		content.WriteString("\n")
		content.WriteString(m.reroll.Value.Description.Description)
	}
	content.WriteString("\n\n")
	content.WriteString(m.reroll.Options.View())

	return tea.NewView(content.String())
}
