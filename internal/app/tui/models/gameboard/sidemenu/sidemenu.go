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

	HandleKeyEnter(msg tea.Msg) error
	HandleKeyEsc(msg tea.Msg) error
	HandleCtrlN(msg tea.Msg) error
}

// Direction
type Direction string

var (
	PreviousMenu Direction = "previous"
	NextMenu     Direction = "next"
	None         Direction = "none"
)

// -- Helper -- //
func HandleKeyArrow(model MenuItem, key tea.KeyPressMsg) Direction {
	menuLength := len(model.GetList().Items())

	previousIndex := model.GetList().Index()
	newList, _ := model.GetList().Update(key)

	model.SetList(newList)
	newIndex := model.GetList().Index()

	switch key.String() {
	case tui.KeyUp:
		if isFirst := newIndex == 0; isFirst && previousIndex == 0 {
			model.SetOpen(false)
			return PreviousMenu
		}
	case tui.KeyDown:
		if isLast := newIndex == menuLength-1; isLast && previousIndex == newIndex {
			model.SetOpen(false)
			return NextMenu
		}
	}
	return None
}
