package gameboard

import (
	"solopg/internal/app/game"
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/models/gameboard/sidemenu"
	"solopg/internal/app/tui/models/gameboard/sidemenu/codexmenu"
	"solopg/internal/app/tui/models/gameboard/sidemenu/dicemenu"
	"solopg/internal/app/tui/models/gameboard/sidemenu/oraclemenu"
	"solopg/internal/infrastructure/t"
	"solopg/types/size"
	"strings"
	"time"

	"charm.land/bubbles/v2/cursor"
	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// -- Init -- //
func initTextarea() textarea.Model {
	ta := textarea.New()
	ta.Placeholder = t.Localize("chat.placeholder")
	ta.SetVirtualCursor(false)
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 280

	ta.SetWidth(30)
	ta.SetHeight(3)

	// Remove cursor line styling
	s := ta.Styles()
	s.Focused.CursorLine = lipgloss.NewStyle()
	ta.SetStyles(s)

	ta.ShowLineNumbers = false
	ta.KeyMap.InsertNewline.SetEnabled(false)

	return ta
}

func initViewport() viewport.Model {
	vp := viewport.New(viewport.WithWidth(30), viewport.WithHeight(5))
	// TODO voir pour set autre choses en fonction de message deja present ou non.
	vp.SetContent(t.Localize("chat.welcome"))
	vp.KeyMap.Left.SetEnabled(false)
	vp.KeyMap.Right.SetEnabled(false)

	return vp
}

func initJournal(engine *game.Engine) []string {
	jrnl := make([]string, 0, len(engine.State.Journal.Entries))

	for _, entry := range engine.State.Journal.Entries {
		jrnl = append(jrnl, entry.String())
	}

	return jrnl
}

func initSideMenu(engine *game.Engine) []sidemenu.MenuItem {
	size := size.Size{
		Width:  panelWidth - 4,
		Height: oracleMenuHeight + 5,
	}

	return []sidemenu.MenuItem{
		oraclemenu.NewSideMenu(size, true),
		dicemenu.NewSideMenu(size, false),
		codexmenu.NewSideMenu(codexmenu.CodexMenuParams{
			Size:  size,
			Codex: engine.State.Codex.EnsureInitialized(),
		}, false),
	}
}

func initAuthor(engine *game.Engine) string {
	author := "System"
	if engine.State.Player != nil {
		author = engine.State.Player.Name
	}
	return author
}

// -- Getters & Setters -- //
func (m *model) getActiveItem() sidemenu.MenuItem {
	return m.menu[m.activeMenuIndex]
}

// -- Handlers -- //
func handleEnterInput(m model) (model, tea.Cmd) {
	input := m.textarea.Value()
	if input == "" {
		return m, nil
	}

	msg := m.senderStyle.Render(m.author + ": " + input)
	m.engine.AddJournalEntry(m.author, input)

	m.journal = append(m.journal, msg)

	m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width()).Render(strings.Join(m.journal, "\n")))
	m.textarea.Reset()
	m.viewport.GotoBottom()

	return m, nil
}

func handleSaveInput(m model, msg tui.SaveMsg) (model, tea.Cmd) {
	if msg.Err != nil {
		m.err = msg.Err
		m.journal = append(m.journal, t.Localize("error.save", map[string]any{"Error": msg.Err}))
	} else {
		log := t.Localize("save.success", map[string]any{"Time": time.Now().Format("2006-01-02 15:04:05")})

		m.engine.Log(log)
		m.journal = append(m.journal, log)
		m.err = nil
	}

	m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width()).Render(strings.Join(m.journal, "\n")))
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

func handleMenuDirection(m model, msg sidemenu.Direction) (model, tea.Cmd) {
	previousMenuIndex := m.activeMenuIndex
	menuLength := len(m.menu)

	switch msg {
	case sidemenu.PreviousMenu:
		if m.activeMenuIndex > 0 {
			m.activeMenuIndex--
		}
	case sidemenu.NextMenu:
		if m.activeMenuIndex < menuLength-1 {
			m.activeMenuIndex++
		}
	}

	if previousMenuIndex != m.activeMenuIndex {
		m.menu[previousMenuIndex].SetFocus(false)
		m.menu[m.activeMenuIndex].SetFocus(true)
	}

	return m, nil
}

func handleWindowResize(m *model, msg tea.WindowSizeMsg) {
	chatWidth := max(0, msg.Width-panelWidth-panelGap)

	m.viewport.SetWidth(chatWidth)
	m.textarea.SetWidth(chatWidth)
	m.viewport.SetHeight(max(0, msg.Height-m.textarea.Height()-1))

	if len(m.journal) > 0 {
		m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width()).Render(strings.Join(m.journal, "\n")))
	}
	m.viewport.GotoBottom()
}



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
