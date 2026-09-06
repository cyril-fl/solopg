package codexsidemenu

import (
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/models/codexform"

	tea "charm.land/bubbletea/v2"
)

func (m *Model) updateForm(msg tea.Msg) tea.Cmd {
	key, ok := msg.(tea.KeyPressMsg)
	if ok && key.String() == tui.KeyEsc {
		m.screen = ScreenPage
		return nil
	}
	if ok && key.String() == tui.KeyEnter {
		if m.form.AdvanceOnEnter() {
			return nil
		}
		result, err := m.form.Submit()
		if err != nil {
			m.form.SetError(err)
			return nil
		}
		if err := addCodexEntry(m.engine, result); err != nil {
			m.form.SetError(err)
			return nil
		}
		m.screen = ScreenPage
		m.refreshPage()
		m.page.GotoTop()
		return nil
	}
	m.form.SetError(nil)
	return m.form.Update(msg)
}

func (m Model) formKind() codexform.Kind {
	switch m.active {
	case CodexNPCs:
		return codexform.NPCs
	case CodexMonsters:
		return codexform.Monsters
	case CodexLocations:
		return codexform.Locations
	case CodexObjects:
		return codexform.Objects
	default:
		return codexform.Objectifs
	}
}