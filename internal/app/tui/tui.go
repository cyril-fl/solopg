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

	index := ui.steps.CurrentIndex + 1
	if index < 0 {
		index = 0
	}
	if index > len(ui.steps.Steps) {
		index = len(ui.steps.Steps)
	}

	ui.steps.Steps = append(ui.steps.Steps, make([]contextStep, len(steps))...)
	copy(ui.steps.Steps[index+len(steps):], ui.steps.Steps[index:])
	copy(ui.steps.Steps[index:index+len(steps)], steps)

	return ui
}

func (ui *Ui) Run() error {
	ui.program = tea.NewProgram(newModel(ui.steps))
	_, err := ui.program.Run()
	return err
}
