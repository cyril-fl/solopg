package forgeui

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/characters/archetypes/classes"
	"solopg/internal/domain/card/characters/archetypes/races"

	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type model struct {
	step step

	nameInput textinput.Model
	raceList  list.Model
	classList list.Model

	selectedName  string
	selectedRace  string
	selectedClass string

	cancelled bool
}

func newModel() model {
	nameInput := textinput.New()
	nameInput.Placeholder = "Enter your name"
	nameInput.Focus()

	races := races.List()
	raceList := makeModel(races)

	classes := classes.List()
	classList := makeModel(classes)

	return model{
		step:      stepName,
		nameInput: nameInput,
		raceList:  raceList,
		classList: classList,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.nameInput.Focus())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.nameInput.SetWidth(msg.Width - 2)
		m.raceList.SetSize(msg.Width, msg.Height-2)
		m.classList.SetSize(msg.Width, msg.Height-2)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case tui.KeyCtrlC, tui.KeyQuit:
			m.cancelled = true
			return m, tea.Quit
		}
	}

	switch m.step {
	case stepName:
		return m.updateName(msg)
	case stepRace:
		return m.updateRace(msg)
	case stepClass:
		return m.updateClass(msg)
	// TODO Rajouter une step pour chaque stats en fonction de la race
	case stepConfirm:
		return m.updateConfirm(msg)
	default:
		return m, nil
	}
}
