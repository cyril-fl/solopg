package forgeui

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/characters/archetypes/classes"
	"solopg/internal/domain/card/characters/archetypes/races"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type step int

const (
	stepName step = iota
	stepRace
	stepClass
	stepConfirm
)

func (m model) updateName(msg tea.Msg) (tea.Model, tea.Cmd) {
	updatedInput, cmd := m.nameInput.Update(msg)
	m.nameInput = updatedInput

	if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == tui.KeyEnter {
		name := strings.TrimSpace(m.nameInput.Value())
		if name == "" {
			return m, nil
		}
		m.selectedName = name
		m.step = stepRace
		m.raceList.ResetSelected()
		return m, nil
	}

	return m, cmd
}

func (m model) updateRace(msg tea.Msg) (tea.Model, tea.Cmd) {
	updatedList, cmd := m.raceList.Update(msg)
	m.raceList = updatedList

	if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == tui.KeyEnter {
		selected, ok := m.raceList.SelectedItem().(tui.Item[*races.Race])
		if !ok {
			return m, nil
		}
		m.selectedRace = selected.Value().GetName()
		m.step = stepClass
		m.classList.ResetSelected()
		return m, nil
	}

	return m, cmd
}

func (m model) updateClass(msg tea.Msg) (tea.Model, tea.Cmd) {
	updatedList, cmd := m.classList.Update(msg)
	m.classList = updatedList

	if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == tui.KeyEnter {
		selected, ok := m.classList.SelectedItem().(tui.Item[*classes.Class])
		if !ok {
			return m, nil
		}
		m.selectedClass = selected.Value().GetName()
		m.step = stepConfirm
		return m, nil
	}

	return m, cmd
}

func (m model) updateConfirm(msg tea.Msg) (tea.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == tui.KeyEnter {
		return m, tea.Quit
	}

	return m, nil
}
