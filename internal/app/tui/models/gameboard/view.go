package gameboard

import (
	"fmt"
	"solopg/internal/app/game"
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/effects"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
	viewportView := m.viewport.View()
	chatView := viewportView + "\n" + m.textarea.View()
	if m.showPanel {
		chatView = lipgloss.JoinHorizontal(
			lipgloss.Top,
			chatView,
			strings.Repeat(" ", panelGap),
			renderSidePanel(m.engine, lipgloss.Height(chatView), m.oracleList.View()),
		)
	}

	v := tea.NewView(chatView)
	c := m.textarea.Cursor()
	if c != nil {
		c.Y += lipgloss.Height(viewportView)
	}
	v.Cursor = c
	v.AltScreen = true
	return v
}

func (m model) GetFooter() []string {
	return []string{"Ctrl+S: Sauvegarder"}
}

func renderSidePanel(engine *game.Engine, height int, oracleView string) string {
	var content strings.Builder

	renderCharacterInfo(&content, engine)
	renderLocationInfo(&content, engine)
	renderStatsInfo(&content, engine)

	return composeSidePanel(&content, height, oracleView)
}

func composeSidePanel(content *strings.Builder, height int, oracleView string) string {
	renderPanel := func(content string, height int) string {
		return lipgloss.NewStyle().
			Width(panelWidth - 4).
			Height(max(0, height)).
			Padding(0, 1).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("5")).
			Render(content)
	}

	statsView := strings.TrimRight(content.String(), "\n")
	sideView := statsView

	if oracleView != tui.EmptyKey {
		sideView = lipgloss.JoinVertical(lipgloss.Left, statsView, "\nORACLES\n", oracleView)
	}

	return renderPanel(sideView,height)
}

func renderCharacterInfo(content *strings.Builder, engine *game.Engine)  {
	var characterName = "PERSONNAGE: "

	isEngineNil := engine == nil || engine.State == nil

	if isEngineNil || engine.State.Player == nil {
		characterName += "Joueur inconnu"
	} else {
		characterName += engine.State.Player.Name
	}

	content.WriteString(characterName)
	content.WriteString("\n")
}

func renderLocationInfo(content *strings.Builder, engine *game.Engine)  {
	var locationName = "LIEU: "

	isEngineNil := engine == nil || engine.State == nil

	if isEngineNil || engine.State.CurrentLocation == nil {
		locationName += "Lieu inconnu"
	} else {
		locationName += engine.State.CurrentLocation.Name
	}

	content.WriteString(locationName)
	content.WriteString("\n")
}

func renderStatsInfo(content *strings.Builder, engine *game.Engine)  {
	content.WriteString("STATS:\n")

	if engine.State.Player == nil {
		content.WriteString("Aucune statistique")
	} else {
		for _, stat := range effects.ListStats() {
			fmt.Fprintf(content, "%-10s %d\n", stat, engine.State.Player.Stats[stat])
		}
	}
}	


