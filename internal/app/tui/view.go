package tui

import tea "charm.land/bubbletea/v2"

const footer = "\n\nCtrl+Q: Quitter"

func (m model) View() tea.View {
	content := ""

	if m.current() != nil {
		content = m.current().View().Content
	}

	return tea.NewView(content + footer)
}
