package tui

import (
	tea "charm.land/bubbletea/v2"
)

type Ui struct {
	program *tea.Program
}

type ModelList []contextStep

func New(subsystem ModelList) *Ui {
	return &Ui{
		program: tea.NewProgram(newModel(subsystem)),
	}
}

func (ui *Ui) Run() error {
	_, err := ui.program.Run()
	if err != nil {
		return err
	}

	return nil
}
