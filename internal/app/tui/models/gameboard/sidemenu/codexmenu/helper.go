package codexmenu

import (
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/models/gameboard/sidemenu"
	"solopg/internal/infrastructure/t"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

// -- Side panel -- //
func makeCodexList(params CodexMenuParams) []list.Item {
	pages := []*codexMenuItem{
		newMenuItem(sideMenuItemsParams{
			id: "codex.locations",
			// size:  params.Size,
			table: params.Codex.LocationsTable,
			form:  getLocationForm,

			// Logger: params.Codex.AddLocation,
		}),
		newMenuItem(sideMenuItemsParams{
			id: "codex.npcs",
			// size:  params.Size,
			table: params.Codex.NpcsTable,
			form:  getNpcForm,
			// Logger: addNPCToCodex,
		}),
		newMenuItem(sideMenuItemsParams{
			id: "codex.monsters",
			// size:  params.Size,
			table: params.Codex.MonstersTable,
			form:  getMonsterForm,
			// Logger: addNPCToCodex,
		}),
		newMenuItem(sideMenuItemsParams{
			id: "codex.objects",
			// size:  params.Size,
			table: params.Codex.ObjectsTable,
			form:  getObjectForm,
			// Logger: params.Codex.AddObject,
		}),
		newMenuItem(sideMenuItemsParams{
			id: "codex.objectives",
			// size:  params.Size,
			table: params.Codex.ObjectifsTable,
			form:  getObjectifForm,
			// Logger: params.Codex.AddObjective,
		}),
	}

	items := make([]list.Item, 0, len(pages))

	for _, page := range pages {
		items = append(items, tui.NewItem(t.Localize(page.id), "", page))
	}

	return items
}

// -- Getters & Setters -- //
func (m *codexMenu) getCurrentPage() *codexMenuItem {
	for _, item := range m.list.Items() {
		page, ok := item.(tui.Item[*codexMenuItem])
		if !ok {
			continue
		}
		if page.Value().isOpen {
			return page.Value()
		}
	}
	return nil
}

func (m *codexMenu) getSelectedPage() *codexMenuItem {
	if item, ok := m.list.SelectedItem().(tui.Item[*codexMenuItem]); ok {
		return item.Value()
	}
	return nil
}

func (m *codexMenuItem) setShowForm(show bool) {
	m.showForm = show
}

func (m *codexMenuItem) toggleShowForm() {
	m.showForm = !m.showForm
}

// -- Helper -- //
func (m *codexMenu) handleOpenPage(selected *codexMenuItem) {
	for _, item := range m.list.Items() {
		item, ok := item.(tui.Item[*codexMenuItem])

		if page := item.Value(); ok && page != selected {
			page.isOpen = false
		}

	}
	selected.isOpen = true
}

func (m *codexMenu) handleKeyEsc(params sidemenu.UpdateParams) (tea.Model, tea.Cmd) {
	currentPage := m.getCurrentPage()

	if currentPage == nil {
		return params.Delegate(params.Msg)
	}

	if currentPage.showForm {
		currentPage.toggleShowForm()
		return params.Model, sendMsg()
	}

	if currentPage.isOpen {
		currentPage.isOpen = false
		return params.Model, sendMsg()
	}

	return params.Model, sendMsg()
}

func (m *codexMenu) delegateInputToForm(params sidemenu.UpdateParams) (tea.Model, tea.Cmd) {
	currentPage := m.getCurrentPage()

	if currentPage == nil {
		return params.Delegate(params.Msg)
	}

	if !currentPage.showForm {
		return params.Delegate(params.Msg)
	}

	_, cmd := currentPage.form.Update(params.Msg)

	return params.Model, tea.Batch(
		cmd,
		sendMsg(),
	)
}
