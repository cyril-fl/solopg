package gameui

import (
	"charm.land/bubbles/v2/list"
)

// menuItem satisfies list.DefaultItem for use with list.NewDefaultDelegate.
type menuItem struct {
	title       string
	description string
}

func (entry menuItem) Title() string       { return entry.title }
func (entry menuItem) Description() string { return entry.description }
func (entry menuItem) FilterValue() string { return entry.title + " " + entry.description }

func newMenuList() list.Model {
	items := []list.Item{
		menuItem{title: ". Generate report", description: ""},
		menuItem{title: ". Generate wallet", description: ""},
	}

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false

	listModel := list.New(items, delegate, 0, 0)
	listModel.Title = ""
	// TODO use config
	listModel.SetShowFilter(false)
	listModel.SetShowPagination(false)
	listModel.SetShowHelp(false)
	listModel.SetShowStatusBar(false)

	return listModel
}
