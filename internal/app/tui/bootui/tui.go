package bootui

import (
	"fmt"
	"solopg/internal/app/tui"
	"solopg/internal/domain/campaign"

	tea "charm.land/bubbletea/v2"
)

type Ui struct {
	program *tea.Program
}

func NewUi(saves []campaign.Campaign) *Ui {
	return &Ui{
		program: tea.NewProgram(newModel(saves)),
	}
}

func (ui *Ui) SelectSave() (*campaign.Campaign, error) {
	res, err := ui.program.Run()
	if err != nil {
		return nil, err
	}

	finalModel, ok := res.(model)
	if !ok {
		return nil, fmt.Errorf("unexpected model type %T", res)
	}

	if finalModel.cancelled {
		return nil, tui.ErrSelectionCancelled
	}

	return finalModel.selected, nil
}
