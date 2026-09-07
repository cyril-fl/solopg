package codexmenu

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/codex"
	"solopg/internal/infrastructure/t"
	"solopg/types/size"

	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
)

type Screen uint8

const (
	ScreenClosed Screen = iota
	ScreenPage
	// ScreenForm
)

// -- Side panel menu -- //
type Model struct {
	codex           *codex.Codex // TODO : degager l'engine.
	addJournalEntry func(message string)
	menu            list.Model
	page            viewport.Model
	// form            codexform.Model
	screen Screen
	active id
}

type SideMenuParams struct {
	Size   size.Size
	Codex  *codex.Codex
	Logger func(message string)
}

func NewModel(params SideMenuParams) Model {
	items := makeCodexList()

	menu := list.New(items, list.NewDefaultDelegate(), params.Size.Width, params.Size.Height)
	tui.ConfigureList(&menu)

	codex := params.Codex.EnsureInitialized()

	return Model{
		codex:           codex,
		addJournalEntry: params.Logger,
		menu:            menu,
		page:            viewport.New(viewport.WithWidth(params.Size.Width), viewport.WithHeight(params.Size.Height)),
		screen:          ScreenClosed,
	}
}

func makeCodexList() []list.Item {
	links := []link{
		{id: CodexNPCs, name: t.Localize("codex.npcs")},
		{id: CodexMonsters, name: t.Localize("codex.monsters")},
		{id: CodexLocations, name: t.Localize("codex.locations")},
		{id: CodexObjects, name: t.Localize("codex.objects")},
		{id: CodexObjectifs, name: t.Localize("codex.objectives")},
	}

	items := make([]list.Item, 0, len(links))

	for _, link := range links {
		items = append(items, tui.NewItem(link.name, "", link))
	}

	return items
}

func NewMenuItem(params SideMenuParams) *CodexMenuItem {
	items := makeCodexList()

	menu := list.New(items, list.NewDefaultDelegate(), params.Size.Width, params.Size.Height)
	tui.ConfigureList(&menu)

	codex := params.Codex.EnsureInitialized()
	return &CodexMenuItem{
		id:     "codex",
		isOpen: false,

		list:            menu,
		codex:           codex.EnsureInitialized(),
		addJournalEntry: params.Logger,
	}
}

type CodexMenuItem struct {
	id              string
	list            list.Model
	isOpen          bool
	codex           *codex.Codex
	addJournalEntry func(message string)
}

func (m *CodexMenuItem) ID() string {
	return m.id
}
func (m *CodexMenuItem) IsOpen() bool {
	return m.isOpen
}
func (m *CodexMenuItem) SetOpen(open bool) {
	m.isOpen = open
}

func (m *CodexMenuItem) GetFooter() []string {
	return []string{t.Localize("add"), t.Localize("back_to_chat")}
}

func (m *CodexMenuItem) HandleKeyEnter(msg tea.Msg) error {
	return nil
}
func (m *CodexMenuItem) HandleKeyEsc(msg tea.Msg) error {
	return nil
}
func (m *CodexMenuItem) HandleKeyArrow(msg tea.Msg) error {
	return nil
}
func (m *CodexMenuItem) HandleCtrlN(msg tea.Msg) error {
	return nil
}

// --
func (m Model) MenuItemsCount() int { return len(m.menu.Items()) }
func (m Model) PageOpen() bool      { return m.screen == ScreenPage }

// func (m Model) FormOpen() bool      { return m.screen == ScreenForm }
// func (m Model) HandlesEscape() bool { return m.PageOpen() || m.FormOpen() }
func (m Model) HandlesEscape() bool { return m.PageOpen() }

func (m *Model) SetSize(width, height int) {
	m.page.SetWidth(width)
	m.page.SetHeight(height)
	// if m.FormOpen() {
	// 	m.form.SetWidth(width)
	// }
}

func (m *Model) Update(msg tea.Msg) tea.Cmd {
	// if m.FormOpen() {
	// 	return m.updateForm(msg)
	// }
	if !m.PageOpen() {
		return nil
	}

	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return nil
	}
	switch key.String() {
	case tui.KeyEsc:
		m.screen = ScreenClosed
	// case tui.KeyCtrlN:
	// m.form = codexform.New(m.formKind(), m.page.Width())
	// m.screen = ScreenForm
	case tui.KeyUp, tui.KeyDown:
		previous := m.menu.Index()
		m.UpdateMenu(msg)
		if previous != m.menu.Index() {
			m.selectActivePage()
		}
	}
	return nil
}
