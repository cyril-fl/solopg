package hintmenu

import (
	"solopg/app/shared/services/i19n"
	"solopg/app/shared/utils/transform"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// - View - //
func (m *HintMenu) GetMenuView() string {
	title := transform.Uppercase(i19n.Localize(m.id))
	title = lipgloss.NewStyle().Bold(true).Render(title)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		m.list.View(),
	)
}

func (m *HintMenu) GetView() string {
	return ""
}

func (m *HintMenu) GetFooter() []string {
	return []string{
		i19n.Localize("cmd.shift+enter:roll"),
	}
}

// - Handlers --//
func (m *HintMenu) HandleWindowResize(msg tea.WindowSizeMsg) tea.Cmd {
	// m.list.SetSize(msg.Width, msg.Height)
	return nil
}
