package bootui

import (
	"solopg/internal/domain/campaign"

	"charm.land/bubbles/v2/list"
)

type saveItem struct {
	title       string
	description string
	save        *campaign.Campaign
}

func (entry saveItem) Title() string       { return entry.title }
func (entry saveItem) Description() string { return entry.description }
func (entry saveItem) FilterValue() string { return entry.title }

func configureList(menu *list.Model) {
	menu.Title = ""
	menu.SetShowFilter(false)
	menu.SetShowPagination(false)
	menu.SetShowHelp(false)
	menu.SetShowStatusBar(false)
	menu.SetShowTitle(false)
}
