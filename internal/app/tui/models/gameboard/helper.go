package gameboard

import (
	"fmt"
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
