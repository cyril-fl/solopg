package onboardforgecharacter

import (
	"fmt"
	"solopg/internal/infrastructure/t"
	"strings"

	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	content := strings.Builder{}

	content.WriteString(t.Local.MustLocalize(&goi18n.LocalizeConfig{MessageID: "stats"}))
	buildsChoiceStr := make([]string, len(m.reroll.Value))
	for i, modifier := range m.reroll.Value {
		key := "stat." + string(modifier.Stat)
		label := t.Local.MustLocalize(&goi18n.LocalizeConfig{
			MessageID: key,
			DefaultMessage: &goi18n.Message{
				ID:    key,
				Other: string(modifier.Stat),
			},
		})
		buildsChoiceStr[i] = fmt.Sprintf("%s: %d", label, modifier.Value)
	}
	content.WriteString(strings.Join(buildsChoiceStr, ", "))
	content.WriteString(" " + t.Local.MustLocalize(&goi18n.LocalizeConfig{MessageID: "attempt"}) + ": ")
	content.WriteString(fmt.Sprintf("%d", m.reroll.Attempt))

	content.WriteString("\n\n")
	content.WriteString(m.reroll.Options.View())

	return tea.NewView(content.String())
}
