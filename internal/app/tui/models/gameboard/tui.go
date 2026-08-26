package gameboard

import (
	"fmt"
	"solopg/internal/app/game"

	tea "charm.land/bubbletea/v2"
)

type Ui struct {
	program *tea.Program
}

type UiParams struct {
	Engine *game.Engine
	OnSave func() error
}

func NewUi(params UiParams) *Ui {
	return &Ui{
		program: tea.NewProgram(NewModel(params)),
	}
}

func (ui *Ui) Start() error {
	res, err := ui.program.Run()
	if err != nil {
		return err
	}

	finalModel, ok := res.(model)
	if !ok {
		return fmt.Errorf("unexpected model type %T", res)
	}

	_ = finalModel

	return nil
}
