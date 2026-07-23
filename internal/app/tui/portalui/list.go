package portalui

import (
	// "solopg/internal/domain/campaign"

	"charm.land/bubbles/v2/list"
)

type choiceItem struct {
	title string
	retry bool
}

func (entry choiceItem) Title() string       { return entry.title }
func (entry choiceItem) Description() string { return "" }
func (entry choiceItem) FilterValue() string { return entry.title }

func configureList(menu *list.Model) {
	menu.Title = ""
	menu.SetShowFilter(false)
	menu.SetShowPagination(false)
	menu.SetShowHelp(false)
	menu.SetShowStatusBar(false)
	menu.SetShowTitle(false)
}
