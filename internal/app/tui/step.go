package tui

import (
	tea "charm.land/bubbletea/v2"
)

type stepList struct {
	Steps        []Step
	CurrentIndex int
}

type Step struct {
	Submodel tea.Model
	Resolve  func(ctx *Context, value any) error
	Skip     func(ctx *Context) bool
}

func (sl *stepList) isOutOfBounds() bool {
	return len(sl.Steps) == 0 || sl.CurrentIndex < 0 || sl.CurrentIndex >= len(sl.Steps)
}

func (sl *stepList) currentStep() *Step {
	if sl.isOutOfBounds() {
		return nil
	}
	return &sl.Steps[sl.CurrentIndex]
}

func (sl *stepList) nextStep() *Step {
	if sl.isOutOfBounds() {
		return nil
	}

	sl.CurrentIndex++

	return sl.currentStep()
}

func (sl *stepList) previousStep() *Step {
	if sl.isOutOfBounds() {
		return nil
	}

	sl.CurrentIndex--

	return sl.currentStep()
}

func (sl *stepList) getCurrentSubmodel() tea.Model {
	current := sl.currentStep()
	if current == nil {
		return nil
	}

	submodel := current.Submodel
	if submodel == nil {
		return nil
	}

	return submodel
}

func (sl *stepList) setCurrentSubmodel(submodel tea.Model) {
	current := sl.currentStep()
	if current == nil {
		return
	}

	current.Submodel = submodel
}
