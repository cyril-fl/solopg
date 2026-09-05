package gameboard

import (
	"fmt"
	"solopg/internal/app/tui"
	"solopg/internal/domain/gameplay"
	"solopg/internal/infrastructure/t"
	"strings"
	"time"

	"charm.land/bubbles/v2/cursor"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

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

func handleSaveInput(m model, msg tui.SaveMsg) (model, tea.Cmd) {
	if msg.Err != nil {
		m.err = msg.Err
		m.messages = append(m.messages, t.Local.MustLocalize(&goi18n.LocalizeConfig{MessageID: "error.save", TemplateData: map[string]any{"Error": msg.Err}}))
	} else {
		log := t.Local.MustLocalize(&goi18n.LocalizeConfig{MessageID: "save.success", TemplateData: map[string]any{"Time": time.Now().Format("2006-01-02 15:04:05")}})

		m.engine.Log(log)
		m.messages = append(m.messages, log)
		m.err = nil
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

func handleCursorBlink(m model, msg cursor.BlinkMsg) (model, tea.Cmd) {
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

// TODO ----  refactor ----
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
	m.diceList.SetSize(panelWidth-4, diceMenuHeight)
	m.codex.SetMenuSize(panelWidth-4, codexMenuHeight)
	m.codex.SetSize(chatWidth, max(0, msg.Height))
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
	diceMenuHeight   = 7
	codexMenuHeight  = 6
)

func handlePanelNavigation(m model, msg tea.KeyPressMsg) (model, tea.Cmd) {
	if m.activeMenu == oracleMenu {
		atStart := m.oracleList.Index() == 0
		atEnd := m.oracleList.Index() >= len(m.oracleList.Items())-1
		if msg.String() == tui.KeyDown && atEnd && len(m.diceList.Items()) > 0 {
			m.activeMenu = diceMenu
			m.diceList.Select(0)
			return m, nil
		}
		if msg.String() == tui.KeyUp && atStart {
			return m, nil
		}
		var cmd tea.Cmd
		m.oracleList, cmd = m.oracleList.Update(msg)
		return m, cmd
	}

	if m.activeMenu == diceMenu {
		atStart := m.diceList.Index() == 0
		atEnd := m.diceList.Index() >= len(m.diceList.Items())-1
		if msg.String() == tui.KeyUp && atStart && len(m.oracleList.Items()) > 0 {
			m.activeMenu = oracleMenu
			m.oracleList.Select(len(m.oracleList.Items()) - 1)
			return m, nil
		}
		if msg.String() == tui.KeyDown && atEnd && m.codex.MenuItemsCount() > 0 {
			m.activeMenu = codexMenu
			m.codex.SelectMenu(0)
			return m, nil
		}
		var cmd tea.Cmd
		m.diceList, cmd = m.diceList.Update(msg)
		return m, cmd
	}

	atStart := m.codex.MenuIndex() == 0
	atEnd := m.codex.MenuIndex() >= m.codex.MenuItemsCount()-1
	if msg.String() == tui.KeyUp && atStart && len(m.oracleList.Items()) > 0 {
		m.activeMenu = diceMenu
		m.diceList.Select(len(m.diceList.Items()) - 1)
		return m, nil
	}
	if msg.String() == tui.KeyDown && atEnd {
		return m, nil
	}
	return m, m.codex.UpdateMenu(msg)
}

func handleDiceRoll(m model) (model, tea.Cmd) {
	selected, ok := m.diceList.SelectedItem().(tui.Item[gameplay.Dice])
	if !ok {
		return m, nil
	}

	dice := selected.Value()

	result := dice.Roll()

	message := fmt.Sprintf("%s : %d", dice.GetName(), result)

	m.engine.AddJournalEntry("Dice", message)
	m.messages = append(m.messages, message)

	m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width()).Render(strings.Join(m.messages, "\n")))
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
		m.messages = append(m.messages, t.Local.MustLocalize(&goi18n.LocalizeConfig{MessageID: "error.oracle_action", TemplateData: map[string]any{"Error": err}}))
	} else {
		critical := ""
		if result.Critical {
			critical = " (" + t.Local.MustLocalize(&goi18n.LocalizeConfig{MessageID: "critical"}) + ")"
		}
		message := fmt.Sprintf("Oracle %s — jet de %d : %v%s", oracle.ID, result.Roll, result.Result, critical)
		m.engine.AddJournalEntry("Oracle", message)
		m.messages = append(m.messages, message)
	}

	m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width()).Render(strings.Join(m.messages, "\n")))
	m.viewport.GotoBottom()

	return m, nil
}
