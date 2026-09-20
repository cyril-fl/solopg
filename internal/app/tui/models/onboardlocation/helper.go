package onboardlocation

import (
	"solopg/internal/app/tui"

	tea "charm.land/bubbletea/v2"
)

/*
	TODO LOW Factoriser avec onboardforgecharacter/helper.go mais faire attention a la perte de controle du flux
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
