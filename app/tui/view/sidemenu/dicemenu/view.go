package dicemenu

import (
	"solopg/app/services/i18n"
	"solopg/app/utils/transform"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// - View - //
func (m *DiceMenu) GetMenuView() string {
	title := transform.Uppercase(i18n.Localize(m.id))
	title = lipgloss.NewStyle().Bold(true).Render(title)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		m.list.View(),
	)
}

func (m *DiceMenu) GetView() string {
	return ""
}

func (m *DiceMenu) GetFooter() []string {
	return []string{
		i18n.Localize("shift-enter:roll"),
	}
}

// - Handlers - //
func (m *DiceMenu) HandleWindowResize(msg tea.WindowSizeMsg) tea.Cmd {
	// m.list.SetSize(msg.Width, msg.Height)
	return nil
}
