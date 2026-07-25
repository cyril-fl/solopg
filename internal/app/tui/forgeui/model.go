package forgeui

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/attributes"
	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/characters/archetypes"

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

	races := archetypes.ListRaces()
	raceItems := make([]list.Item, 0, len(races))
	for _, race := range races {
		if !race.Playable {
			continue
		}

		raceItems = append(raceItems, raceItem{
			title: race.String(),
			race:  race,
		})
	}

	classes := archetypes.ListClasses()
	classItems := make([]list.Item, 0, len(classes))
	for _, class := range archetypes.ListClasses() {
		if !class.Playable {
			continue
		}

		classItems = append(classItems, classItem{
			title: class.String(),
			class: class,
		})
	}

	raceList := list.New(raceItems, list.NewDefaultDelegate(), 0, 0)
	classList := list.New(classItems, list.NewDefaultDelegate(), 0, 0)
	configureList(&raceList)
	configureList(&classList)

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
	// TODO ajouter les stuff de base en fonct de la class
	case stepConfirm:
		return m.updateConfirm(msg)
	default:
		return m, nil
	}
}

func (m model) buildCharacter() (*characters.Character, error) {
	return characters.New(characters.Template{
		Name:   m.selectedName,
		Rarity: attributes.F,
		Class:  m.selectedClass,
		Race:   m.selectedRace,
	})
}
