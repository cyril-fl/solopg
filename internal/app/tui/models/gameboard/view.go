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
	chatView := viewportView + "\n" + m.textarea.View()
	if m.codex.PageOpen() || m.codex.FormOpen() {
		chatView = m.codex.View()
	}
	if m.showPanel {
		chatView = lipgloss.JoinHorizontal(
			lipgloss.Top,
			chatView,
			strings.Repeat(" ", panelGap),
			renderSidePanel(m.engine, lipgloss.Height(chatView), m.oracleList.View(), m.diceList.View(), m.codex.MenuView()),
		)
	}

	v := tea.NewView(chatView)
	c := m.textarea.Cursor()
	if m.codex.PageOpen() || m.codex.FormOpen() {
		c = nil
	}
	if c != nil {
		c.Y += lipgloss.Height(viewportView)
	}
	v.Cursor = c
	v.AltScreen = true
	return v
}

func (m model) GetFooter() []string {
	footer := []string{t.Localize("save")}
	if m.codex.FormOpen() {
		return append(footer, t.Localize("validate"), t.Localize("cancel"))
	}
	if m.codex.PageOpen() {
		footer = append(footer, t.Localize("add"), t.Localize("back_to_chat"))
	}
	return footer
}

func renderSidePanel(engine *game.Engine, height int, oracleView, diceView, codexView string) string {
	var content strings.Builder

	renderCharacterInfo(&content, engine)
	renderLocationInfo(&content, engine)
	renderStatsInfo(&content, engine)

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

func renderCharacterInfo(content *strings.Builder, engine *game.Engine) {
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

func renderLocationInfo(content *strings.Builder, engine *game.Engine) {
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

func renderStatsInfo(content *strings.Builder, engine *game.Engine) {
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
