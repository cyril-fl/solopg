package portalui

import (
	"fmt"
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/locations"
	"solopg/internal/domain/gameplay"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type model struct {
	attempt          int
	choiceList       list.Model
	selectedLocation locations.Location

	cancelled bool
}

func newModel() model {
	choiceList := []list.Item{
		choiceItem{title: "Accept", retry: false},
		choiceItem{title: "Reroll", retry: true},
	}

	choiceListModel := list.New(choiceList, list.NewDefaultDelegate(), 0, 0)
	configureList(&choiceListModel)

	selectedLocation, err := gameplay.DrawLocations()
	if err != nil {
		fmt.Println("Error drawing locations:", err)
	}

	return model{
		choiceList:       choiceListModel,
		attempt:          0,
		selectedLocation: *selectedLocation,
		cancelled:        false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.choiceList.SetSize(msg.Width, max(0, msg.Height-5))
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case tui.KeyCtrlC, tui.KeyQuit:
			m.cancelled = true
			return m, tea.Quit

		case tui.KeyEnter:
			if selected, ok := m.choiceList.SelectedItem().(choiceItem); ok {
				if selected.retry && m.attempt < 3 {
					m.attempt++
					newLocation, err := gameplay.DrawLocations()
					if err != nil {
						fmt.Println("Error drawing locations:", err)
						return m, nil
					}
					m.selectedLocation = *newLocation
				} else {
					return m, tea.Quit
				}
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.choiceList, cmd = m.choiceList.Update(msg)

	return m, cmd
}
