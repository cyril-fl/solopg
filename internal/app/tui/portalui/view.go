package portalui

import (
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	content := strings.Builder{}

	content.WriteString("Location: ")
	content.WriteString(m.drawLocations.Name)
	if m.drawLocations.Description.Description != "" {
		content.WriteString("\n")
		content.WriteString(m.drawLocations.Description.Description)
	}
	content.WriteString("\n\n")
	content.WriteString(m.choiceList.View())

	return tea.NewView(content.String())
}
