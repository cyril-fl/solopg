package onboardlocation

import (
	"fmt"
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/locations"
	"solopg/internal/domain/gameplay"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type model struct {
	attempt       int
	choiceList    list.Model
	drawLocations locations.Location
}

func NewModel() tea.Model {
	locations, err := gameplay.DrawLocations()
	if err != nil {
		fmt.Println("Error drawing locations:", err)
	}

	return model{
		choiceList:    makeModel(),
		attempt:       1,
		drawLocations: *locations,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		handleWindowResize(&m, msg)

	case tea.KeyMsg:
		switch msg.String() {
		case tui.KeyEnter:
			return handleEnterInput(m)
		}
	}

	var cmd tea.Cmd
	m.choiceList, cmd = m.choiceList.Update(msg)
	return m, cmd
}
