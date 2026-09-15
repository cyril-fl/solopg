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
	menu := m.getMenuActiveElement()

	return menu.HandleUpdate(sidemenu.UpdateParams{
		Model:    m,
		Msg:      msg,
		Delegate: m.handleUpdate,
		Refresh:  m.refreshViewport,
	})
}

func (m model) handleUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Tea messages
	case tea.WindowSizeMsg:
		return m.handleWindowResize(msg)
	case tea.KeyPressMsg:
		m.handleKeyPress(msg)
		return m.handleCommand(msg)

	// Cmd messages
	case codexmenu.Msg:
		return m.handleCodexAction(msg)
	case oraclemenu.Msg:
		return m.handleOracleRolled(msg)
	case dicemenu.Msg:
		return m.handleDiceRolled(msg)
	case tui.SaveMsg:
		return m.handleSaveInput(msg)
	case cursor.BlinkMsg:
		return m.handleCursorBlink(msg)
	}

	return m, nil
}

// HandlesEscape lets the global router forward Escape while a Codex page is open.
func (m model) HandlesEscape() bool {
	// return m.codexMenu.HandlesEscape()
	// TODO refactor la maniere dont on hndle escape
	return true
}
