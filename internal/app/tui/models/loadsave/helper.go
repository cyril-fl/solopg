package loadsave

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/campaign"
	"solopg/internal/infrastructure/t"

	"charm.land/bubbles/v2/list"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

func makeItems(data []campaign.Campaign) []list.Item {
	items := []list.Item{
		tui.NewItem[*campaign.Campaign](t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "new_save"}), "", nil),
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
