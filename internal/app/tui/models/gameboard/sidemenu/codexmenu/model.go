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

// -- Side panel menu -- //
type Model struct {
	codex           *codex.Codex // TODO :  voir s il degage ou si je fait des sorte de call back ect comme addjournm
	addJournalEntry func(message string)
	menu            *CodexMenu
	viewport        viewport.Model
}

type SideMenuParams struct {
	Size   size.Size
	Codex  *codex.Codex
	Logger func(message map[string]string) error
}

type SideMenuItemsParams struct {
	Size  size.Size
	Table any //*codex.Table[any]
	// Logger func(message map[string]string) error

}

func NewModel(params SideMenuParams) Model {
	return Model{
		codex: params.Codex.EnsureInitialized(),
		// addJournalEntry: params.Logger,
		menu:     newMenu(params),
		viewport: viewport.New(viewport.WithWidth(params.Size.Width), viewport.WithHeight(params.Size.Height)),
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

// -- CodexMenu -- //

// Menu
type CodexMenu struct {
	list  *list.Model
	pages []*CodexMenuItem
}

func newMenu(params SideMenuParams) *CodexMenu {
	items := makeCodexList()
	menu := list.New(items, list.NewDefaultDelegate(), params.Size.Width, params.Size.Height)
	tui.ConfigureList(&menu)

	return &CodexMenu{
		list: &menu,
		pages: []*CodexMenuItem{
			NewMenuItem(SideMenuItemsParams{
				Size:  params.Size,
				Table: params.Codex.LocationsTable,
				// Logger: params.Codex.AddLocation,
			}),
			NewMenuItem(SideMenuItemsParams{
				Size:  params.Size,
				Table: params.Codex.NpcsTable,
				// Logger: addNPCToCodex,
			}),
			NewMenuItem(SideMenuItemsParams{
				Size:  params.Size,
				Table: params.Codex.MonstersTable,
				// Logger: addNPCToCodex,
			}),
			NewMenuItem(SideMenuItemsParams{
				Size:  params.Size,
				Table: params.Codex.ObjectsTable,
				// Logger: params.Codex.AddObject,
			}),
			NewMenuItem(SideMenuItemsParams{
				Size:  params.Size,
				Table: params.Codex.ObjectifsTable,
				// Logger: params.Codex.AddObjective,
			}),
		},
	}
}

func (m *CodexMenu) ID() string {
	return "codex"
}
func (m *CodexMenu) IsOpen() bool {
	return m.isOpen()
}
func (m *CodexMenu) SetOpen(open bool) {
	if open {
		if m.currentPage() == nil && len(m.pages) > 0 {
			m.pages[0].SetOpen(true)
		}
	} else {
		for _, item := range m.pages {
			item.SetOpen(false)
		}
	}
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

func (m *CodexMenu) isOpen() bool {
	for _, item := range m.pages {
		if item.IsOpen() {
			return true
		}
	}
	return false
}

func (m *CodexMenu) currentPage() *CodexMenuItem {
	for _, item := range m.pages {
		if item.IsOpen() {
			return item
		}
	}
	return nil
}

func (m *CodexMenu) pageTitle() string {
	current := m.currentPage()
	if current == nil {
		return ""
	}
	return current.ID()
}

func (m *CodexMenu) openPage(id string) {
	var toOpen *CodexMenuItem
	for _, i := range m.pages {
		if i.ID() == id {
			if i.IsOpen() {
				break
			}

			toOpen = i

			break
		} else {
			i.SetOpen(false)
		}
	}

	toOpen.SetOpen(true)
	// 	selected, ok := m.menu.SelectedItem().(tui.Item[link])
	// 	if !ok {
	// 		return false
	// 	}
	// 	m.active = selected.Value().id
	// 	m.screen = ScreenPage
	// 	m.refreshPage()
	// 	m.viewport.GotoTop()
	// 	return true
}

func (m *CodexMenu) closePage(id string) {
	for _, i := range m.pages {
		if i.ID() == id {
			i.SetOpen(false)
			break
		}
	}
}

// Items
type CodexMenuItem struct {
	id              string
	list            list.Model
	isOpen          bool
	table           any // *codex.Table[any]
	addJournalEntry func(message string)
}

func NewMenuItem(params SideMenuItemsParams) *CodexMenuItem {
	items := makeCodexList()

	menu := list.New(items, list.NewDefaultDelegate(), params.Size.Width, params.Size.Height)
	tui.ConfigureList(&menu)

	return &CodexMenuItem{
		id:     "codex",
		isOpen: false,
		list:   menu,
		table:  params.Table,
		// addJournalEntry: params.Logger,
	}
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

func (m *CodexMenuItem) getFooter() []string {
	return []string{t.Localize("add"), t.Localize("back_to_chat")}
}

func (m *CodexMenuItem) handleKeyEnter(msg tea.Msg) error {
	return nil
}
func (m *CodexMenuItem) handleKeyEsc(msg tea.Msg) error {
	return nil
}
func (m *CodexMenuItem) handleKeyArrow(msg tea.Msg) error {
	return nil
}
func (m *CodexMenuItem) handleCtrlN(msg tea.Msg) error {
	return nil
}
