package gameboard

import (
	"solopg/internal/app/game"
	"solopg/internal/app/tui"

	"charm.land/bubbles/v2/cursor"
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type model struct {
	viewport    viewport.Model
	author      string
	messages    []string
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error
	showPanel   bool
	oracleList  list.Model
	codexList   list.Model
	activeMenu  panelMenu
	codexPage   viewport.Model
	pageOpen    bool

	engine *game.Engine
	save   func() error
}

type panelMenu uint8

const (
	oracleMenu panelMenu = iota
	codexMenu
)

// NewModel returns the game view for embedding in the main TUI router.
func NewModel(params UiParams) model {
	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.SetVirtualCursor(false)
	ta.Focus()

	// TODO cherche ce que ca fait
	// ta.Prompt = "┃ "
	ta.CharLimit = 280

	ta.SetWidth(30)
	ta.SetHeight(3)

	// Remove cursor line styling
	s := ta.Styles()
	s.Focused.CursorLine = lipgloss.NewStyle()
	ta.SetStyles(s)

	ta.ShowLineNumbers = false

	vp := viewport.New(viewport.WithWidth(30), viewport.WithHeight(5))
	// TODO voir pour set autre choses en fonction de message deja present ou non.
	vp.SetContent(`Welcome to the chat room!
Type a message and press Enter to send.`)
	vp.KeyMap.Left.SetEnabled(false)
	vp.KeyMap.Right.SetEnabled(false)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	journalContent := make([]string, 0, len(params.Engine.State.Journal.Entries))

	for _, entry := range params.Engine.State.Journal.Entries {
		journalContent = append(journalContent, entry.String())
	}

	return model{
		textarea:    ta,
		author:      "Me",
		messages:    journalContent,
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,

		oracleList: makeOracleModel(),
		codexList:  makeCodexModel(),
		activeMenu: oracleMenu,
		codexPage:  viewport.New(viewport.WithWidth(30), viewport.WithHeight(5)),

		engine: params.Engine,
		save:   params.OnSave,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

// HandlesEscape lets the global router forward Escape while a Codex page is open.
func (m model) HandlesEscape() bool {
	return m.pageOpen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		handleWindowResize(&m, msg)

	case tea.KeyPressMsg:
		switch msg.String() {
		case tui.KeySave:
			return m, saveCmd(m.save)
		case tui.KeyEnter:
			if m.pageOpen {
				return m, nil
			}
			if m.showPanel && m.textarea.Value() == "" && m.activeMenu == oracleMenu {
				return handleOracleRoll(m)
			}
			if m.showPanel && m.textarea.Value() == "" && m.activeMenu == codexMenu {
				return openCodexPage(m)
			}
			return handleEnterInput(m)
		case "up", "down":
			if m.pageOpen {
				var cmd tea.Cmd
				m.codexPage, cmd = m.codexPage.Update(msg)
				return m, cmd
			}
			if m.showPanel && m.textarea.Value() == "" {
				return handlePanelNavigation(m, msg)
			}
			return handleDefaultInput(m, msg)
		case tui.KeyEsc:
			if m.pageOpen {
				m.pageOpen = false
				return m, nil
			}
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
