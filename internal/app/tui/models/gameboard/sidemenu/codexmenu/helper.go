package codexmenu

import (
	"fmt"
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/models/gameboard/sidemenu"
	"solopg/internal/infrastructure/t"
	"solopg/internal/platform/message"

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
			addEntryToCodex: func(message string) {
				fmt.Println("Adding location to codex:", message)
				// params.Codex.AddLocation,
			},
		}),
		newMenuItem(sideMenuItemsParams{
			id: "codex.npcs",
			// size:  params.Size,
			table: params.Codex.NpcsTable,
			form:  getNpcForm,
			addEntryToCodex: func(message string) {
				fmt.Println("Adding NPC to codex:", message)
				// params.Codex.AddNPC,
			},
		}),
		newMenuItem(sideMenuItemsParams{
			id: "codex.monsters",
			// size:  params.Size,
			table: params.Codex.MonstersTable,
			form:  getMonsterForm,
			addEntryToCodex: func(message string) {

				fmt.Println("Adding Monster to codex:", message)
				// params.Codex.AddMonster,
			},
		}),
		newMenuItem(sideMenuItemsParams{
			id: "codex.objects",
			// size:  params.Size,
			table: params.Codex.ObjectsTable,
			form:  getObjectForm,
			addEntryToCodex: func(message string) {
				fmt.Println("Adding Item to codex:", message)
				// params.Codex.AddObject,
			},
		}),
		newMenuItem(sideMenuItemsParams{
			id: "codex.objectives",
			// size:  params.Size,
			table: params.Codex.ObjectifsTable,
			form:  getObjectifForm,
			addEntryToCodex: func(message string) {
				fmt.Println("Adding Objective to codex:", message)
				// params.Codex.AddObjective,
			},
		}),
	}

	items := make([]list.Item, 0, len(pages))

	for _, page := range pages {
		items = append(items, tui.NewItem(t.Localize(page.id), "", page))
	}

	return items
}

// -- Getters & Setters -- //
// Codex menu
func (m *codexMenu) getCurrentPage() *codexMenuItem {
	for _, item := range m.list.Items() {

		if page, ok := item.(tui.Item[*codexMenuItem]); !ok {
			continue
		} else if page.Value().isOpen {
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

// Codex menu iteme
func (m *codexMenuItem) setOpen(open bool) {
	if open {
		m.isOpen = true
		return
	}

	m.isOpen = false
	m.showForm = false
}

func (m *codexMenuItem) setShowForm(show bool) {
	m.showForm = show
}

func (m *codexMenuItem) toggleShowForm() {
	m.showForm = !m.showForm
}

// -- Helper -- //
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

func (m *codexMenu) handleKeyShiftEnter(params sidemenu.UpdateParams) (tea.Model, tea.Cmd) {
	currentPage := m.getCurrentPage()

	if currentPage != nil && m.IsOpen() {
		if currentPage.showForm {
			return params.Model, message.SendFormMsg[message.FormSubmit]()
		}

		currentPage.setShowForm(true)
		return params.Model, sendMsg()
	}

	if selectedPage := m.getSelectedPage(); selectedPage != nil {
		m.handleOpenPage(selectedPage)
	}

	return params.Model, sendMsg()
}

func (m *codexMenu) handleOpenPage(selected *codexMenuItem) {
	for _, item := range m.list.Items() {
		item, ok := item.(tui.Item[*codexMenuItem])

		if page := item.Value(); ok && page != selected {
			page.setOpen(false)
		}

	}
	selected.setOpen(true)
}

func (m *codexMenu) handleFormSubmit(params sidemenu.UpdateParams) (tea.Model, tea.Cmd) {
	form := m.GetFormFromCurrentPage()
	if form == nil {
		return params.Model, nil
	}

	form.Validate()

	if form.HasErrors() {
		return params.Model, message.SendFormMsg[message.FormError]()
	}

	return params.Model, m.handleFormPost()
}

func (m *codexMenu) handleFormPost() tea.Cmd {
	fmt.Println("Form submitted successfully.")

	currentPage := m.getCurrentPage()
	if currentPage == nil {
		return nil
	}

	form := currentPage.form
	if form == nil {
		return nil
	}

	currentPage.addEntryToCodex(form.GetValuesAsString())

	return sendMsg()
}

func (m *codexMenu) updatePageOnRedirection() {
	currentPage := m.getCurrentPage()

	if currentPage != nil && m.IsOpen() {
		selected := m.getSelectedPage()
		m.handleOpenPage(selected)
	}
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
