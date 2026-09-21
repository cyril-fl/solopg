package loadsave

import (
	"solopg/app/domain/campaign"
	"solopg/app/services/t"
	"solopg/app/tui/models"

	"charm.land/bubbles/v2/list"
)

func makeItems(data []campaign.Campaign) []list.Item {
	items := []list.Item{
		models.NewItem[*campaign.Campaign](t.Localize("new_save"), "", nil),
	}

	for _, s := range data {
		newItem := models.NewItem(s.Title(), s.Description(), &s)
		items = append(items, newItem)
	}

	return items
}

func makeModel(data []campaign.Campaign) list.Model {
	items := makeItems(data)
	model := list.New(items, list.NewDefaultDelegate(), 0, 0)
	models.ConfigureList(&model)

	return model
}
