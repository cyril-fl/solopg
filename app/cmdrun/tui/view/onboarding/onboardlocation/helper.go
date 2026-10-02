package onboardlocation

import (
	tea "charm.land/bubbletea/v2"
)

func handleWindowResize(m *model, msg tea.WindowSizeMsg) {
	m.reroll.Options.SetSize(msg.Width, max(0, msg.Height-3))
}
