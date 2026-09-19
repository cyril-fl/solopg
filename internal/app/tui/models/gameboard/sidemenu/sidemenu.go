package sidemenu

import (
	"solopg/internal/app/tui"
	"solopg/types/direction"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

// - Contract - //
type MenuItem interface {
	ID() string
	IsOpen() bool
	SetOpen(bool)

	SetFocus(bool)

	GetList() *list.Model
	SetList(list.Model)

	GetMenuView() string
	GetView() string
	GetFooter() []string

	HandleWindowResize(msg tea.WindowSizeMsg) tea.Cmd
	HandleUpdate(params UpdateParams) (tea.Model, tea.Cmd)

	HandleDirectionInput(direction direction.Direction)
}

// - Context - //
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

// - Update - //
type UpdateParams struct {
	Model    tea.Model
	Msg      tea.Msg
	Delegate tui.DelegateUpdateFunc
	Refresh  tui.RefreshViewFunc
}

// - Helper - //
func HandleKeyArrow(metadata Context, key tea.KeyPressMsg) direction.Direction {
	list, newdirection := direction.GetListDirection(metadata.Menu.GetList(), key)

	metadata.Menu.SetList(list)

	isFirstElementToNext := metadata.IsFirstMenuElement() && newdirection == direction.Next
	isLastElementToPrevious := metadata.IsLastMenuElement() && newdirection == direction.Previous
	isBetweenElements := !(metadata.IsFirstMenuElement() || metadata.IsLastMenuElement())

	if isFirstElementToNext || isBetweenElements || isLastElementToPrevious {
		metadata.Menu.SetOpen(false)
	}

	return newdirection
}
