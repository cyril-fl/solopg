package gameboard

import (
	"solopg/internal/app/game"
	"solopg/internal/app/tui/models/gameboard/sidemenu/metadatasidemenu"
	"solopg/internal/infrastructure/t"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
	viewportView := m.viewport.View()
	baseView := viewportView + "\n" + m.textarea.View()

	if m.activeMenuItem.IsOpen() {
		baseView = m.activeMenuItem.GetView()
	} else {
		var menuTitles = []string{}
		for _, menuItem := range m.menu {
			menuTitles = append(menuTitles, menuItem.ID())
		}

		baseView = lipgloss.JoinHorizontal(
			lipgloss.Top,
			baseView,
			strings.Repeat(" ", panelGap),
			renderSidePanel(m.engine, lipgloss.Height(baseView), menuTitles...),
		)
	}

	view := tea.NewView(baseView)
	// cursor := m.textarea.Cursor()

	// if m.codexMenu.PageOpen() || m.codexMenu.FormOpen() {
	// 	c = nil
	// }

	// if m.codexMenu.PageOpen() {
	// 	cursor = nil
	// }

	// if cursor != nil {
	// 	cursor.Y += lipgloss.Height(viewportView)
	// }

	// view.Cursor = cursor
	// view.AltScreen = true
	return view
}

func (m model) GetFooter() []string {
	footer := []string{t.Localize("save")}
	return append(footer, m.activeMenuItem.GetFooter()...)
}

func renderSidePanel(engine *game.Engine, height int, Titles ...string) string {
	var content strings.Builder

	composeSidePanel(&content, engine)

	renderPanel := func(content string, height int) string {
		return lipgloss.NewStyle().
			Width(panelWidth-4).
			Height(max(0, height-2)).
			Padding(0, 1).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("5")).
			Render(content)
	}

	statsView := strings.TrimRight(content.String(), "\n")
	sideView := statsView

	if len(Titles) > 0 {
		sideView = lipgloss.JoinVertical(lipgloss.Left, statsView)
		for i, title := range Titles {
			if i < len(Titles) {
				sideView = lipgloss.JoinVertical(lipgloss.Left, sideView, "\n"+t.Localize(title))
			}
		}
	}

	return renderPanel(sideView, height)
}

func composeSidePanel(content *strings.Builder, engine *game.Engine) {
	metadatasidemenu.GetCharacterInfo(content, engine)
	metadatasidemenu.GetLocationInfo(content, engine)
	metadatasidemenu.GetStatInfo(content, engine)
}
