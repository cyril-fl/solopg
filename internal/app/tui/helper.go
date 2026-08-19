package tui

import tea "charm.land/bubbletea/v2"

func handleEvent(m model, msg tea.Msg) (model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
	case tea.KeyMsg:
		switch msg.String() {
		case KeyQuit:
			return m, tea.Quit
		}
	}

	return m, nil
}

func handleProcess(m model, msg tea.Msg) (model, tea.Cmd) {
	current := m.current()

	if current == nil {
		return m, nil
	}

	var cmd tea.Cmd
	current, cmd = current.Update(msg)

	m.steps[m.step].Model = current

	return m, cmd
}
