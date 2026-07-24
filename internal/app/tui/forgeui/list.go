package forgeui

import (
	"solopg/internal/domain/card/characters/archetypes"

	"charm.land/bubbles/v2/list"
)

type raceItem struct {
	title string
	race  archetypes.Race
}

func (entry raceItem) Title() string       { return entry.title }
func (entry raceItem) Description() string { return "" }
func (entry raceItem) FilterValue() string { return entry.title }

type classItem struct {
	title string
	class archetypes.Class
}

func (entry classItem) Title() string       { return entry.title }
func (entry classItem) Description() string { return "" }
func (entry classItem) FilterValue() string { return entry.title }

func configureList(menu *list.Model) {
	menu.Title = ""
	menu.SetShowFilter(false)
	menu.SetShowPagination(false)
	menu.SetShowHelp(false)
	menu.SetShowStatusBar(false)
	menu.SetShowTitle(false)
}
