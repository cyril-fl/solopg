package gameboard

import (
	"solopg/internal/app/game"
	"solopg/internal/app/tui/models/gameboard/sidemenu/metadatasidemenu"
	"solopg/internal/infrastructure/t"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const (
	panelWidth       = 34
	panelGap         = 1
	oracleMenuHeight = 1
)

func (m model) View() tea.View {
	base := m.viewport.View() + "\n" + m.textarea.View()

	if menu := m.getActiveItem(); menu != nil && menu.IsOpen() {
		base = menu.GetView()
	} 

	var sidemenu = []string{}
	for _, item := range m.menu {
		sidemenu = append(sidemenu, item.GetMenuView())
	}

	composedView := lipgloss.JoinHorizontal(
		lipgloss.Top,
		base,
		strings.Repeat(" ", panelGap),
		renderSidePanel(m.engine, lipgloss.Height(base), sidemenu...),
	)

	view := tea.NewView(composedView)

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
	return append(footer, m.getActiveItem().GetFooter()...)
}

// -- Helper -- //
func renderSidePanel(engine *game.Engine, height int, Titles ...string) string {
	var content strings.Builder

	composeSidePanel(&content, engine)

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

func renderPanel(content string, height int) string {
	return lipgloss.NewStyle().
		Width(panelWidth-4).
		Height(max(0, height-2)).
		Padding(0, 1).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("5")).
		Render(content)
}

func composeSidePanel(content *strings.Builder, engine *game.Engine) {
	metadatasidemenu.GetCharacterInfo(content, engine)
	metadatasidemenu.GetLocationInfo(content, engine)
	metadatasidemenu.GetStatInfo(content, engine)
}
