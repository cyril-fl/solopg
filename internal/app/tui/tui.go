package tui

import tea "charm.land/bubbletea/v2"

type Ui struct {
	list *contextStepList

	program *tea.Program
}

func New() *Ui {
	return &Ui{list: &contextStepList{}}
}

func (ui *Ui) Add(steps ...Step) *Ui {
	ui.list.Steps = append(ui.list.Steps, steps...)
	return ui
}

func (ui *Ui) Insert(steps ...Step) *Ui {
	if len(steps) == 0 {
		return ui
	}

	insertAt := ui.list.CurrentIndex + 1
	if insertAt < 0 {
		insertAt = 0
	}
	if insertAt > len(ui.list.Steps) {
		insertAt = len(ui.list.Steps)
	}

	oldSteps := ui.list.Steps
	newSteps := make([]contextStep, 0, len(oldSteps)+len(steps))

	newSteps = append(newSteps, oldSteps[:insertAt]...)
	newSteps = append(newSteps, steps...)
	newSteps = append(newSteps, oldSteps[insertAt:]...)

	ui.list.Steps = newSteps
	return ui
}

func (ui *Ui) Run() error {
	ui.program = tea.NewProgram(newModel(ui.list))
	_, err := ui.program.Run()
	return err
}
