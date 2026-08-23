package tui

import tea "charm.land/bubbletea/v2"

func (m model) handleEvent(msg tea.Msg) (model, tea.Cmd) {
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

func (m model) handleProcess(msg tea.Msg) (model, tea.Cmd) {
	current := m.stepList.getCurrentSubmodel()

	if current == nil {
		return m, nil
	}

	var cmd tea.Cmd
	current, cmd = current.Update(msg)

	m.stepList.setCurrentSubmodel(current)

	return m, cmd
}

func (m model) handleResolution(resolution ResolutionMsg) (tea.Model, tea.Cmd) {
	if resolution.Err != nil {
		m.err = NormalizeError(resolution.Err)
		if m.err != nil {
			return m, tea.Quit
		}

		return m, nil
	}

	if !resolution.Completed {
		return m, nil
	}

	if err := m.resolveCurrentStep(resolution.Value); err != nil {
		m.err = err
		return m, tea.Quit
	}

	return m.moveToNextStep()
}