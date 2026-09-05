package codex

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
	active CodexKind
}

func NewModel(engine *game.Engine, width, menuHeight int) Model {
	links := []CodexLink{
		{name: t.Localize("codex.npcs"), Kind: CodexNPCs},
		{name: t.Localize("codex.monsters"), Kind: CodexMonsters},
		{name: t.Localize("codex.locations"), Kind: CodexLocations},
		{name: t.Localize("codex.objects"), Kind: CodexObjects},
		{name: t.Localize("codex.objectives"), Kind: CodexObjectifs},
	}

	items := make([]list.Item, 0, len(links))
	for _, link := range links {
		items = append(items, tui.NewItem(link.name, "", link))
	}

	menu := list.New(items, list.NewDefaultDelegate(), width, menuHeight)
	tui.ConfigureList(&menu)

	return Model{
		engine: engine,
		menu:   menu,
		page:   viewport.New(viewport.WithWidth(width), viewport.WithHeight(menuHeight)),
		screen: ScreenClosed,
	}
}

func (m Model) MenuView() string    { return m.menu.View() }
func (m Model) MenuIndex() int      { return m.menu.Index() }
func (m Model) MenuItemsCount() int { return len(m.menu.Items()) }
func (m Model) PageOpen() bool      { return m.screen == ScreenPage }
func (m Model) FormOpen() bool      { return m.screen == ScreenForm }
func (m Model) HandlesEscape() bool { return m.FormOpen() }
func (m Model) View() string {
	switch m.screen {
	case ScreenPage:
		return m.page.View()
	case ScreenForm:
		return m.form.View()
	default:
		return ""
	}
}

func (m *Model) SetSize(width, height int) {
	m.page.SetWidth(width)
	m.page.SetHeight(height)
	if m.FormOpen() {
		m.form.SetWidth(width)
	}
}

func (m *Model) SetMenuSize(width, height int) {
	m.menu.SetSize(width, height)
}

func (m *Model) SelectMenu(index int) { m.menu.Select(index) }

func (m *Model) UpdateMenu(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.menu, cmd = m.menu.Update(msg)
	return cmd
}

func (m *Model) OpenPage() bool {
	selected, ok := m.menu.SelectedItem().(tui.Item[CodexLink])
	if !ok {
		return false
	}
	m.active = selected.Value().Kind
	m.screen = ScreenPage
	m.refreshPage()
	m.page.GotoTop()
	return true
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
	case "ctrl+n":
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

func (m *Model) updateForm(msg tea.Msg) tea.Cmd {
	key, ok := msg.(tea.KeyPressMsg)
	if ok && key.String() == tui.KeyEsc {
		m.screen = ScreenPage
		return nil
	}
	if ok && key.String() == tui.KeyEnter {
		if m.form.AdvanceOnEnter() {
			return nil
		}
		result, err := m.form.Submit()
		if err != nil {
			m.form.SetError(err)
			return nil
		}
		if err := AddCodexEntry(m.engine, result); err != nil {
			m.form.SetError(err)
			return nil
		}
		m.screen = ScreenPage
		m.refreshPage()
		m.page.GotoTop()
		return nil
	}
	m.form.SetError(nil)
	return m.form.Update(msg)
}

func (m *Model) selectActivePage() {
	selected, ok := m.menu.SelectedItem().(tui.Item[CodexLink])
	if !ok {
		return
	}
	m.active = selected.Value().Kind
	m.refreshPage()
	m.page.GotoTop()
}

func (m *Model) refreshPage() {
	m.page.SetContent(RenderCodexPage(m.engine, CodexLink{Kind: m.active, name: m.pageTitle()}))
}

func (m Model) pageTitle() string {
	for _, item := range m.menu.Items() {
		selected, ok := item.(tui.Item[CodexLink])
		if ok && selected.Value().Kind == m.active {
			return selected.Value().name
		}
	}
	return ""
}

func (m Model) formKind() codexform.Kind {
	switch m.active {
	case CodexNPCs:
		return codexform.NPCs
	case CodexMonsters:
		return codexform.Monsters
	case CodexLocations:
		return codexform.Locations
	case CodexObjects:
		return codexform.Objects
	default:
		return codexform.Objectifs
	}
}

/*
	TODO ne devrais pas appartenir a Code mais a T

Localizer devrai eetre un objetc
et exporter des fonction direct

var Localizer *Localizer

	func MustLocalize(lc *LocalizeConfig) string {
		localized, err := Localize(lc)
		if err != nil {
			panic(err)
		}
		return localized
	}
*/
