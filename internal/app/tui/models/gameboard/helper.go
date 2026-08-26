package gameboard

import (
	"fmt"
	"solopg/internal/app/tui"
	"strings"
	"time"

	"charm.land/bubbles/v2/cursor"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func handleWindowResize(m *model, msg tea.WindowSizeMsg) {
	m.viewport.SetWidth(msg.Width)
	m.textarea.SetWidth(msg.Width)
	// Reserve the input and the separator above it. The parent TUI already
	// reserved the footer height before forwarding the window size.
	m.viewport.SetHeight(max(0, msg.Height-m.textarea.Height()-1))

	if len(m.messages) > 0 {
		// Wrap content before setting it.
		m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width()).Render(strings.Join(m.messages, "\n")))
	}
	m.viewport.GotoBottom()
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
