package onboardforgecharacter

import (
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/tui"
	"solopg/app/tui/models"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	reroll models.RerollModel[[]stats.Modifier]
}

func NewModel() tea.Model {
	return model{
		reroll: models.NewRerollModel(
			models.NewOptionsModel(models.DefaultRerollOptions),
			drowBuild,
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
			return m, m.reroll.HandleEnterInput()
		}
	}

	var cmd tea.Cmd
	m.reroll.Options, cmd = m.reroll.Options.Update(msg)
	return m, cmd
}
