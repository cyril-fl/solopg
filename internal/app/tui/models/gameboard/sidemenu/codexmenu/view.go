package codexmenu

// func (m Model) MenuView() string { return m.menu.View() }
// func (m Model) MenuIndex() int   { return m.menu.Index() }

/*
TODO changer par currentParge.View()
*/
func (m Model) View() string {
	// switch m.screen {
	// case ScreenPage:
	return m.viewport.View()
	// case ScreenForm:
	// return m.form.View()
	// default:
	// return ""
	// }
}

// w
// -- page -- //
// func (m *Model) selectActivePage() {
// 	selected, ok := m.menu.SelectedItem().(tui.Item[link])
// 	if !ok {
// 		return
// 	}
// 	m.active = selected.Value().id
// 	m.refreshPage()
// 	m.viewport.GotoTop()
// }

// func (m *Model) refreshPage() {
// 	m.viewport.SetContent(m.renderCodexPage(link{id: m.active, name: m.pageTitle()}))
// }

// func (m *Model) renderCodexPage(link link) string {
// 	title := link.name
// 	var entries []string

// 	switch link.id {
// 	case CodexNPCs:
// 		entries = m.codex.NpcsTable.Summaries()
// 	case CodexMonsters:
// 		entries = m.codex.MonstersTable.Summaries()
// 	case CodexLocations:
// 		entries = m.codex.LocationsTable.Summaries()
// 	case CodexObjects:
// 		entries = m.codex.ObjectsTable.Summaries()
// 	case CodexObjectifs:
// 		entries = m.codex.ObjectifsTable.Summaries()
// 	}

// 	if len(entries) == 0 {
// 		entries = []string{t.Localize("codex.empty")}
// 	}
// 	return title + "\n\n" + strings.Join(entries, "\n\n")
// }

func (m *Model) GetView() string {
	return m.View()
}
