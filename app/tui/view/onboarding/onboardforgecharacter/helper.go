package onboardforgecharacter

import (
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/services/process/generatestats"

	tea "charm.land/bubbletea/v2"
)

func handleWindowResize(m *model, msg tea.WindowSizeMsg) {
	m.reroll.Options.SetSize(msg.Width, max(0, msg.Height-3))
}

func drowBuild() ([]stats.Modifier, error) {
	generator := generatestats.NewValuelessGenerator()
	generator.GenerateFromOracle()

	if generator.HasError() {
		return nil, generator.GetError()
	}

	return generator.GetModifiers(), nil
}