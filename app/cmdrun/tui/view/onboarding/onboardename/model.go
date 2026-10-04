package onboardename

import (
	cmdruntui "solopg/app/cmdrun/tui"
	sharedtui "solopg/app/shared/tui"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type model struct {
	Input     textinput.Model
	Name      string
	Cancelled bool
}

func NewModel() model {
	input := textinput.New()
	input.Focus()

	return model{
		Input:     input,
		Name:      "",
		Cancelled: false,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Input.SetWidth(msg.Width - 2)
		return m, nil
	}

	newInput, cmd := m.Input.Update(msg)

	m.Input = newInput
	m.Name = newInput.Value()

	if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == sharedtui.KEY_ENTER {
		name := strings.TrimSpace(m.Name)
		if name == sharedtui.KEY_EMPTY {
			return m, nil
		}

		m.Name = name

		return m, cmdruntui.SendResolutionMsg(cmdruntui.ResolutionMsg{
			Completed: true,
			Value:     m.Name,
		})
	}

	return m, cmd
}
