package codexsidemenu

import (
	"solopg/internal/app/game"
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/models/codexform"
	"solopg/internal/infrastructure/t"

	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
)

type Screen uint8

const (
	ScreenClosed Screen = iota
	ScreenPage
	ScreenForm
)

// Model owns the Codex menu and all Codex-specific screens and transitions.
type Model struct {
	// TODO : le model ne devrai pas avoir l'engine.
	engine *game.Engine
	menu   list.Model
	page   viewport.Model
	form   codexform.Model
	screen Screen
	active id
}

// Side panel menu
func NewModel(engine *game.Engine, width, menuHeight int) Model {
	items := makeCodexList()

	menu := list.New(items, list.NewDefaultDelegate(), width, menuHeight)
	tui.ConfigureList(&menu)

	return Model{
		engine: engine,
		menu:   menu,
		page:   viewport.New(viewport.WithWidth(width), viewport.WithHeight(menuHeight)),
		screen: ScreenClosed,
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



// --

func (m Model) MenuItemsCount() int { return len(m.menu.Items()) }
func (m Model) PageOpen() bool      { return m.screen == ScreenPage }
func (m Model) FormOpen() bool      { return m.screen == ScreenForm }
func (m Model) HandlesEscape() bool { return m.FormOpen() }


func (m *Model) SetSize(width, height int) {
	m.page.SetWidth(width)
	m.page.SetHeight(height)
	if m.FormOpen() {
		m.form.SetWidth(width)
	}
}

func (m *Model) Update(msg tea.Msg) tea.Cmd {
	if m.FormOpen() {
		return m.updateForm(msg)
	}
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
	case tui.KeyCtrlN:
		m.form = codexform.New(m.formKind(), m.page.Width())
		m.screen = ScreenForm
	case tui.KeyUp, tui.KeyDown:
		previous := m.menu.Index()
		m.UpdateMenu(msg)
		if previous != m.menu.Index() {
			m.selectActivePage()
		}
	}
	return nil
}

