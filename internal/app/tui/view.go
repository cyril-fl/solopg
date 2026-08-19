package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

const footer = "\n\nCtrl+Q: Quitter"

func (m model) View() tea.View {
	if m.err != nil {
		return tea.NewView(fmt.Sprintf("Erreur: %v%s", m.err, footer))
	}

	if current := m.current(); current != nil {
		view := current.View()
		view.Content += footer
		return view
	}

	return tea.NewView(footer)
}
