package tui

import (
	"solopg/internal/domain/campaign"
	"solopg/internal/domain/card/characters/archetypes/classes"
	"solopg/internal/domain/card/characters/archetypes/races"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	step  step
	steps []Step
	context Context
}

type Context struct {
	SelectedSave *campaign.Campaign
	SelectedName string
	SelectedRace *races.Race	
	SelectedClass *classes.Class
}

func newModel(steps []Step) model {
	return model{
		step:  step1,
		steps: steps,
	}
}

func (m model) current() tea.Model {
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

	m, cmd = handleEvent(m, msg)
	if cmd != nil {
		return m, cmd
	}

	if m.current() == nil {
		return m, nil
	}
m, cmd = changeStep(m, msg)

m, processCmd := handleProcess(m, msg)

return m, tea.Sequence(cmd, processCmd)
}
