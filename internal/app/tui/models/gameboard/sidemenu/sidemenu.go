package sidemenu

import (
	"solopg/internal/app/tui"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

// -- Contract - //
type MenuItem interface {
	ID() string
	IsOpen() bool
	SetOpen(bool)

	SetFocus(bool)

	GetList() list.Model
	SetList(list.Model)

	GetMenuView() string
	GetView() string
	GetFooter() []string

	HandleKeyShiftEnter(msg tea.Msg) tea.Cmd
	HandleKeyEsc(msg tea.Msg) error
	HandleCtrlN(msg tea.Msg) error
}

// Direction
type Direction string

var (
	Previous Direction = "previous"
	Next     Direction = "next"
	None     Direction = "none"
)

type Context struct {
	Menu             MenuItem
	CurrentMenuIndex int
	SibblingCount    int
}

func (m *Context) IsLastMenuElement() bool {
	return m.CurrentMenuIndex == m.SibblingCount-1
}

func (m *Context) IsFirstMenuElement() bool {
	return m.CurrentMenuIndex == 0
}

// -- Helper -- //
func HandleKeyArrow(metadata Context, key tea.KeyPressMsg) Direction {
	menu := metadata.Menu
	list := menu.GetList()
	itemCount := len(list.Items())

	previousIndex := list.Index()
	updatedList, _ := list.Update(key)

	menu.SetList(updatedList)
	currentIndex := list.Index()

	switch key.String() {
	case tui.KeyUp:
		if metadata.IsFirstMenuElement() {
			return None
		}
		if isFirstEl := currentIndex == 0; isFirstEl && previousIndex == 0 {
			menu.SetOpen(false)
			return Previous
		}
	case tui.KeyDown:
		if metadata.IsLastMenuElement() {
			return None
		}
		if isLastEL := currentIndex == itemCount-1; isLastEL && previousIndex == currentIndex {
			menu.SetOpen(false)
			return Next
		}
	}
	return None
}
