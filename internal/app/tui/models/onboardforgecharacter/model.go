package onboardforgecharacter

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/effects"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type model struct {
	attempt       int
	choiceList    list.Model
	buildsChoice		[]effects.Modifier
}

func NewModel() tea.Model {
	return model{
		choiceList:    makeChoiceModel(),
		attempt:       1,
		buildsChoice:  drowBuild(),
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
	m.choiceList, cmd = m.choiceList.Update(msg)
	return m, cmd
}
