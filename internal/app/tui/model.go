package tui

import (
	"solopg/internal/domain/campaign"
	"solopg/internal/domain/card/characters/archetypes/classes"
	"solopg/internal/domain/card/characters/archetypes/races"
	"solopg/internal/domain/card/effects"
	"solopg/internal/domain/card/locations"
	"solopg/types/step"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	steps   *contextStepList
	context Context
	size    *tea.WindowSizeMsg
	err     error
}

type Context struct {
	SelectedSave     *campaign.Campaign
	SelectedName     string
	SelectedRace     *races.Race
	SelectedClass    *classes.Class
	SelectedBuild    []effects.Modifier
	SelectedLocation *locations.Location
}

type contextStepList = step.List[Context]
type contextStep = step.Step[Context]
type Step = contextStep

func newModel(steps *contextStepList) model {
	return model{
		steps: steps,
	}
}

func (m *model) resolveCurrentStep(value any) error {
	list := m.steps
	current := list.CurrentStep()

	isResolvable := current != nil && current.Resolve != nil
	if !isResolvable {
		return nil
	}

	return current.Resolve(&m.context, value)
}

func (m *model) selectNextStep() bool {
	for {
		next := m.steps.NextStep()
		if next == nil {
			return false
		}

		if next.Skip == nil || !next.Skip(&m.context) {
			return true
		}
	}
}

func (m *model) forwardNextStep() (tea.Model, tea.Cmd) {
	if !m.selectNextStep() {
		return *m, tea.Quit
	}

	submodel := m.steps.GetCurrentSubmodel()
	if submodel == nil {
		return *m, tea.Quit
	}

	initCmd := submodel.Init()
	if m.size == nil {
		return *m, initCmd
	}

	return *m, tea.Sequence(initCmd, func() tea.Msg {
		return *m.size
	})
}

func (m model) Init() tea.Cmd {
	submodel := m.steps.GetCurrentSubmodel()
	if submodel == nil {
		return nil
	}

	return submodel.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.size = &size
	}

	var cmd tea.Cmd
	m, cmd = m.handleEvent(msg)
	if cmd != nil {
		return m, cmd
	}

	if resolution, ok := msg.(ResolutionMsg); ok {
		return m.handleResolution(resolution)
	}

	if m.steps.GetCurrentSubmodel() == nil {
		return m, nil
	}

	return m.handleSubmodel(msg)
}

type DelegateUpdateFunc func(msg tea.Msg) (tea.Model, tea.Cmd)
type RefreshViewFunc func(bool)
