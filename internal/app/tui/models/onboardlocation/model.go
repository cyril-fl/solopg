package onboardlocation

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/locations"
	"solopg/internal/domain/gameplay/portal"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	reroll tui.RerollModel[*locations.Location]
}

func NewModel() tea.Model {
	return model{
		reroll: tui.NewRerollModel(
			tui.NewOptionsModel(tui.DefaultRerollOptions),
			portal.Teleport,
			3,
		),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		handleWindowResize(&m, msg)

	case tea.KeyMsg:
		switch msg.String() {
		case tui.KeyEnter:
			return handleEnterInput(m)
		}
	}

	var cmd tea.Cmd
	m.reroll.Options, cmd = m.reroll.Options.Update(msg)
	return m, cmd
}
