package portal

import (
	"fmt"
	"solopg/internal/app/tui"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	content := strings.Builder{}

	content.WriteString("Location: ")
	content.WriteString(m.drawLocations.Name)
	content.WriteString(" Attempt: ")
	content.WriteString(fmt.Sprintf("%d", m.attempt))

	if m.drawLocations.Description.Description != tui.EmptyKey {
		content.WriteString("\n")
		content.WriteString(m.drawLocations.Description.Description)
	}
	content.WriteString("\n\n")
	content.WriteString(m.choiceList.View())

	return tea.NewView(content.String())
}
