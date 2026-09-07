package gameboard

import (
	"fmt"
	"solopg/internal/app/game"
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/effects"
	"solopg/internal/infrastructure/t"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
	viewportView := m.viewport.View()
	baseView := viewportView + "\n" + m.textarea.View()

	// TODO remplacer par un ShowView()
	// if m.codexMenu.PageOpen() || m.codexMenu.FormOpen() {
	// 	chatView = m.codexMenu.View()
	// }

	// if m.codexMenu.PageOpen() {
	// 	baseView = m.codexMenu.View()
	// }

	// if m.showPanel {
	baseView = lipgloss.JoinHorizontal(
		lipgloss.Top,
		baseView,
		strings.Repeat(" ", panelGap),
		renderSidePanel(m.engine, lipgloss.Height(baseView), m.oracleMenu.View(), m.diceMenu.View(), m.codexMenu.MenuView()),
	)
	// }

	view := tea.NewView(baseView)
	cursor := m.textarea.Cursor()

	// if m.codexMenu.PageOpen() || m.codexMenu.FormOpen() {
	// 	c = nil
	// }

	if m.codexMenu.PageOpen() {
		cursor = nil
	}

	if cursor != nil {
		cursor.Y += lipgloss.Height(viewportView)
	}

	view.Cursor = cursor
	view.AltScreen = true
	return view
}

func (m model) GetFooter() []string {
	footer := []string{t.Localize("save")}
	// if m.codexMenu.FormOpen() {
	// 	return append(footer, t.Localize("validate"), t.Localize("cancel"))
	// }
	// if m.codexMenu.PageOpen() {
	// 	footer = append(footer, t.Localize("add"), t.Localize("back_to_chat"))
	// }
	footer = append(footer, m.activeMenuItem.GetFooter()...)
	return footer
}

func renderSidePanel(engine *game.Engine, height int, oracleView, diceView, codexView string) string {
	var content strings.Builder

	getComponentCharacterInfo(&content, engine)
	getComponentLocationInfo(&content, engine)
	getComponentStatInfo(&content, engine)

	return composeSidePanel(&content, height, oracleView, diceView, codexView)
}

func composeSidePanel(content *strings.Builder, height int, oracleView, diceView, codexView string) string {
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

	if oracleView != tui.EmptyKey {
		sideView = lipgloss.JoinVertical(lipgloss.Left, statsView,
			"\n"+t.Localize("oracles")+"\n", oracleView,
			"\n"+t.Localize("dice")+"\n", diceView,
			"\n"+t.Localize("codex")+"\n", codexView)
	}

	return renderPanel(sideView, height)
}

func getComponentCharacterInfo(content *strings.Builder, engine *game.Engine) {
	var characterName = t.Localize("character")

	isEngineNil := engine == nil || engine.State == nil

	if isEngineNil || engine.State.Player == nil {
		characterName += t.Localize("unknown_player")
	} else {
		characterName += engine.State.Player.Name
	}

	content.WriteString(characterName)
	content.WriteString("\n")
}

func getComponentLocationInfo(content *strings.Builder, engine *game.Engine) {
	var locationName = t.Localize("place")

	isEngineNil := engine == nil || engine.State == nil

	if isEngineNil || engine.State.CurrentLocation == nil {
		locationName += t.Localize("unknown_place")
	} else {
		locationName += engine.State.CurrentLocation.Name
	}

	content.WriteString(locationName)
	content.WriteString("\n")
}

func getComponentStatInfo(content *strings.Builder, engine *game.Engine) {
	content.WriteString(t.Localize("stats_upper") + "\n")

	if engine == nil || engine.State == nil || engine.State.Player == nil {
		content.WriteString(t.Localize("no_stats"))
	} else {
		for _, stat := range effects.ListStats() {
			key := "stat." + string(stat)
			label := t.Localize(key)
			fmt.Fprintf(content, "%-10s %d\n", label, engine.State.Player.Stats[stat])
		}
	}
}
