package sidemenu

import (
	"solopg/internal/app/tui"
	"solopg/types/direction"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

// -- Contract - //
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

	// REFACTOR
	HandleKeyShiftEnter(msg tea.Msg) tea.Cmd
	HandleWindowResize(msg tea.WindowSizeMsg) tea.Cmd

	HandleUpdate(params UpdateParams) (tea.Model, tea.Cmd)
}
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

// Update
type UpdateParams struct {
	Model    tea.Model
	Msg      tea.Msg
	Delegate tui.DelegateUpdateFunc
	Refresh  tui.RefreshViewFunc
}

// -- Helper -- //
// func HandleKeyArrow(metadata Context, key tea.KeyPressMsg) Direction {
// 	// Todo diviser le code en deux, ce qui gerer la drection, et ce qui gere les side effect. Ainsi on pouurait utiliser cette fonction aussi pour des liste normal ex les fomrulaire.
// 	// Egalemnt deplacer cette nouvvele fonction dans "list"
// 	menu := metadata.Menu
// 	list := menu.GetList()
// 	itemCount := len(list.Items())

// 	previousIndex := list.Index()
// 	updatedList, _ := list.Update(key)

// 	menu.SetList(updatedList)
// 	currentIndex := list.Index()

// 	switch key.String() {
// 	case tui.KeyUp:
// 		if metadata.IsFirstMenuElement() {
// 			return None
// 		}
// 		if isFirstEl := currentIndex == 0; isFirstEl && previousIndex == 0 {
// 			menu.SetOpen(false)
// 			return Previous
// 		}
// 	case tui.KeyDown:
// 		if metadata.IsLastMenuElement() {
// 			return None
// 		}
// 		if isLastEL := currentIndex == itemCount-1; isLastEL && previousIndex == currentIndex {
// 			menu.SetOpen(false)
// 			return Next
// 		}
// 	}
// 	return None
// }

func HandleKeyArrow(metadata Context, key tea.KeyPressMsg) direction.Direction {
	list, direction := direction.GetListDirection(metadata.Menu.GetList(), key)

	metadata.Menu.SetList(list)

	if !(metadata.IsFirstMenuElement() || metadata.IsLastMenuElement()) {
		metadata.Menu.SetOpen(false)
	}

	return direction
}

