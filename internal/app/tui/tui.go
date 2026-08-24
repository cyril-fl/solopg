package tui

import (
	tea "charm.land/bubbletea/v2"
)

type Ui struct {
	steps   *contextStepList
	program *tea.Program
}

func New() *Ui {
	return &Ui{steps: &contextStepList{}}
}

func (ui *Ui) Add(steps ...Step) *Ui {
	ui.steps.Steps = append(ui.steps.Steps, steps...)
	return ui
}

func (ui *Ui) Insert(steps ...Step) *Ui {
	if len(steps) == 0 {
		return ui
	}

	insertAt := ui.steps.CurrentIndex + 1
	if insertAt < 0 {
		insertAt = 0
	}
	if insertAt > len(ui.steps.Steps) {
		insertAt = len(ui.steps.Steps)
	}

	oldSteps := ui.steps.Steps
	newSteps := make([]contextStep, 0, len(oldSteps)+len(steps))

	newSteps = append(newSteps, oldSteps[:insertAt]...)
	newSteps = append(newSteps, steps...)
	newSteps = append(newSteps, oldSteps[insertAt:]...)

	ui.steps.Steps = newSteps
	return ui
}

func (ui *Ui) Run() error {
	ui.program = tea.NewProgram(newModel(ui.steps))
	_, err := ui.program.Run()
	return err
}
