package gameboard

import (
	"solopg/internal/app/game"
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/models/codex"
	"solopg/internal/domain/gameplay"
	"solopg/internal/infrastructure/t"

	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"

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
	diceList    list.Model
	codexList   list.Model
	activeMenu  panelMenu

	codexView CodexView

	engine *game.Engine
	save   func() error
}

type panelMenu uint8

const (
	oracleMenu panelMenu = iota
	diceMenu
	codexMenu
)

type UiParams struct {
	Engine *game.Engine
	OnSave func() error
}

// NewModel returns the game view for embedding in the main TUI router.
func NewModel(params UiParams) model {
	ta := textarea.New()
	ta.Placeholder = t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "chat.placeholder"})
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
	vp.SetContent(t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "chat.welcome"}))
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
		diceList:   makeDiceModel(),
		codexList:  codex.MakeCodexListModel(panelWidth-4, codexMenuHeight),
		activeMenu: oracleMenu,
		codexView: CodexView{
			page: codexViewModel[viewport.Model]{
				model: viewport.New(viewport.WithWidth(30), viewport.WithHeight(5)),
				open:  false,
			},
		},

		engine: params.Engine,
		save:   params.OnSave,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

// HandlesEscape lets the global router forward Escape while a Codex page is open.
func (m model) HandlesEscape() bool {
	return m.codexView.form.open
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		handleWindowResize(&m, msg)

	case tea.KeyPressMsg:
		if m.codexView.form.open {
			return HandleCodexFormKey(m, msg)
		}
		switch msg.String() {
		case tui.KeySave:
			return m, saveCmd(m.save)
		case "ctrl+n":
			if m.codexView.page.open {
				return OpenCodexForm(m)
			}
		case tui.KeyEnter:
			if m.codexView.page.open {
				return m, nil
			}
			if m.showPanel && m.textarea.Value() == "" && m.activeMenu == oracleMenu {
				return handleOracleRoll(m)
			}
			if m.showPanel && m.textarea.Value() == "" && m.activeMenu == diceMenu {
				return handleDiceRoll(m)
			}
			if m.showPanel && m.textarea.Value() == "" && m.activeMenu == codexMenu {
				return OpenCodexPage(m)
			}
			return handleEnterInput(m)
		case "up", "down":
			if m.codexView.page.open {
				return HandleCodexPageNavigation(m, msg)
			}
			if m.showPanel && m.textarea.Value() == "" {
				return handlePanelNavigation(m, msg)
			}
			return handleDefaultInput(m, msg)
		case tui.KeyEsc:
			if m.codexView.page.open {
				m.codexView.page.open = false
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

// - Submodels - //
func makeDiceModel() list.Model {
	options := gameplay.ListDices()

	items := make([]list.Item, 0, len(options))
	for _, option := range options {
		items = append(items, tui.NewItem(option.GetName(), "", option))
	}

	model := list.New(items, list.NewDefaultDelegate(), panelWidth-4, diceMenuHeight)
	tui.ConfigureList(&model)
	return model
}

func makeOracleModel() list.Model {
	items := make([]list.Item, 0)
	for _, oracle := range gameplay.GetOracle() {
		if oracle.Visible {
			key := "oracle." + oracle.ID
			name := t.Localizer.MustLocalize(&goi18n.LocalizeConfig{
				MessageID: key,
				DefaultMessage: &goi18n.Message{
					ID:    key,
					Other: oracle.ID,
				},
			})
			items = append(items, tui.NewItem(name, "", oracle))
		}
	}

	model := list.New(items, list.NewDefaultDelegate(), panelWidth-4, oracleMenuHeight)
	tui.ConfigureList(&model)
	return model
}