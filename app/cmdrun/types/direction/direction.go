package direction

import (
	cmdruntui "solopg/app/cmdrun/tui"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type Direction string

var (
	Previous Direction = "previous"
	Next     Direction = "next"
	None     Direction = "none"
)

func GetListDirection(list *list.Model, key tea.KeyPressMsg) (list.Model, Direction) {
	direction := None
	itemCount := len(list.Items())

	previousIndex := list.Index()
	updatedList, _ := list.Update(key)
	currentIndex := list.Index()

	switch key.String() {
	case cmdruntui.KeyUp:
		if isFirstEl := currentIndex == 0; isFirstEl && previousIndex == 0 {
			direction = Previous
		}
	case cmdruntui.KeyDown:
		if isLastEL := currentIndex == itemCount-1; isLastEL && previousIndex == currentIndex {
			direction = Next
		}
	}

	return updatedList, direction
}
