package onboardforgecharacter

import (
	"solopg/internal/app/tui"

	"solopg/internal/domain/card/attributes/stats"
	"solopg/internal/domain/gameplay/oracle"

	tea "charm.land/bubbletea/v2"
)

/*
	TODO LOW Factoriser avec onboardlocation/helper.go mais faire attention a la perte de controle du flux
*/

func handleWindowResize(m *model, msg tea.WindowSizeMsg) {
	m.reroll.Options.SetSize(msg.Width, max(0, msg.Height-3))
}

func handleEnterInput(m model) (tea.Model, tea.Cmd) {
	isSelected, ok := m.reroll.Options.SelectedItem().(tui.Item[bool])
	if !ok {
		return m, nil
	}

	if isSelected.Value() || m.reroll.IsOutOfLimit() {
		return m, func() tea.Msg {
			return tui.ResolutionMsg{Completed: true, Value: m.reroll.Value}
		}
	}

	m.reroll.Reroll()

	return m, nil
}

func drowBuild() ([]stats.Modifier, error) {
	oracle, err := getStatGenerationOracle()
	if err != nil {
		return nil, err
	}

	build, err := generateCharacterBuild(*oracle)
	if err != nil {
		return nil, err
	}

	return build, nil
}

func getStatGenerationOracle() (*oracle.Oracle, error) {
	return oracle.GetByID("stat_generation")
}

func generateCharacterBuild(rules oracle.Oracle) ([]stats.Modifier, error) {
	statsList := stats.List()
	build := []stats.Modifier{}

	for _, stat := range statsList {
		roll, err := oracle.Roll[int](rules)
		if err != nil {
			return nil, err
		}

		// jsonlog.JsonifiedLog(roll)

		build = append(build, stats.Modifier{
			Stat:  stat,
			Value: roll.Result,
		})
	}
	return build, nil
}
