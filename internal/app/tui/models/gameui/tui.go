package gameui

import (
	"fmt"
	"solopg/internal/app/game"
	"solopg/internal/app/tui"

	"charm.land/bubbles/v2/cursor"
	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Ui struct {
	program *tea.Program
}

type UiParams struct {
	Engine *game.Engine
	OnSave func() error
}

func NewUi(params UiParams) *Ui {
	return &Ui{
		program: tea.NewProgram(NewModel(params)),
	}
}

func (ui *Ui) Start() error {
	res, err := ui.program.Run()
	if err != nil {
		return err
	}

	finalModel, ok := res.(model)
	if !ok {
		return fmt.Errorf("unexpected model type %T", res)
	}

	_ = finalModel

	return nil
}

type model struct {
	viewport    viewport.Model
	author      string
	messages    []string
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error

	engine *game.Engine
	save   func() error
}

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
		textarea: ta,
		author:   "Me",
		messages:    journalContent,
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,

		engine: params.Engine,
		save:   params.OnSave,
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
		switch msg.String() {
		case tui.KeySave:
			return m, saveCmd(m.save)
		case tui.KeyEnter:
			return handleEnterInput(m)
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
