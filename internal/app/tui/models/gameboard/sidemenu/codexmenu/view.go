package codexmenu

import (
	"solopg/internal/infrastructure/t"

	"charm.land/lipgloss/v2"
)

func (m *CodexMenu) GetMenuView() string {
	title := lipgloss.NewStyle().Bold(true).Render(t.Localize(m.id))
	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		m.list.View(),
	)
}

// func (m Model) View() string {
// 	// switch m.screen {
// 	// case ScreenPage:
// 	return m.viewport.View()
// 	// case ScreenForm:
// 	// return m.form.View()
// 	// default:
// 	// return ""
// 	// }
// }

func (m *CodexMenu) GetView() string {
	return "CODEX MENU VIEW"
}


func (m *CodexMenu) GetFooter() []string {
	// return []string{t.Localize("add"), t.Localize("back_to_chat")}
	return []string{}
}

/*
TODO changer par currentParge.View()
*/


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