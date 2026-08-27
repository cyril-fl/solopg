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
	if m.pageOpen {
		chatView = m.codexPage.View()
	}
	if m.formOpen {
		chatView = m.form.View()
	}
	if m.showPanel {
		chatView = lipgloss.JoinHorizontal(
			lipgloss.Top,
			chatView,
			strings.Repeat(" ", panelGap),
			renderSidePanel(m.engine, lipgloss.Height(chatView), m.oracleList.View(), m.codexList.View()),
		)
	}

	v := tea.NewView(chatView)
	c := m.textarea.Cursor()
	if m.pageOpen || m.formOpen {
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
	footer := []string{"Ctrl+S: Sauvegarder"}
	if m.formOpen {
		return append(footer, "Entrée: Valider", "Échap: Annuler")
	}
	if m.pageOpen {
		footer = append(footer, "Ctrl+N: Ajouter", "Échap: retour au chat")
	}
	return footer
}

func renderSidePanel(engine *game.Engine, height int, oracleView, codexView string) string {
	var content strings.Builder

	renderCharacterInfo(&content, engine)
	renderLocationInfo(&content, engine)
	renderStatsInfo(&content, engine)

	return composeSidePanel(&content, height, oracleView, codexView)
}

func composeSidePanel(content *strings.Builder, height int, oracleView, codexView string) string {
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
		sideView = lipgloss.JoinVertical(lipgloss.Left, statsView, "\nORACLES\n", oracleView, "\nCODEX\n", codexView)
	}

	return renderPanel(sideView, height)
}

func renderCharacterInfo(content *strings.Builder, engine *game.Engine) {
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

func renderLocationInfo(content *strings.Builder, engine *game.Engine) {
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

func renderStatsInfo(content *strings.Builder, engine *game.Engine) {
	content.WriteString("STATS:\n")

	if engine == nil || engine.State == nil || engine.State.Player == nil {
		content.WriteString("Aucune statistique")
	} else {
		for _, stat := range effects.ListStats() {
			fmt.Fprintf(content, "%-10s %d\n", stat, engine.State.Player.Stats[stat])
		}
	}
}
