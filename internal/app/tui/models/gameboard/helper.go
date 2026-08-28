package gameboard

import (
	"solopg/internal/app/tui"
	"solopg/internal/infrastructure/t"
	"strings"
	"time"

	"charm.land/bubbles/v2/cursor"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
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
	m.codexView.page.model.SetWidth(chatWidth)
	m.codexView.page.model.SetHeight(max(0, msg.Height))
	if m.codexView.form.open {
		m.codexView.form.model.SetWidth(chatWidth)
	}
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

// TODO refactor
func handlePanelNavigation(m model, msg tea.KeyPressMsg) (model, tea.Cmd) {
	if m.activeMenu == oracleMenu {
		atStart := m.oracleList.Index() == 0
		atEnd := m.oracleList.Index() >= len(m.oracleList.Items())-1
		if msg.String() == tui.KeyDown && atEnd && len(m.codexList.Items()) > 0 {
			m.activeMenu = codexMenu
			m.codexList.Select(0)
			return m, nil
		}
		if msg.String() == tui.KeyUp && atStart {
			return m, nil
		}
		var cmd tea.Cmd
		m.oracleList, cmd = m.oracleList.Update(msg)
		return m, cmd
	}

	atStart := m.codexList.Index() == 0
	atEnd := m.codexList.Index() >= len(m.codexList.Items())-1
	if msg.String() == tui.KeyUp && atStart && len(m.oracleList.Items()) > 0 {
		m.activeMenu = oracleMenu
		m.oracleList.Select(len(m.oracleList.Items()) - 1)
		return m, nil
	}
	if msg.String() == tui.KeyDown && atEnd {
		return m, nil
	}
	var cmd tea.Cmd
	m.codexList, cmd = m.codexList.Update(msg)
	return m, cmd
}

func handleDefaultInput(m model, msg tea.Msg) (model, tea.Cmd) {
	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func handleSaveInput(m model, msg tui.SaveMsg) (model, tea.Cmd) {
	if msg.Err != nil {
		m.err = msg.Err
		m.messages = append(m.messages, t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "error.save", TemplateData: map[string]any{"Error": msg.Err}}))
	} else {
		log := t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "save.success", TemplateData: map[string]any{"Time": time.Now().Format("2006-01-02 15:04:05")}})
		m.engine.Log(log)
		m.messages = append(m.messages, log)
		m.err = nil
	}

	m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width()).Render(strings.Join(m.messages, "\n")))
	m.viewport.GotoBottom()
	return m, nil
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
