package onboardforgecharacter

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/effects"
	"solopg/internal/domain/gameplay"

	tea "charm.land/bubbletea/v2"
)

/* TODO factoriser avec onboardlocation/helper.go
mais faire attention a la perte de controle du flux
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

func drowBuild() ([]effects.Modifier, error) {
	oracle, err := getStatGenerationOracle()
	if err != nil {
		return nil, err
	}

	build, err := generateCharacterBuild(oracle)
	if err != nil {
		return nil, err
	}

	return build, nil
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
