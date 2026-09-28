package tui

import (
	"solopg/app/domain/campaign"
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/domain/card/characters/classes"
	"solopg/app/domain/card/characters/races"
	"solopg/app/domain/card/locations"
	"solopg/app/types/step"

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
	SelectedBuild    []stats.Modifier
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
	if current := m.steps.CurrentStep(); current != nil && current.Resolve != nil {
		return current.Resolve(&m.context, value)
	}

	return nil
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

func (m *model) forwardNextStep() (*model, tea.Cmd) {
	if !m.selectNextStep() {
		return m, tea.Quit
	}

	submodel := m.steps.GetCurrentSubmodel()
	if submodel == nil {
		return m, tea.Quit
	}

	initCmd := submodel.Init()
	if m.size == nil {
		return m, initCmd
	}

	return m, tea.Sequence(initCmd, func() tea.Msg {
		return *m.size
	})
}

func (m model) Init() tea.Cmd {
	if submodel := m.steps.GetCurrentSubmodel(); submodel != nil {
		return submodel.Init()
	}

	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.size = &msg
		return m.handleResize(msg)
	case tea.KeyMsg:
		return m.handleEvent(msg)

	case ResolutionMsg:
		return m.handleResolution(msg)

	default:
		return m.handleSubmodelUpdate(msg)
	}
}

type DelegateUpdateFunc func(msg tea.Msg) (tea.Model, tea.Cmd)

type EscapeSupport interface {
	HandlesEscape() bool
}
