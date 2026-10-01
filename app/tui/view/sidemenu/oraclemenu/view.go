package oraclemenu

import (
	"solopg/app/services/i19n"
	"solopg/app/utils/transform"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// - View - //
func (m *OracleMenu) GetMenuView() string {
	title := transform.Uppercase(i19n.Localize(m.id))
	title = lipgloss.NewStyle().Bold(true).Render(title)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		m.list.View(),
	)
}

func (m *OracleMenu) GetView() string {
	return ""
}

func (m *OracleMenu) GetFooter() []string {
	return []string{
		i19n.Localize("cmd.shift+enter:roll"),
	}
}

// - Handlers --//
func (m *OracleMenu) HandleWindowResize(msg tea.WindowSizeMsg) tea.Cmd {
	// m.list.SetSize(msg.Width, msg.Height)
	return nil
}
