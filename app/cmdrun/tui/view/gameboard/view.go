package gameboard

import (
	"solopg/app/cmdrun/services/game"
	"solopg/app/cmdrun/tui/view/sidemenu/metadatamenu"
	"solopg/app/cmdrun/types/viewoptions"
	"solopg/app/shared/services/i19n"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

/*
TODO LOW Trouver un meilleur moyen de gerer les sizes,
trouver un moyen de les distibuer al'interieur des sous composant avec
par exemple un contexte disponible, un peu comme Engine et States.
*/
const (
	PANEL_WHIDTH  = 34
	PANEL_GAP     = 1
	ORACLE_HEIGHT = 1
)

func (m model) View() tea.View {
	base := m.viewport.View()

	if !m.isMenuActiveElementOpen() {
		base += "\n" + m.textarea.View()
	}

	var sidemenu = []string{}
	for _, item := range m.menu {
		sidemenu = append(sidemenu, item.GetMenuView())
	}

	composedView := lipgloss.JoinHorizontal(
		lipgloss.Top,
		base,
		strings.Repeat(" ", PANEL_GAP),
		renderSidePanel(m.engine, lipgloss.Height(base), sidemenu...),
	)

	view := tea.NewView(composedView)
	view.Cursor = getCursor(m)
	view.AltScreen = true

	return view
}

func (m model) GetFooter() []string {
	footer := []string{i19n.Localize("cmd.ctrl+s:save")}
	return append(footer, m.getMenuActiveElement().GetFooter()...)
}

// - Viewport - //

func (m *model) refreshViewport(params viewoptions.RefreshOption) (*model, tea.Cmd) {
	content := m.getContent()
	content = lipgloss.NewStyle().
		Width(m.viewport.Width()).
		Render(content)

	m.viewport.SetContent(content)

	if params.Positionreset {
		m.viewport.GotoTop()
	}

	return m, nil
}

func (m *model) getContent() string {
	if m.isMenuActiveElementOpen() {
		return m.getMenuActiveElement().GetView()
	}

	return strings.Join(m.journal, "\n")
}

// - Helper - //
func renderSidePanel(engine *game.Engine, height int, views ...string) string {
	var content strings.Builder

	composeSidePanel(&content, engine)

	statsview := strings.TrimRight(content.String(), "\n")
	menuview := []string{}
	for _, view := range views {
		menuview = append(menuview, view)
	}

	verticalView := lipgloss.JoinVertical(lipgloss.Left, statsview, strings.Join(menuview, "\n"))

	return renderPanel(verticalView, height)
}

func renderPanel(content string, height int) string {
	return lipgloss.NewStyle().
		Width(PANEL_WHIDTH-4).
		Height(max(0, height-2)).
		Padding(0, 1).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("5")).
		Render(content)
}

func composeSidePanel(content *strings.Builder, engine *game.Engine) {
	metadatamenu.GetCharacterInfo(content, engine)
	metadatamenu.GetLocationInfo(content, engine)
	metadatamenu.GetStatInfo(content, engine)
}

func getCursor(m model) *tea.Cursor {
	cursor := m.textarea.Cursor()
	cursor.Y += lipgloss.Height(m.viewport.View())

	if item := m.getMenuActiveElement(); item != nil && item.IsOpen() {
		cursor = nil
	}

	return cursor
}
