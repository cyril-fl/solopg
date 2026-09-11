package gameboard

import (
	"solopg/internal/app/game"
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/models/gameboard/sidemenu"
	"solopg/internal/app/tui/models/gameboard/sidemenu/codexmenu"
	"solopg/internal/app/tui/models/gameboard/sidemenu/dicemenu"
	"solopg/internal/app/tui/models/gameboard/sidemenu/oraclemenu"
	"solopg/internal/infrastructure/t"

	"charm.land/bubbles/v2/cursor"
	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// - Model - //
type model struct {
	author      string
	journal     []string
	senderStyle lipgloss.Style

	textarea textarea.Model
	mainview viewport.Model

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
		journal:     initJournal(params.Engine),
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),

		textarea: initTextarea(),
		mainview: initViewport(t.Localize("chat.welcome")),

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
		m.handleWindowResize(msg)

	case tea.KeyPressMsg:
		menu := m.getMenuActiveElement()

		switch msg.String() {
		// Global
		case tui.KeyEnter:
			m.handleEnterInput()
		case tui.KeyUp, tui.KeyDown:
			m.handleMenuDirection(msg)

		// Subviews
		case tui.KeyCtrlN:
			menu.HandleCtrlN(msg)
		case tui.KeyEsc:
			menu.HandleKeyEsc(msg)

		// Command
		case tui.KeyShiftEnter:
			return m, menu.HandleKeyShiftEnter(msg)
		case tui.KeySave:
			return m, saveCmd(m.save)

		default:
			return handleDefaultInput(m, msg)
		}

	case codexmenu.Msg:
		// return m, nil
	case oraclemenu.Msg:
		// return m, nil
	case dicemenu.Msg:
		m.handleDiceRolled(msg)
	case tui.SaveMsg:
		m.handleSaveInput(msg)
	case cursor.BlinkMsg:
		return m.handleCursorBlink(msg)
	}

	m.refreshViewport(true)

	return m, nil
}

// HandlesEscape lets the global router forward Escape while a Codex page is open.
func (m model) HandlesEscape() bool {
	// return m.codexMenu.HandlesEscape()
	// TODO refactor la maniere dont on hndle escape
	return true
}
