package codexmenu

import (
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/models/gameboard/sidemenu"
	"solopg/internal/domain/codex"
	"solopg/internal/platform/form"
	"solopg/internal/platform/message"

	"solopg/types/size"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

// -- CodexMenu -- //
// Menu
func NewSideMenu(params CodexMenuParams, focused bool) *codexMenu {
	items := makeCodexList(params)
	menu := list.New(items, list.NewDefaultDelegate(), params.Size.Width, params.Size.Height)
	tui.ConfigureList(&menu)
	tui.SetListFocus(&menu, focused)

	return &codexMenu{
		id:      "codex",
		focused: focused,
		list:    menu,
	}
}

type CodexMenuParams struct {
	Size  size.Size
	Codex *codex.Codex
	// Logger func(message map[string]string) error
}

type codexMenu struct {
	id      string
	focused bool
	list    list.Model
}

// / Getters and Setters
func (m *codexMenu) ID() string {
	return m.id
}

func (m *codexMenu) IsOpen() bool {
	for _, item := range m.list.Items() {
		page, ok := item.(tui.Item[*codexMenuItem])
		if !ok {
			continue
		}
		if page.Value().isOpen {
			return true
		}
	}
	return false
}

// SetOpen sets the open state of the CodexMenu. If open true, @it does nothing. If open false, it closes all pages in the menu.
func (m *codexMenu) SetOpen(open bool) {
	if open {
		return
	}

	for _, item := range m.list.Items() {
		item, ok := item.(tui.Item[*codexMenuItem])
		if page := item.Value(); ok {
			page.isOpen = false
			page.showForm = false
		}
	}
}

func (m *codexMenu) SetFocus(focused bool) {
	m.focused = focused
	tui.SetListFocus(&m.list, focused)
}

func (m *codexMenu) GetList() list.Model {
	return m.list
}

func (m *codexMenu) SetList(list list.Model) {
	m.list = list
}

func (menu *codexMenu) GetFormFromCurrentPage() *form.Model {
	if currentPage := menu.getCurrentPage(); currentPage != nil && menu.IsOpen() {
		return currentPage.form
	}
	return nil
}

func (menu *codexMenu) SetFormOnCurrentPage(form form.Model) {
	if currentPage := menu.getCurrentPage(); currentPage != nil && menu.IsOpen() {
		currentPage.form = &form
	}
}

// / Handlers
func (m *codexMenu) HandleKeyShiftEnter(msg tea.Msg) tea.Cmd {
	if currentPage := m.getCurrentPage(); currentPage != nil && m.IsOpen() {
		currentPage.toggleShowForm()
	}

	if nextPage := m.getSelectedPage(); nextPage != nil {
		m.handleOpenPage(nextPage)
	}

	return func() tea.Msg {
		return Msg{}
	}
}

func (m *codexMenu) HandleUpdate(params sidemenu.UpdateParams) (tea.Model, tea.Cmd) {
	// Arrow keys are translated by a field into form navigation messages.
	// Those messages are sent back through the Bubble Tea update loop, so they
	// must be routed to the form as well instead of being delegated globally.
	switch params.Msg.(type) {
	case message.FormNextField, message.FormPreviousField:
		return m.delegateInputToForm(params)
	}

	switch msg := params.Msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {
		case tui.KeyEsc:
			return m.handleKeyEsc(params)
		case tui.KeyUp, tui.KeyDown:
			return m.delegateInputToForm(params)
		default:
			return m.delegateInputToForm(params)
		}
	}

	return params.Delegate(params.Msg)

}

// Items
type codexMenuItem struct {
	id     string
	isOpen bool
	table  codex.Table

	form     *form.Model
	showForm bool
	// addJournalEntry func(message string)
}

type sideMenuItemsParams struct {
	id    string
	table codex.Table
	form  func() *form.Model
	// Logger func(message map[string]string) error
}

func newMenuItem(params sideMenuItemsParams) *codexMenuItem {
	return &codexMenuItem{
		id:     params.id,
		isOpen: false,
		table:  params.table,

		form:     params.form(),
		showForm: false,
		// addJournalEntry: params.Logger,
	}
}

// Message
type Msg struct {
}

func sendMsg() tea.Cmd {
	return func() tea.Msg {
		return Msg{}
	}
}
