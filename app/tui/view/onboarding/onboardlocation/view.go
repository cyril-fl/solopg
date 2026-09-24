package onboardlocation

import (
	"fmt"
	"solopg/app/components/page"
	"solopg/app/services/t"
	"solopg/app/tui"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	// TODO a un moment autour d'ici il y a yb proble avec la view et un "scrol" se fait
	// le probleme avec dans un rerol, deha ic avec le build
	

	page := page.NewPage(page.Template{
		Title:    "create_character",
		Subtitle: "location",
		Body:     m.getContent(),
	})

	return page.View()
}

func (m model) getContent() string {
	content := strings.Builder{}

	content.WriteString(m.reroll.Value.Name)
	content.WriteString(strings.Join([]string{" ", t.Localize("attempt"), ": "}, ""))
	content.WriteString(fmt.Sprintf("%d", m.reroll.Attempt))

	if m.reroll.Value.Description.Description != tui.EmptyKey {
		content.WriteString("\n")
		content.WriteString(m.reroll.Value.Description.Description)
	}

	content.WriteString("\n\n")
	content.WriteString(m.reroll.Options.View())
	return content.String()
}
