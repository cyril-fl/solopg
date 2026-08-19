package tui

import (
	"solopg/internal/domain/campaign"
	"solopg/internal/domain/card/characters/archetypes/classes"
	"solopg/internal/domain/card/characters/archetypes/races"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	step    step
	steps   []Step
	context Context
	err     error
	size    *tea.WindowSizeMsg
}

type Context struct {
	SelectedSave  *campaign.Campaign
	SelectedName  string
	SelectedRace  *races.Race
	SelectedClass *classes.Class
}

func newModel(steps []Step) model {
	return model{
		step:  step1,
		steps: steps,
	}
}

func (m model) current() tea.Model {
	if len(m.steps) == 0 || int(m.step) >= len(m.steps) {
		return nil
	}
	return m.steps[m.step].Model
}

func (m model) Init() tea.Cmd {
	if m.current() == nil {
		return nil
	}

	return m.current().Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.size = &size
	}

	m, cmd = handleEvent(m, msg)
	if cmd != nil {
		return m, cmd
	}

	if resolution, ok := msg.(ResolutionMsg); ok {
		if resolution.Err != nil {
			m.err = NormalizeError(resolution.Err)
			if m.err != nil {
				return m, tea.Quit
			}
			return m, nil
		}

		if resolution.Completed {
			if err := resolveStep(&m, resolution.Value); err != nil {
				m.err = err
				return m, tea.Quit
			}
			return advanceStep(&m)
		}
		return m, nil
	}

	if m.current() == nil {
		return m, nil
	}
	m, cmd = handleProcess(m, msg)
	return m, cmd
}
