package gameui

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	footerStyleMuted = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	settingsStyle    = lipgloss.NewStyle()
)

func (root model) View() tea.View {
	switch root.screen {
	case screenSettings:
		return tea.NewView(root.settingsView())
	default:
		return tea.NewView(root.mainView())
	}
}

func (root model) mainView() string {
	var builder strings.Builder
	builder.WriteString(root.itemList.View())
	builder.WriteString(root.footer())
	return builder.String()
}

func (root model) footer() string {
	line := "q: quitter  s: settings"
	if root.width <= 0 {
		return "\n" + footerStyleMuted.Render(line)
	}
	return "\n" + footerStyleMuted.Width(root.width).Render(line)
}

func (root model) settingsView() string {
	body := "Réglages\n\n(b ou échap : retour, q : quitter)"
	if root.width <= 0 {
		return body
	}
	return settingsStyle.Width(root.width).Render(body)
}
