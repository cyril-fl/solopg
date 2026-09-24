package onboardforgecharacter

import (
	"fmt"
	"solopg/app/components/page"
	"solopg/app/services/t"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	page := page.NewPage(page.Template{
		Title:    "create_character",
		Subtitle: "stats",
		Body:     m.getContent(),
	})

	return page.View()
}

func (m model) getContent() string {
	content := strings.Builder{}

	buildsChoiceStr := make([]string, len(m.reroll.Value))
	
	for i, modifier := range m.reroll.Value {
		key := "stat." + string(modifier.Stat)
		label := t.Localize(key)
		buildsChoiceStr[i] = fmt.Sprintf("%s: %d", label, modifier.Value)
	}

	content.WriteString(strings.Join(buildsChoiceStr, ", "))
	content.WriteString(strings.Join([]string{" ", t.Localize("attempt"), ": "}, ""))
	content.WriteString(fmt.Sprintf("%d", m.reroll.Attempt))

	content.WriteString("\n\n")
	content.WriteString(m.reroll.Options.View())

	return content.String()
}
