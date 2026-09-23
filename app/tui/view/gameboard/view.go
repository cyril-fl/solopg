package gameboard

import (
	"solopg/app/services/game"
	"solopg/app/services/t"
	"solopg/app/tui/view/sidemenu/metadatamenu"
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
		strings.Repeat(" ", panelGap),
		renderSidePanel(m.engine, lipgloss.Height(base), sidemenu...),
	)

	view := tea.NewView(composedView)
	view.Cursor = getCursor(m)
	view.AltScreen = true

	return view
}

func (m model) GetFooter() []string {
	footer := []string{t.Localize("save")}
	return append(footer, m.getMenuActiveElement().GetFooter()...)
}

// - Viewport - //
func (m *model) refreshViewport(resetPosition bool) (*model, tea.Cmd) {
	content := m.getContent()
	content = lipgloss.NewStyle().
		Width(m.viewport.Width()).
		Render(content)

	m.viewport.SetContent(content)

	if resetPosition {
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
