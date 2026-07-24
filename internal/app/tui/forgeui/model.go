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
	selectedRace  archetypes.OG_Race
	selectedClass archetypes.OG_Class

	cancelled bool
}

func newModel() model {
	nameInput := textinput.New()
	nameInput.Placeholder = "Enter your name"
	nameInput.Focus()

	raceItems := make([]list.Item, 0, len(archetypes.PlayableRaces))
	for _, race := range archetypes.PlayableRaces {
		raceItems = append(raceItems, raceItem{
			title: race.String(),
			race:  race,
		})
	}

	classItems := make([]list.Item, 0, len(archetypes.PlayableClasses))
	for _, class := range archetypes.PlayableClasses {
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
	case stepConfirm:
		return m.updateConfirm(msg)
	default:
		return m, nil
	}
}

func (m model) buildCharacter() (*characters.Character, error) {
	return characters.NewCharacter(characters.CharacterTemplate{
		Name:   m.selectedName,
		Rarity: attributes.F,
		Class:  m.selectedClass,
		Race:   m.selectedRace,
	})
}
