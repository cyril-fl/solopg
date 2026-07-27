package portalui

import (
	"solopg/internal/app/tui"

	"charm.land/bubbles/v2/list"
)

func makeItems() []list.Item {
	items := []list.Item{
		tui.NewItem("Accept", "", false),
		tui.NewItem("Reroll", "", true),
	}

	return items
}

func makeModel() list.Model {
	items := makeItems()

	model := list.New(items, list.NewDefaultDelegate(), 0, 0)
	tui.ConfigureList(&model)

	return model
}
