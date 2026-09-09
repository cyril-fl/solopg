package gameboard

import (
	"solopg/internal/app/tui"
	"solopg/internal/infrastructure/t"
	"strings"
	"time"

	"charm.land/bubbles/v2/cursor"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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
		m.messages = append(m.messages, t.Localize("error.save", map[string]any{"Error": msg.Err}))
	} else {
		log := t.Localize("save.success", map[string]any{"Time": time.Now().Format("2006-01-02 15:04:05")})

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
	// if msg.Width >= panelWidth+panelGap+minimumChatWidth {
	// 	// m.showPanel = true
	chatWidth = msg.Width - panelWidth - panelGap
	// } else {
	// 	// m.showPanel = false
	// }

	m.viewport.SetWidth(chatWidth)
	m.textarea.SetWidth(chatWidth)
	// m.oracleMenu.SetSize(panelWidth-4, oracleMenuHeight)
	// m.diceMenu.SetSize(panelWidth-4, diceMenuHeight)
	// m.codexMenu.SetMenuSize(panelWidth-4, codexMenuHeight)
	// m.codexMenu.SetSize(chatWidth, max(0, msg.Height))
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

// func handlePanelNavigation(m model, msg tea.KeyPressMsg) (model, tea.Cmd) {
	//TODO trouver un moyen pour que ce soit en fonction de curent menu et que si on est a la fin ou au debut cuurent menu change.
	// TODO s'occuper de ça en suite !

// 	if m.activeMenu == oracleMenu {
// 		atStart := m.oracleMenu.Index() == 0
// 		atEnd := m.oracleMenu.Index() >= len(m.oracleMenu.Items())-1
// 		if msg.String() == tui.KeyDown && atEnd && len(m.diceMenu.Items()) > 0 {
// 			m.activeMenu = diceMenu
// 			m.diceMenu.Select(0)
// 			return m, nil
// 		}
// 		if msg.String() == tui.KeyUp && atStart {
// 			return m, nil
// 		}
// 		var cmd tea.Cmd
// 		m.oracleMenu, cmd = m.oracleMenu.Update(msg)
// 		return m, cmd
// 	}

// 	if m.activeMenu == diceMenu {
// 		atStart := m.diceMenu.Index() == 0
// 		atEnd := m.diceMenu.Index() >= len(m.diceMenu.Items())-1
// 		if msg.String() == tui.KeyUp && atStart && len(m.oracleMenu.Items()) > 0 {
// 			m.activeMenu = oracleMenu
// 			m.oracleMenu.Select(len(m.oracleMenu.Items()) - 1)
// 			return m, nil
// 		}
// 		// if msg.String() == tui.KeyDown && atEnd && m.codexMenu.MenuItemsCount() > 0 {
// 		// 	m.activeMenu = codexMenu
// 		// 	m.codexMenu.SelectMenu(0)
// 		// 	return m, nil
// 		// }
// 		var cmd tea.Cmd
// 		m.diceMenu, cmd = m.diceMenu.Update(msg)
// 		return m, cmd
// 	}

// 	// atStart := m.codexMenu.MenuIndex() == 0
// 	// atEnd := m.codexMenu.MenuIndex() >= m.codexMenu.MenuItemsCount()-1
// 	if msg.String() == tui.KeyUp && atStart && len(m.oracleMenu.Items()) > 0 {
// 		m.activeMenu = diceMenu
// 		m.diceMenu.Select(len(m.diceMenu.Items()) - 1)
// 		return m, nil
// 	}
// 	if msg.String() == tui.KeyDown && atEnd {
// 		return m, nil
// 	}
// 	return m, m.codexMenu.UpdateMenu(msg)
// }

// func handleDiceRoll(m model) (model, tea.Cmd) {
// 	selected, ok := m.diceMenu.SelectedItem().(tui.Item[gameplay.Dice])
// 	if !ok {
// 		return m, nil
// 	}

// 	dice := selected.Value()

// 	result := dice.Roll()

// 	message := fmt.Sprintf("%s : %d", dice.GetName(), result)

// 	m.engine.AddJournalEntry("Dice", message)
// 	m.messages = append(m.messages, message)

// 	m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width()).Render(strings.Join(m.messages, "\n")))
// 	m.viewport.GotoBottom()

// 	return m, nil
// }
