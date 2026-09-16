package codexmenu

import (
	"fmt"
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/models/gameboard/sidemenu"
	"solopg/internal/domain/codex"
	"solopg/internal/platform/form"
	"solopg/internal/platform/message"

	"solopg/types/direction"
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
		if page, ok := item.(tui.Item[*codexMenuItem]); !ok {
			continue
		} else if page.Value().isOpen {
			return true
		}
	}
	return false
}

func (m *codexMenu) SetOpen(open bool) {
	if open {
		return
	}

	for _, item := range m.list.Items() {
		item, ok := item.(tui.Item[*codexMenuItem])
		if page := item.Value(); ok {
			page.setOpen(false)
		}
	}
}

func (m *codexMenu) SetFocus(focused bool) {
	m.focused = focused
	tui.SetListFocus(&m.list, focused)
}

func (m *codexMenu) GetList() *list.Model {
	return &m.list
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
func (m *codexMenu) HandleUpdate(params sidemenu.UpdateParams) (tea.Model, tea.Cmd) {
	switch msg := params.Msg.(type) {
	case message.FormNextField, message.FormPreviousField:
		return m.delegateInputToForm(params)
	case message.FormSubmit:
		return m.handleFormSubmit(params)
	case message.FormError:
		fmt.Println("Form validation error occurred.")
		return params.Model, nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case tui.KeyEsc:
			return m.handleKeyEsc(params)
		case tui.KeyUp, tui.KeyDown:
			return m.delegateInputToForm(params)
		case tui.KeyEnter:
			return m.delegateInputToForm(params)
		case tui.CmdShiftEnter, tui.CmdAltEnter:
			return m.handleKeyShiftEnter(params)
		default:
			return m.delegateInputToForm(params)
		}
	}
	return params.Delegate(params.Msg)
}

func (m *codexMenu) HandleDirectionInput(direction direction.Direction) {
	_ = direction

	m.updatePageOnRedirection()
}

// Items
type codexMenuItem struct {
	id     string
	isOpen bool
	table  codex.Table

	form            *form.Model
	showForm        bool
	addEntryToCodex func(message string)
}

type sideMenuItemsParams struct {
	id              string
	table           codex.Table
	form            func() *form.Model
	addEntryToCodex func(message string)
}

func newMenuItem(params sideMenuItemsParams) *codexMenuItem {
	return &codexMenuItem{
		id:     params.id,
		isOpen: false,
		table:  params.table,

		form:            params.form(),
		showForm:        false,
		addEntryToCodex: params.addEntryToCodex,
	}
}

// Message
func sendMsg() tea.Cmd {
	return func() tea.Msg {
		return Msg{}
	}
}

type Msg struct {
}
