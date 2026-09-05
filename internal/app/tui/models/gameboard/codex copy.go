package gameboard

// import (
// 	"solopg/internal/app/tui"
// 	"solopg/internal/app/tui/models/codex"
// 	"solopg/internal/app/tui/models/codexform"

// 	tea "charm.land/bubbletea/v2"
// )	

// type pageID = string 

// const (
// 	page pageID = "page"
// 	form pageID = "form"
// )

// type CodexView struct {
// 	// page   codexViewModel[viewport.Model]
// 	// form        codexViewModel[codexform.Model]
// 	pages 		[]codexViewModel
// 	openedPage 	string
// }

// func (cv CodexView) openPage(id string) {
// 	cv.openedPage = id
// }

// type codexViewModel[T any] struct {
// 	id 	string
// 	model   T
// 	// open     bool
// }

// func newCodexView[T any](id string, model T) (codexViewModel[T]) {
// 	return codexViewModel[T]{
// 		id: id,
// 		model: model,
// 	}
// }


// func OpenCodexPage(m model) (model, tea.Cmd) {
// 	selected, ok := m.codexList.SelectedItem().(tui.Item[codex.CodexLink])
// 	if !ok {
// 		return m, nil
// 	}

// 	link := selected.Value()
// 	m.codexView.openedPage(page)
// 	m.codexView.page.model.SetWidth(m.viewport.Width())
// 	// The Codex page gets one additional footer line (Escape: retour au chat).
// 	// Reserve that line immediately; the next resize event will recalculate it.
// 	m.codexView.page.model.SetHeight(m.viewport.Height() + m.textarea.Height())
// 	m.codexView.page.model.SetContent(codex.RenderCodexPage(m.engine, link))
// 	m.codexView.page.model.GotoTop()
// 	return m, nil
// }

// func HandleCodexPageNavigation(m model, msg tea.KeyPressMsg) (model, tea.Cmd) {
// 	previousIndex := m.codexList.Index()
// 	var cmd tea.Cmd
// 	m.codexList, cmd = m.codexList.Update(msg)
// 	if m.codexList.Index() != previousIndex {
// 		selected, ok := m.codexList.SelectedItem().(tui.Item[codex.CodexLink])
// 		if ok {
// 			m.codexView.page.model.SetContent(codex.RenderCodexPage(m.engine, selected.Value()))
// 			m.codexView.page.model.GotoTop()
// 		}
// 	}
// 	return m, cmd
// }

// func OpenCodexForm(m model) (model, tea.Cmd) {
// 	selected, ok := m.codexList.SelectedItem().(tui.Item[codex.CodexLink])
// 	if !ok {
// 		return m, nil
// 	}

// 	var kind codexform.Kind
// 	switch selected.Value().Kind {
// 	case codex.CodexNPCs:
// 		kind = codexform.NPCs
// 	case codex.CodexMonsters:
// 		kind = codexform.Monsters
// 	case codex.CodexLocations:
// 		kind = codexform.Locations
// 	case codex.CodexObjects:
// 		kind = codexform.Objects
// 	case codex.CodexObjectifs:
// 		kind = codexform.Objectifs
// 	}

// 	m.codexView.form.model = codexform.New(kind, m.viewport.Width())
// 	m.codexView.form.open = true
// 	return m, nil
// }

// func HandleCodexFormKey(m model, key tea.KeyPressMsg) (model, tea.Cmd) {
// 	if key.String() == tui.KeyEsc {
// 		m.codexView.form.open = false
// 		return m, nil
// 	}
// 	if key.String() == tui.KeyEnter {
// 		if m.codexView.form.model.AdvanceOnEnter() {
// 			return m, nil
// 		}
// 		result, err := m.codexView.form.model.Submit()
// 		if err != nil {
// 			m.codexView.form.model.SetError(err)
// 			return m, nil
// 		}
// 		if err := codex.AddCodexEntry(m.engine, result); err != nil {
// 			m.codexView.form.model.SetError(err)
// 			return m, nil
// 		}

// 		m.codexView.form.open = false
// 		m.RefreshCodexPage()
// 		return m, nil
// 	}

// 	m.codexView.form.model.SetError(nil)
// 	return m, m.codexView.form.model.Update(key)
// }

// func (m *model) RefreshCodexPage() {
// 	selected, ok := m.codexList.SelectedItem().(tui.Item[codex.CodexLink])
// 	if !ok {
// 		return
// 	}
// 	m.codexView.page.model.SetContent(codex.RenderCodexPage(m.engine, selected.Value()))
// 	m.codexView.page.model.GotoTop()
// }
