package gameboard

import (
	"solopg/internal/app/game"
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/models/gameboard/sidemenu"

	"charm.land/bubbles/v2/cursor"
	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// - Model - //
type model struct {
	author      string
	journal    []string
	senderStyle lipgloss.Style

	textarea textarea.Model
	viewport viewport.Model

	menu            []sidemenu.MenuItem
	activeMenuIndex int

	engine *game.Engine
	save   func() error

	err error
}

type UiParams struct {
	Engine *game.Engine
	OnSave func() error
}

func NewModel(params UiParams) model {
	return model{
		author:      initAuthor(params.Engine),
		journal:    initJournal(params.Engine),
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),

		textarea: initTextarea(),
		viewport: initViewport(),

		engine: params.Engine,
		save:   params.OnSave,

		menu:            initSideMenu(params.Engine),
		activeMenuIndex: 0,

		err: nil,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		handleWindowResize(&m, msg)

	case tea.KeyPressMsg:
		// if m.codexMenu.FormOpen() {
		// 	return m, m.codexMenu.Update(msg)
		// }
		menu := m.getActiveItem()

		switch msg.String() {
		case tui.KeySave:
			return m, saveCmd(m.save)
		case tui.KeyCtrlN:
			// if m.codexMenu.PageOpen() {
			// 	return m, m.codexMenu.Update(msg)
			// }
			menu.HandleCtrlN(msg)

		case tui.KeyEnter:
			menu.HandleKeyEnter(msg)
			// if m.codexMenu.PageOpen() {
			// 	return m, m.codexMenu.Update(msg)
			// }

			// // if m.showPanel && m.textarea.Value() == "" {
			// if m.activeMenu == oracleMenu {
			// 	return handleOracleRoll(m)
			// }

			// if m.activeMenu == diceMenu {
			// 	return handleDiceRoll(m)
			// }
			// if m.activeMenu == codexMenu {
			// 	m.codexMenu.OpenPage()
			// 	return m, nil
			// }
			// // }

			// return handleEnterInput(m)
		case tui.KeyUp, tui.KeyDown:
			direction := sidemenu.HandleKeyArrow(menu, msg)
			return handleMenuDirection(m, direction)
		case tui.KeyEsc:
			menu.HandleKeyEsc(msg)
			// if m.codexMenu.PageOpen() {
			// 	return m, m.codexMenu.Update(msg)
			// }
		default:
			return handleDefaultInput(m, msg)
		}

	case cursor.BlinkMsg:
		return handleCursorBlink(m, msg)
	case tui.SaveMsg:
		return handleSaveInput(m, msg)
	}

	return m, nil
}

// HandlesEscape lets the global router forward Escape while a Codex page is open.
// func (m model) HandlesEscape() bool {
// 	// return m.codexMenu.HandlesEscape()
// }
