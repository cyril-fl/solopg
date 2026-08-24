package portal

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

	cancelled      bool
	emitResolution bool
}

func newModel(emitResolution bool) model {
	locations, err := gameplay.DrawLocations()
	if err != nil {
		fmt.Println("Error drawing locations:", err)
	}
	if locations == nil {
		return model{choiceList: makeModel(), cancelled: true, emitResolution: emitResolution}
	}
	return model{
		choiceList:     makeModel(),
		attempt:        0,
		drawLocations:  *locations,
		cancelled:      false,
		emitResolution: emitResolution,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// Reserve space for the location header and the parent's footer.
		m.choiceList.SetSize(msg.Width, max(0, msg.Height-6))
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case tui.KeyQuit:
			m.cancelled = true
			return m, tea.Quit

		case tui.KeyEnter:
			if selected, ok := m.choiceList.SelectedItem().(tui.Item[bool]); ok {
				if selected.Value() && m.attempt < 3 {
					m.attempt++
					newLocation, err := gameplay.DrawLocations()
					if err != nil {
						return m, nil
					}
					m.drawLocations = *newLocation
				} else {
					if m.emitResolution {
						location := m.drawLocations
						return m, func() tea.Msg {
							return tui.ResolutionMsg{Completed: true, Value: &location}
						}
					}
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
