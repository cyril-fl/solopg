package loadsave

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/campaign"

	"charm.land/bubbles/v2/list"
)

func makeItems(data []campaign.Campaign) []list.Item {
	items := []list.Item{
		tui.NewItem[*campaign.Campaign]("New Save", "", nil),
	}

	for _, s := range data {
		newItem := tui.NewItem(s.Title(), s.Description(), &s)
		items = append(items, newItem)
	}

	return items
}

func makeModel(data []campaign.Campaign) list.Model {
	items := makeItems(data)
	model := list.New(items, list.NewDefaultDelegate(), 0, 0)
	tui.ConfigureList(&model)

	return model
}
