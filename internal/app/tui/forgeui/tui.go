package forgeui

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/characters"

	tea "charm.land/bubbletea/v2"
)

type Ui struct {
	program *tea.Program
}

func New() *Ui {
	return &Ui{
		program: tea.NewProgram(newModel()),
	}
}

func (ui *Ui) CreateCharacter() (*characters.Character, error) {
	res, err := ui.program.Run()
	if err != nil {
		return nil, err
	}

	finalModel, ok := res.(model)
	if !ok {
		return nil, nil
	}

	if finalModel.cancelled {
		return nil, tui.ErrCreationCancelled
	}

	return finalModel.buildCharacter()
}
