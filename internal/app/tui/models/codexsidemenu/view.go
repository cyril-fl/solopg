package codexsidemenu

import (
	"solopg/internal/app/game"
	"solopg/internal/app/tui"
	"solopg/internal/infrastructure/t"
	"strings"
)

func (m Model) MenuView() string { return m.menu.View() }
func (m Model) MenuIndex() int   { return m.menu.Index() }

/*
TODO changer par currentParge.View()
*/
func (m Model) View() string {
	switch m.screen {
	case ScreenPage:
		return m.page.View()
	case ScreenForm:
		return m.form.View()
	default:
		return ""
	}
}

// -- page -- //
func (m *Model) selectActivePage() {
	selected, ok := m.menu.SelectedItem().(tui.Item[link])
	if !ok {
		return
	}
	m.active = selected.Value().id
	m.refreshPage()
	m.page.GotoTop()
}

func (m *Model) refreshPage() {
	m.page.SetContent(RenderCodexPage(m.engine, link{id: m.active, name: m.pageTitle()}))
}

func (m Model) pageTitle() string {
	for _, item := range m.menu.Items() {
		selected, ok := item.(tui.Item[link])
		if ok && selected.Value().id == m.active {
			return selected.Value().name
		}
	}
	return ""
}

func (m *Model) OpenPage() bool {
	selected, ok := m.menu.SelectedItem().(tui.Item[link])
	if !ok {
		return false
	}
	m.active = selected.Value().id
	m.screen = ScreenPage
	m.refreshPage()
	m.page.GotoTop()
	return true
}

func RenderCodexPage(engine *game.Engine, link link) string {
	title := link.name
	var entries []string
	if engine != nil && engine.State != nil && engine.State.Codex != nil {
		codexData := engine.State.Codex.EnsureInitialized()
		switch link.id {
		case CodexNPCs:
			entries = codexData.NpcsTable.Summaries()
		case CodexMonsters:
			entries = codexData.MonstersTable.Summaries()
		case CodexLocations:
			entries = codexData.LocationsTable.Summaries()
		case CodexObjects:
			entries = codexData.ObjectsTable.Summaries()
		case CodexObjectifs:
			entries = codexData.ObjectifsTable.Summaries()
		}
	}
	if len(entries) == 0 {
		entries = []string{t.Localize("codex.empty")}
	}
	return title + "\n\n" + strings.Join(entries, "\n\n")
}
