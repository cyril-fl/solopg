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
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

func (m model) View() tea.View {
	viewportView := m.viewport.View()
	chatView := viewportView + "\n" + m.textarea.View()
	if m.codexView.page.open {
		chatView = m.codexView.page.model.View()
	}
	if m.codexView.form.open {
		chatView = m.codexView.form.model.View()
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
	if m.codexView.page.open || m.codexView.form.open {
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
	footer := []string{t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "save"})}
	if m.codexView.form.open {
		return append(footer, t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "validate"}), t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "cancel"}))
	}
	if m.codexView.page.open {
		footer = append(footer, t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "add"}), t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "back_to_chat"}))
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
		sideView = lipgloss.JoinVertical(lipgloss.Left, statsView, "\n"+t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "oracles"})+"\n", oracleView, "\n"+t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "codex"})+"\n", codexView)
	}

	return renderPanel(sideView, height)
}

func renderCharacterInfo(content *strings.Builder, engine *game.Engine) {
	var characterName = t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "character"})

	isEngineNil := engine == nil || engine.State == nil

	if isEngineNil || engine.State.Player == nil {
		characterName += t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "unknown_player"})
	} else {
		characterName += engine.State.Player.Name
	}

	content.WriteString(characterName)
	content.WriteString("\n")
}

func renderLocationInfo(content *strings.Builder, engine *game.Engine) {
	var locationName = t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "place"})

	isEngineNil := engine == nil || engine.State == nil

	if isEngineNil || engine.State.CurrentLocation == nil {
		locationName += t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "unknown_place"})
	} else {
		locationName += engine.State.CurrentLocation.Name
	}

	content.WriteString(locationName)
	content.WriteString("\n")
}

func renderStatsInfo(content *strings.Builder, engine *game.Engine) {
	content.WriteString(t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "stats_upper"}) + "\n")

	if engine == nil || engine.State == nil || engine.State.Player == nil {
		content.WriteString(t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "no_stats"}))
	} else {
		for _, stat := range effects.ListStats() {
			key := "stat." + string(stat)
			label := t.Localizer.MustLocalize(&goi18n.LocalizeConfig{
				MessageID: key,
				DefaultMessage: &goi18n.Message{
					ID:    key,
					Other: string(stat),
				},
			})
			fmt.Fprintf(content, "%-10s %d\n", label, engine.State.Player.Stats[stat])
		}
	}
}
