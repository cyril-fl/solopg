package portalui

import (
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	content := strings.Builder{}

	content.WriteString(m.selectedLocation.Name)
	content.WriteString("\n\n")
	content.WriteString(m.choiceList.View())
	content.WriteString("\n\nq: quit  enter: select")

	return tea.NewView(content.String())
}
