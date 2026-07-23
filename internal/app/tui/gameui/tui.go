package gameui

import (
	"solopg/internal/app/game"

	tea "charm.land/bubbletea/v2"
)

func initialModel() model {
	return model{
		itemList: newMenuList(),
	}
}

type Ui struct {
	engine  *game.Engine
	program *tea.Program
}

func NewUi(engine *game.Engine) *Ui {
	return &Ui{
		engine: engine,
	}
}

func (ui *Ui) Start() error {
	ui.program = tea.NewProgram(initialModel())

	_, err := ui.program.Run()
	return err
}
