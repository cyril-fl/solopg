package tui

import (
	"fmt"
	"strings"

	"solopg/internal/infrastructure/t"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
	if m.err != nil {
		errorLabel := t.Localize("error.label")
		return tea.NewView(alignFooter(fmt.Sprintf("%s: %v", errorLabel, m.err), m.windowHeight(), mainFooter()))
	}

	if current := m.steps.GetCurrentSubmodel(); current != nil {
		view := current.View()
		view.Content = alignFooter(view.Content, m.windowHeight(), mergeFooter(current))
		return view
	}

	return tea.NewView(alignFooter("", m.windowHeight(), mainFooter()))
}

/* 
	TODO
	LOW
	Modifier le footer pour qu'il devienne une "class" en soit.
*/
func mainFooter() []string {
	return []string{t.Localize("quit")}
}

type HasFooter interface {
	GetFooter() []string
}

func mergeFooter(current tea.Model) []string {
	footer := append([]string(nil), mainFooter()...)
	if provider, ok := current.(HasFooter); ok {
		footer = append(footer, provider.GetFooter()...)
	}
	return footer
}

func alignFooter(content string, height int, footerParts []string) string {
	footer := strings.Join(footerParts, " | ")

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

func (m model) windowHeight() int {
	if m.size == nil {
		return 0
	}
	return m.size.Height
}
