package codexmenu

import (
	"solopg/internal/app/tui"
	"solopg/internal/infrastructure/t"

	"charm.land/bubbles/v2/list"
)

// -- Side panel menu -- //
func makeCodexList(params CodexMenuParams) ([]list.Item) {
	pages := []*CodexMenuItem{
		NewMenuItem(SideMenuItemsParams{
			id:    "codex.locations",
			Size:  params.Size,
			Table: params.Codex.LocationsTable,
			// Logger: params.Codex.AddLocation,
		}),
		NewMenuItem(SideMenuItemsParams{
			id:    "codex.npcs",
			Size:  params.Size,
			Table: params.Codex.NpcsTable,
			// Logger: addNPCToCodex,
		}),
		NewMenuItem(SideMenuItemsParams{
			id:    "codex.monsters",
			Size:  params.Size,
			Table: params.Codex.MonstersTable,
			// Logger: addNPCToCodex,
		}),
		NewMenuItem(SideMenuItemsParams{
			id:    "codex.objects",
			Size:  params.Size,
			Table: params.Codex.ObjectsTable,
			// Logger: params.Codex.AddObject,
		}),
		NewMenuItem(SideMenuItemsParams{
			id:    "codex.objectives",
			Size:  params.Size,
			Table: params.Codex.ObjectifsTable,
			// Logger: params.Codex.AddObjective,
		}),
	}

	items := make([]list.Item, 0, len(pages))

	for _, page := range pages {
		items = append(items, tui.NewItem(t.Localize(page.id), "", page))
	}

	return items
}

// -- Helper -- //
func (m *CodexMenu) handleOpenPage(selected *CodexMenuItem) {
	for _, item := range m.list.Items() {
		item, ok := item.(tui.Item[*CodexMenuItem])
		
		if page := item.Value(); ok && page != selected {
			page.isOpen = false
		}
	
	}
	selected.isOpen = true
}

