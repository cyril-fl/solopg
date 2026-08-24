package portal

import (
	"fmt"
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/locations"

	tea "charm.land/bubbletea/v2"
)

type Ui struct {
	program *tea.Program
}

func NewUi() *Ui {
	return &Ui{
		program: tea.NewProgram(newModel(false)),
	}
}

// NewModel returns a portal view that can be embedded in the parent TUI.
func NewModel() tea.Model {
	return newModel(true)
}

func (ui *Ui) SelectLocation() (*locations.Location, error) {
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

	return &finalModel.drawLocations, nil
}
