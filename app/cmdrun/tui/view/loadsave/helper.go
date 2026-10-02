package loadsave

import (
	"solopg/app/cmdrun/domain/campaign"
	"solopg/app/cmdrun/tui/models"
	"solopg/app/shared/services/i19n"
	"solopg/app/shared/utils/transform"

	"charm.land/bubbles/v2/list"
)

func makeItems(data []campaign.Campaign) []list.Item {
	items := []list.Item{
		models.NewItem[*campaign.Campaign](transform.Capitalize(i19n.Localize("campaign:new")), "", nil),
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
