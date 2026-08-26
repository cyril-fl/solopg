package gameboard

import (
	"fmt"
	"solopg/internal/app/game"
	"solopg/internal/app/tui"
	"solopg/internal/domain/gameplay"
	"strings"
	"time"

	"charm.land/bubbles/v2/cursor"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func handleWindowResize(m *model, msg tea.WindowSizeMsg) {
	chatWidth := msg.Width
	if msg.Width >= panelWidth+panelGap+minimumChatWidth {
		m.showPanel = true
		chatWidth = msg.Width - panelWidth - panelGap
	} else {
		m.showPanel = false
	}

	m.viewport.SetWidth(chatWidth)
	m.textarea.SetWidth(chatWidth)
	m.oracleList.SetSize(panelWidth-4, oracleMenuHeight)
	m.codexList.SetSize(panelWidth-4, codexMenuHeight)
	m.codexPage.SetWidth(chatWidth)
	m.codexPage.SetHeight(max(0, msg.Height))
	// Reserve the input and the separator above it. The parent TUI already
	// reserved the footer height before forwarding the window size.
	m.viewport.SetHeight(max(0, msg.Height-m.textarea.Height()-1))

	if len(m.messages) > 0 {
		// Wrap content before setting it.
		m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width()).Render(strings.Join(m.messages, "\n")))
	}
	m.viewport.GotoBottom()
}

const (
	panelWidth       = 34
	panelGap         = 1
	minimumChatWidth = 40
	oracleMenuHeight = 1
	codexMenuHeight  = 6
)

func makeOracleModel() list.Model {
	items := make([]list.Item, 0)
	for _, oracle := range gameplay.GetOracle() {
		if oracle.Visible {
			items = append(items, tui.NewItem(oracle.ID, "", oracle))
		}
	}

	model := list.New(items, list.NewDefaultDelegate(), panelWidth-4, oracleMenuHeight)
	tui.ConfigureList(&model)
	return model
}

type codexLink struct {
	name string
	kind codexKind
}

type codexKind uint8

const (
	codexNPCs codexKind = iota
	codexMonsters
	codexLocations
	codexObjects
	codexObjectifs
)

func makeCodexModel() list.Model {
	links := []codexLink{
		{name: "PNJ", kind: codexNPCs},
		{name: "Monstres", kind: codexMonsters},
		{name: "Lieux", kind: codexLocations},
		{name: "Objets", kind: codexObjects},
		{name: "Objectifs", kind: codexObjectifs},
	}
	items := make([]list.Item, 0, len(links))
	for _, link := range links {
		items = append(items, tui.NewItem(link.name, "", link))
	}

	model := list.New(items, list.NewDefaultDelegate(), panelWidth-4, codexMenuHeight)
	tui.ConfigureList(&model)
	return model
}

func handlePanelNavigation(m model, msg tea.KeyPressMsg) (model, tea.Cmd) {
	if m.activeMenu == oracleMenu {
		atStart := m.oracleList.Index() == 0
		atEnd := m.oracleList.Index() >= len(m.oracleList.Items())-1
		if msg.String() == "down" && atEnd && len(m.codexList.Items()) > 0 {
			m.activeMenu = codexMenu
			m.codexList.Select(0)
			return m, nil
		}
		if msg.String() == "up" && atStart {
			return m, nil
		}
		var cmd tea.Cmd
		m.oracleList, cmd = m.oracleList.Update(msg)
		return m, cmd
	}

	atStart := m.codexList.Index() == 0
	atEnd := m.codexList.Index() >= len(m.codexList.Items())-1
	if msg.String() == "up" && atStart && len(m.oracleList.Items()) > 0 {
		m.activeMenu = oracleMenu
		m.oracleList.Select(len(m.oracleList.Items()) - 1)
		return m, nil
	}
	if msg.String() == "down" && atEnd {
		return m, nil
	}
	var cmd tea.Cmd
	m.codexList, cmd = m.codexList.Update(msg)
	return m, cmd
}

func openCodexPage(m model) (model, tea.Cmd) {
	selected, ok := m.codexList.SelectedItem().(tui.Item[codexLink])
	if !ok {
		return m, nil
	}

	link := selected.Value()
	m.pageOpen = true
	m.codexPage.SetWidth(m.viewport.Width())
	// The Codex page gets one additional footer line (Escape: retour au chat).
	// Reserve that line immediately; the next resize event will recalculate it.
	m.codexPage.SetHeight(m.viewport.Height() + m.textarea.Height())
	m.codexPage.SetContent(renderCodexPage(m.engine, link))
	m.codexPage.GotoTop()
	return m, nil
}

func renderCodexPage(engine *game.Engine, link codexLink) string {
	title := link.name
	var entries []string
	if engine != nil && engine.State != nil && engine.State.Codex != nil {
		codexData := engine.State.Codex.EnsureInitialized()
		switch link.kind {
		case codexNPCs:
			entries = codexData.NpcsTable.Summaries()
		case codexMonsters:
			entries = codexData.MonstersTable.Summaries()
		case codexLocations:
			entries = codexData.LocationsTable.Summaries()
		case codexObjects:
			entries = codexData.ObjectsTable.Summaries()
		case codexObjectifs:
			entries = codexData.ObjectifsTable.Summaries()
		}
	}
	if len(entries) == 0 {
		entries = []string{"Aucune entrée dans cette table."}
	}
	return title + "\n\n" + strings.Join(entries, "\n\n")
}

func handleCursorBlink(m model, msg cursor.BlinkMsg) (model, tea.Cmd) {
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func handleEnterInput(m model) (model, tea.Cmd) {
	input := m.textarea.Value()
	if input == "" {
		return m, nil
	}

	msg := m.senderStyle.Render(m.author + ": " + input)
	m.engine.AddJournalEntry(m.author, input)
	m.messages = append(m.messages, msg)

	m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width()).Render(strings.Join(m.messages, "\n")))
	m.textarea.Reset()
	m.viewport.GotoBottom()

	return m, nil
}

func handleOracleRoll(m model) (model, tea.Cmd) {
	selected, ok := m.oracleList.SelectedItem().(tui.Item[*gameplay.Oracle])
	if !ok || selected.Value() == nil {
		return m, nil
	}

	oracle := selected.Value()
	result, err := gameplay.RollOracle[any](oracle)
	if err != nil {
		m.messages = append(m.messages, "Erreur oracle: "+err.Error())
	} else {
		critical := ""
		if result.Critical {
			critical = " (critique)"
		}
		message := fmt.Sprintf("Oracle %s — jet de %d : %v%s", oracle.ID, result.Roll, result.Result, critical)
		m.engine.AddJournalEntry("Oracle", message)
		m.messages = append(m.messages, message)
	}

	m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width()).Render(strings.Join(m.messages, "\n")))
	m.viewport.GotoBottom()
	return m, nil
}

func handleDefaultInput(m model, msg tea.Msg) (model, tea.Cmd) {
	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func handleSaveInput(m model, msg tui.SaveMsg) (model, tea.Cmd) {
	if msg.Err != nil {
		m.err = msg.Err
		m.messages = append(m.messages, "Erreur de sauvegarde: "+msg.Err.Error())
	} else {
		log := fmt.Sprintf("Saved successfully at %s", time.Now().Format("2006-01-02 15:04:05"))
		m.engine.Log(log)
		m.messages = append(m.messages, log)
		m.err = nil
	}

	m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width()).Render(strings.Join(m.messages, "\n")))
	m.viewport.GotoBottom()
	return m, nil
}
