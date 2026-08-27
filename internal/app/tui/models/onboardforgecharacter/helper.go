package onboardforgecharacter

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/effects"
	"solopg/internal/domain/gameplay"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

func makeChoiceItems() []list.Item {
	items := []list.Item{
		tui.NewItem("Accept", "", true),
		tui.NewItem("Reroll", "", false),
	}

	return items
}

func makeChoiceModel() list.Model {
	items := makeChoiceItems()

	model := list.New(items, list.NewDefaultDelegate(), 0, 0)
	tui.ConfigureList(&model)

	return model
}

func makeConfirmationModel(m list.Model) list.Model {
	items := []list.Item{
		tui.NewItem("Accept", "", true),
	}

	model := list.New(items, list.NewDefaultDelegate(), m.Width(), m.Height())
	tui.ConfigureList(&model)

	return model
}

func handleWindowResize(m *model, msg tea.WindowSizeMsg) {
	m.choiceList.SetSize(msg.Width, max(0, msg.Height-3))
}

func handleEnterInput(m model) (tea.Model, tea.Cmd) {
	isSelected, ok := m.choiceList.SelectedItem().(tui.Item[bool])
	if !ok {
		return m, nil
	}

	m.attempt++

	if isSelected.Value() || m.attempt > 3 {
		return m, func() tea.Msg {
			return tui.ResolutionMsg{Completed: true, Value: m.buildsChoice}
		}
	}

	if m.attempt >= 3 {
		m.choiceList = makeConfirmationModel(m.choiceList)
	}

	m.buildsChoice = drowBuild()

	return m, nil
}

func drowBuild() []effects.Modifier {
	oracle, err := getStatGenerationOracle()
	if err != nil {
		return []effects.Modifier{}
	}

	build, err := generateCharacterBuild(oracle)
	if err != nil {
		return []effects.Modifier{}
	}

	return build
}

func getStatGenerationOracle() (*gameplay.Oracle, error) {
	return gameplay.GetOracleByID("stat_generation")
}

func generateCharacterBuild(oracle *gameplay.Oracle) ([]effects.Modifier, error) {
	statsList := effects.ListStats()
	build := []effects.Modifier{}

	for _, stat := range statsList {
		roll, err := gameplay.RollOracle[int](oracle)
		if err != nil {
			return nil, err
		}

		// jsonlog.JsonifiedLog(roll)

		build = append(build, effects.Modifier{
			Stat:  stat,
			Value: roll.Result,
		})
	}
	return build, nil
}
