package portal

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/gameplay"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

func makeItems() []list.Item {
	items := []list.Item{
		tui.NewItem("Accept", "", true),
		tui.NewItem("Reroll", "", false),
	}

	return items
}

func makeModel() list.Model {
	items := makeItems()

	model := list.New(items, list.NewDefaultDelegate(), 0, 0)
	tui.ConfigureList(&model)

	return model
}

func handleWindowResize(m *model, msg tea.WindowSizeMsg) {
	// The parent TUI has already reserved the footer height. Reserve only
	// the location title and description shown above the list.
	m.choiceList.SetSize(msg.Width, max(0, msg.Height-3))
}

func handleEnterInput(m model) (tea.Model, tea.Cmd) {
	isSelected, ok := m.choiceList.SelectedItem().(tui.Item[bool])
	if !ok {
		return m, nil
	}

	if isSelected.Value() || m.attempt >= 3 {
		return m, func() tea.Msg {
			return tui.ResolutionMsg{Completed: true, Value: &m.drawLocations}
		}
	}

	newLocation, err := gameplay.DrawLocations()
	if err != nil {
		return m, nil
	}
	m.drawLocations = *newLocation

	m.attempt++

	return m, nil
}
