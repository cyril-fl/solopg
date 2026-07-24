package forgeui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	switch m.step {
	case stepName:
		return tea.NewView(m.renderName())
	case stepRace:
		return tea.NewView(m.renderRace())
	case stepClass:
		return tea.NewView(m.renderClass())
	case stepConfirm:
		return tea.NewView(m.renderConfirm())
	default:
		return tea.NewView("")
	}
}

func (m model) renderName() string {
	return fmt.Sprintf("Create character\n\nName: %s\n\nenter: continue  q: quit", m.nameInput.View())
}

func (m model) renderRace() string {
	return fmt.Sprintf("Create character\n\nName: %s\n\nChoose race:\n%s\n\nenter: validate  q: quit", m.selectedName, m.raceList.View())
}

func (m model) renderClass() string {
	return fmt.Sprintf("Create character\n\nName: %s\nRace: %s\n\nChoose class:\n%s\n\nenter: validate  q: quit", m.selectedName, m.selectedRace, m.classList.View())
}

func (m model) renderConfirm() string {
	return fmt.Sprintf(
		"Create character\n\nName: %s\nRace: %s\nClass: %s\n\nenter: create  q: quit",
		m.selectedName,
		m.selectedRace,
		m.selectedClass,
	)
	// return jsonlog.JsonifiedLog()
}
