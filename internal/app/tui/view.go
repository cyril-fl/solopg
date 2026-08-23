package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const footer = "\n\nCtrl+Q: Quitter"

func alignFooter(content string, height int) string {
	if height <= 0 {
		return content + footer
	}

	footerHeight := lipgloss.Height(footer)
	contentHeight := lipgloss.Height(content)
	availableHeight := height - footerHeight
	if contentHeight < availableHeight {
		content += strings.Repeat("\n", availableHeight-contentHeight)
	}

	return content + footer
}

func (m model) View() tea.View {
	if m.err != nil {
		return tea.NewView(alignFooter(fmt.Sprintf("Erreur: %v", m.err), m.windowHeight()))
	}

	if current := m.stepList.getCurrentSubmodel(); current != nil {
		view := current.View()
		view.Content = alignFooter(view.Content, m.windowHeight())
		return view
	}

	return tea.NewView(alignFooter("", m.windowHeight()))
}

func (m model) windowHeight() int {
	if m.size == nil {
		return 0
	}
	return m.size.Height
}
