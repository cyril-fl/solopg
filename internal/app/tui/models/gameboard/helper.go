package gameboard

import (
	"fmt"
	"solopg/internal/app/game"
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/models/gameboard/sidemenu"
	"solopg/internal/app/tui/models/gameboard/sidemenu/codexmenu"
	"solopg/internal/app/tui/models/gameboard/sidemenu/dicemenu"
	"solopg/internal/app/tui/models/gameboard/sidemenu/oraclemenu"
	"solopg/internal/infrastructure/t"
	"solopg/types/direction"
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

func initViewport(content string) viewport.Model {
	vp := viewport.New(viewport.WithWidth(30), viewport.WithHeight(5))

	vp.SetContent(content)
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
func (m *model) getMenuActiveElement() sidemenu.MenuItem {
	return m.menu[m.activeMenuIndex]
}

func (m *model) isLastMenuElement() bool {
	return m.activeMenuIndex == len(m.menu)-1
}

func (m *model) isFirstMenuElement() bool {
	return m.activeMenuIndex == 0
}

func (m *model) isMenuActiveElementOpen() bool {
	return m.getMenuActiveElement().IsOpen()
}

// -- Handlers -- //
// Input
func (m *model) handleEnterInput() {
	input := m.textarea.Value()
	if input == "" {
		return
	}

	msg := m.senderStyle.Render(m.author + ": " + input)
	m.engine.AddJournalEntry(m.author, input)

	m.journal = append(m.journal, msg)

	m.mainview.SetContent(lipgloss.NewStyle().Width(m.mainview.Width()).Render(strings.Join(m.journal, "\n")))
	m.textarea.Reset()
	m.mainview.GotoBottom()
}

func (m *model) handleSaveInput(msg tui.SaveMsg) (*model, tea.Cmd) {
	if msg.Err != nil {
		m.err = msg.Err
		m.journal = append(m.journal, t.Localize("error.save", map[string]any{"Error": msg.Err}))
	} else {
		log := t.Localize("save.success", map[string]any{"Time": time.Now().Format("2006-01-02 15:04:05")})

		m.engine.Log(log)
		m.journal = append(m.journal, log)
		m.err = nil
	}

	m.mainview.SetContent(lipgloss.NewStyle().Width(m.mainview.Width()).Render(strings.Join(m.journal, "\n")))
	m.mainview.GotoBottom()

	m.refreshViewport(true)

	return m, nil
}

func handleDefaultInput(m *model, msg tea.Msg) (*model, tea.Cmd) {
	var cmd tea.Cmd

	// TODO ajouter la fonction du side menu ici
	// sideMenu := m.getMenuActiveElement()

	// fmt.Println("Side menu is open %s, forwarding message to side menu", sideMenu.IsOpen())

	// if  {
	// 	// m.pageview, cmd = m.pageview.Update(msg)
	// 	return m, cmd
	// }

	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func (m *model) handleCursorBlink(msg cursor.BlinkMsg) (*model, tea.Cmd) {
	var cmd tea.Cmd
	m.mainview, cmd = m.mainview.Update(msg)
	return m, cmd
}

func (m *model) handleKeyPress(msg tea.KeyPressMsg) {
	switch msg.String() {
	case tui.KeyEnter:
		m.handleEnterInput()
	case tui.KeyUp, tui.KeyDown:
		m.handleDirectionInput(msg)
	default:
		handleDefaultInput(m, msg)
	}
}

func (m *model) handleCommand(msg tea.KeyPressMsg) (*model, tea.Cmd) {
	menu := m.getMenuActiveElement()

	switch msg.String() {
	case tui.CmdShiftEnter, tui.CmdAltEnter:
		return m, menu.HandleKeyShiftEnter(msg)
	case tui.CmdCtrlS:
		return m, saveCmd(m.save)
	}

	m.refreshViewport(false)

	return m, nil
}

// Side Menu Direction
func (m *model) handleDirectionInput(key tea.KeyPressMsg) {
	direction := sidemenu.HandleKeyArrow(sidemenu.Context{
		Menu:             m.getMenuActiveElement(),
		CurrentMenuIndex: m.activeMenuIndex,
		SibblingCount:    len(m.menu),
	}, key)

	m.updateMenuDirection(direction)
}

func (m *model) updateMenuDirection(msg direction.Direction) {
	previousMenuIndex := m.activeMenuIndex
	menuLength := len(m.menu)

	switch msg {
	case direction.Previous:
		if m.activeMenuIndex > 0 {
			m.activeMenuIndex--
		}
	case direction.Next:
		if m.activeMenuIndex < menuLength-1 {
			m.activeMenuIndex++
		}
	}

	if previousMenuIndex != m.activeMenuIndex {
		m.menu[previousMenuIndex].SetFocus(false)
		m.menu[m.activeMenuIndex].SetFocus(true)
	}
}

// Codex Action
func (m *model) handleCodexAction(msg codexmenu.Msg) (*model, tea.Cmd) {
	// Handle the codex action
	m.refreshViewport(true)

	return m, nil
}

// Dice Rolled
func (m *model) handleDiceRolled(msg dicemenu.Msg) (*model, tea.Cmd) {
	message := fmt.Sprintf("%s : %d", msg.Dice, msg.Value)

	m.engine.AddJournalEntry("Dice", message)
	m.journal = append(m.journal, message)

	m.refreshViewport(true)

	return m, nil
}

// Oracle Rolled
func (m *model) handleOracleRolled(msg oraclemenu.Msg) (*model, tea.Cmd) {
	message := fmt.Sprintf("Oracle rolled: %d, Result: %v, Critical: %t", msg.Result.Roll, msg.Result.Result, msg.Result.Critical)

	m.engine.AddJournalEntry("Oracle", message)
	m.journal = append(m.journal, message)

	m.refreshViewport(true)

	return m, nil
}

// Window
func (m *model) handleWindowResize(msg tea.WindowSizeMsg) (*model, tea.Cmd) {
	refreshMainView(m, msg)
	refreshMenuElement(m, msg)

	return m, nil
}

func refreshMainView(m *model, msg tea.WindowSizeMsg) {
	chatWidth := max(0, msg.Width-panelWidth-panelGap)

	m.mainview.SetHeight(max(0, msg.Height-m.textarea.Height()-1))
	m.mainview.SetWidth(chatWidth)
	m.textarea.SetWidth(chatWidth)

	if len(m.journal) > 0 && !m.isMenuActiveElementOpen() {
		m.mainview.SetContent(lipgloss.NewStyle().Width(m.mainview.Width()).Render(strings.Join(m.journal, "\n")))
	}

	m.mainview.GotoBottom()
}

func refreshMenuElement(m *model, msg tea.WindowSizeMsg) {
	activeEl := m.getMenuActiveElement()
	activeEl.HandleWindowResize(msg)
}
