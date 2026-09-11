package codexmenu

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/codex"
	"strings"

	"solopg/types/size"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

// -- CodexMenu -- //
// Menu
func NewSideMenu(params CodexMenuParams, focused bool) *CodexMenu {
	items := makeCodexList(params)
	menu := list.New(items, list.NewDefaultDelegate(), params.Size.Width, params.Size.Height)
	tui.ConfigureList(&menu)
	tui.SetListFocus(&menu, focused)

	return &CodexMenu{
		id:      "codex",
		focused: focused,
		list:    menu,
	}
}

type CodexMenu struct {
	id      string
	focused bool
	list    list.Model
}

type CodexMenuParams struct {
	Size  size.Size
	Codex *codex.Codex
	// Logger func(message map[string]string) error
}

func (m *CodexMenu) ID() string {
	return m.id
}
func (m *CodexMenu) IsOpen() bool {
	for _, item := range m.list.Items() {
		page, ok := item.(tui.Item[*CodexMenuItem])
		if !ok {
			continue
		}
		if page.Value().isOpen {
			return true
		}
	}
	return false
}

// SetOpen sets the open state of the CodexMenu. If open true, it does nothing. If open false, it closes all pages in the menu.
func (m *CodexMenu) SetOpen(open bool) {
	if open {
		return
	}

	for _, item := range m.list.Items() {
		item, ok := item.(tui.Item[*CodexMenuItem])
		if page := item.Value(); ok {
			page.isOpen = false
		}
	}
}

func (m *CodexMenu) SetFocus(focused bool) {
	m.focused = focused
	tui.SetListFocus(&m.list, focused)
}

func (m *CodexMenu) GetList() list.Model {
	return m.list
}
func (m *CodexMenu) SetList(list list.Model) {
	m.list = list
}

func (m *CodexMenu) HandleKeyShiftEnter(msg tea.Msg) tea.Cmd {
	if item, ok := m.list.SelectedItem().(tui.Item[*CodexMenuItem]); ok {
		page := item.Value()
		m.handleOpenPage(page)
	}

	return func() tea.Msg {
		return Msg{}
	}
}
func (m *CodexMenu) HandleKeyEsc(msg tea.Msg) error {
	return nil
}

func (m *CodexMenu) HandleCtrlN(msg tea.Msg) error {
	m.SetOpen(false)
	return nil
}

// Items
type CodexMenuItem struct {
	id              string
	list            list.Model
	isOpen          bool
	table           codex.Table
	addJournalEntry func(message string)
}

type SideMenuItemsParams struct {
	id    string
	Size  size.Size
	Table codex.Table
	// Logger func(message map[string]string) error
}

func NewMenuItem(params SideMenuItemsParams) *CodexMenuItem {
	return &CodexMenuItem{
		id:     params.id,
		isOpen: false,
		// list:   menu,
		table: params.Table,
		// addJournalEntry: params.Logger,
	}
}

func (m *CodexMenuItem) View() string {
	return "Page: " + m.id + "\n\n" + strings.Join(m.table.Summaries(), "\n\n")
}

type Msg struct {
}
