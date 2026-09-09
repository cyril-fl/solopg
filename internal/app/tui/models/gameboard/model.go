package gameboard

import (
	"solopg/internal/app/game"
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/models/gameboard/sidemenu/dicemenu"
	"solopg/internal/app/tui/models/gameboard/sidemenu/oraclemenu"
	"solopg/internal/infrastructure/t"
	"solopg/types/size"

	"charm.land/bubbles/v2/cursor"
	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// - Model - //
type model struct {
	author      string
	messages    []string
	senderStyle lipgloss.Style

	textarea textarea.Model
	viewport viewport.Model

	menu           []MenuItem
	activeMenuItem MenuItem

	engine *game.Engine
	save   func() error

	err error
}

type UiParams struct {
	Engine *game.Engine
	OnSave func() error
}

func NewModel(params UiParams) model {
	menuSize := size.Size{
		Width:  panelWidth - 4,
		Height: oracleMenuHeight,
	}

	oraclem := oraclemenu.NewMenuItem(menuSize)
	dicem := dicemenu.NewMenuItem(menuSize)
	// codexm := codexmenu.NewMenuItem(codexmenu.SideMenuParams{
	// 	Size:  menuSize,
	// 	Codex: params.Engine.State.Codex.EnsureInitialized(),
	// 	// TODO faire en sorte remplacer logger par une cmd .
	// 	Logger: func(message string) {
	// 		params.Engine.AddJournalEntry("Codex", message)
	// 	},
	// })

	return model{
		author:      "Me",
		messages:    initJournalContent(params.Engine),
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),

		textarea: initTextarea(),
		viewport: initViewport(),

		engine: params.Engine,
		save:   params.OnSave,

		menu: []MenuItem{
			oraclem,
			dicem,
			// codexm,
		},
		activeMenuItem: oraclem,

		err: nil,
	}
}

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

func initJournalContent(engine *game.Engine) []string {
	journalContent := make([]string, 0, len(engine.State.Journal.Entries))

	for _, entry := range engine.State.Journal.Entries {
		journalContent = append(journalContent, entry.String())
	}

	return journalContent
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
		switch msg.String() {
		case tui.KeySave:
			return m, saveCmd(m.save)
		case tui.KeyCtrlN:
			// if m.codexMenu.PageOpen() {
			// 	return m, m.codexMenu.Update(msg)
			// }
			m.activeMenuItem.HandleCtrlN(msg)

		case tui.KeyEnter:
			m.activeMenuItem.HandleKeyEnter(msg)
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
			// m.activeMenuItem.HandleKeyArrow(msg)

			// if m.codexMenu.PageOpen() {
			// 	return m, m.codexMenu.Update(msg)
			// }
			// return handlePanelNavigation(m, msg)
		case tui.KeyEsc:
			m.activeMenuItem.HandleKeyEsc(msg)
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
