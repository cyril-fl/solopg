package step

import (
	tea "charm.land/bubbletea/v2"
)

type List[ T any] struct {
	Steps        []Step[ T ]
	CurrentIndex int
}

type Step[ T any] struct {
	Submodel tea.Model
	Resolve  func(ctx *T, value any) error
	Skip     func(ctx *T) bool
}

func (sl *List[ T ]) isOutOfBounds() bool {
	return len(sl.Steps) == 0 || sl.CurrentIndex < 0 || sl.CurrentIndex >= len(sl.Steps)
}

func (sl *List[ T ]) CurrentStep() *Step[ T ] {
	if sl.isOutOfBounds() {
		return nil
	}
	return &sl.Steps[sl.CurrentIndex]
}

func (sl *List[ T ]) NextStep() *Step[ T ] {
	if sl.isOutOfBounds() {
		return nil
	}

	sl.CurrentIndex++

	return sl.CurrentStep()
}

func (sl *List[ T ]) PreviousStep() *Step[ T ] {
	if sl.isOutOfBounds() {
		return nil
	}

	sl.CurrentIndex--

	return sl.CurrentStep()
}

func (sl *List[ T ]) GetCurrentSubmodel() tea.Model {
	current := sl.CurrentStep()
	if current == nil {
		return nil
	}

	submodel := current.Submodel
	if submodel == nil {
		return nil
	}

	return submodel
}

func (sl *List[ T ]) SetCurrentSubmodel(submodel tea.Model) {
	current := sl.CurrentStep()
	if current == nil {
		return
	}

	current.Submodel = submodel
}
