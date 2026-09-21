package codexmenu

import (
	"solopg/app/components/form"
	"solopg/app/components/form/field"
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/domain/card/characters/classes"
	"solopg/app/domain/card/characters/races"
	"solopg/app/domain/gameplay/oracle"
	"solopg/app/services/t"
	"solopg/app/tui/models"
	"solopg/app/tui/view/sidemenu"
	"solopg/config"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

// - Side panel - //
func makeCodexList(params CodexMenuParams) []list.Item {
	pages := []*codexMenuItem{
		newMenuItem(sideMenuItemsParams{
			id:    "codex.locations",
			table: params.Codex.LocationsTable,
			form:  getLocationForm,
		}),
		newMenuItem(sideMenuItemsParams{
			id:    "codex.npcs",
			table: params.Codex.NpcsTable,
			form:  getNpcForm,
		}),
		newMenuItem(sideMenuItemsParams{
			id:    "codex.beastiary",
			table: params.Codex.BeastiaryTable,
			form:  getBeastiaryForm,
		}),
		newMenuItem(sideMenuItemsParams{
			id:    "codex.objects",
			table: params.Codex.ObjectsTable,
			form:  getObjectForm,
		}),
		newMenuItem(sideMenuItemsParams{
			id:    "codex.objectives",
			table: params.Codex.ObjectifsTable,
			form:  getObjectifForm,
		}),
	}

	items := make([]list.Item, 0, len(pages))

	for _, page := range pages {
		items = append(items, models.NewItem(t.Localize(page.id), "", page))
	}

	return items
}

// - Getters & Setters - //
// Codex menu
func (m *codexMenu) getCurrentPage() *codexMenuItem {
	for _, item := range m.list.Items() {

		if page, ok := item.(models.Item[*codexMenuItem]); !ok {
			continue
		} else if page.Value().isOpen {
			return page.Value()
		}
	}
	return nil
}

func (m *codexMenu) getSelectedPage() *codexMenuItem {
	if item, ok := m.list.SelectedItem().(models.Item[*codexMenuItem]); ok {
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

// - Helper - //
// Input
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
			return params.Model, form.SendMsg[form.Validate]()
		}

		currentPage.form = currentPage.newForm()
		currentPage.setShowForm(true)
		return params.Model, sendMsg()
	}

	if selectedPage := m.getSelectedPage(); selectedPage != nil {
		m.handleOpenPage(selectedPage)
	}

	return params.Model, sendMsg()
}

// Page
func (m *codexMenu) handleOpenPage(selected *codexMenuItem) {
	for _, item := range m.list.Items() {
		item, ok := item.(models.Item[*codexMenuItem])

		if page := item.Value(); ok && page != selected {
			page.setOpen(false)
		}

	}
	selected.setOpen(true)
}

func (m *codexMenu) updatePageOnRedirection() {
	currentPage := m.getCurrentPage()

	if currentPage != nil && m.IsOpen() {
		selected := m.getSelectedPage()
		m.handleOpenPage(selected)
	}
}

// Form
func getFormRaceOptions() []models.Item[string] {
	var raceItems []models.Item[string]
	for _, i := range races.List() {
		raceItems = append(raceItems, models.NewItem(i.GetName(), "", i.GetName()))
	}
	return raceItems
}

func getFormClassOptions() []models.Item[string] {
	var classItems []models.Item[string]
	for _, i := range classes.List() {
		classItems = append(classItems, models.NewItem(i.GetName(), "", i.GetName()))
	}
	classItems = append(classItems, models.NewItem("None", "", ""))
	return classItems
}

func getFormOracle() []models.Item[string] {
	list, err := oracle.ListFromFolder(config.Current.Documents.Folders.Encounters)
	if err != nil {
		return nil
	}

	var oracleItems []models.Item[string]
	for _, i := range list {
		oracleItems = append(oracleItems, models.NewItem(t.Localize("oracles."+i), "", i))
	}

	return oracleItems
}

func getFormStatsFields() []form.Field {
	fields := []form.Field{}

	for _, i := range stats.List() {
		f := field.TextField(field.TextTemplate[string]{
			ID:           i.String(),
			Label:        t.Localize(i.String()),
			Validator:    nil,
			Required:     false,
			Defaultvalue: "",
		})
		fields = append(fields, f)
	}

	return fields
}

func (m *codexMenu) handleFormSubmit(params sidemenu.UpdateParams) (tea.Model, tea.Cmd) {
	currentForm := m.GetFormFromCurrentPage()
	if currentForm == nil {
		return params.Model, nil
	}

	currentForm.Submit()

	if currentForm.HasErrors() {
		return params.Model, form.SendErrorMsg(currentForm.GetError())
	}

	return params.Model, m.handleFormPost()
}

func (m *codexMenu) handleFormPost() tea.Cmd {
	currentPage := m.getCurrentPage()
	if currentPage == nil {
		return nil
	}

	currentForm := currentPage.form
	if currentForm == nil {
		return nil
	}

	err := currentPage.table.AddFromMappedValues(currentForm.GetValuesAsMappedString())
	if err != nil {
		currentPage.form.SetError(err)
		return form.SendErrorMsg(err)
	}

	currentPage.setShowForm(false)
	// currentForm.Reset()
	/*
		TODO LOW log l'enregistrement du codex ca dans le main view et les log
		verrifier que tout est persister
	*/
	return sendMsg()
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
