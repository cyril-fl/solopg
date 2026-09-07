package codexform

import (
	"solopg/internal/app/tui"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type Kind string

const (
	NPCs      Kind = "NPCs"
	Monsters  Kind = "Monsters"
	Locations Kind = "Locations"
	Objects   Kind = "Objects"
	Objectifs Kind = "Objectifs"
)

type Result struct {
	Kind   Kind
	Values map[string]string
}

type field struct {
	name        string
	label       string
	input       textinput.Model
	choices     []string
	choiceIndex int
}

type Model struct {
	kind   Kind
	fields []field
	index  int
	err    string
}

func New(kind Kind, width int) Model {
	names := fieldsFor(kind)
	fields := make([]field, 0, len(names))
	for _, item := range names {
		input := textinput.New()
		input.SetWidth(max(10, width-18))
		input.Prompt = "> "
		current := field{name: item[0], label: item[1], input: input, choices: choicesFor(kind, item[0])}
		if len(current.choices) > 0 {
			current.input.SetValue(current.choices[0])
		}
		fields = append(fields, current)
	}

	model := Model{kind: kind, fields: fields}
	if len(model.fields) > 0 {
		model.fields[0].input.Focus()
	}
	return model
}

func (m *Model) Update(msg tea.Msg) tea.Cmd {
	key, ok := msg.(tea.KeyPressMsg)
	if ok {
		switch key.String() {
		case tui.KeyUp:
			if m.currentIsChoice() {
				m.previousChoice()
				return nil
			}
			m.previousField()
			return nil
		case tui.KeyDown:
			if m.currentIsChoice() {
				m.nextChoice()
				return nil
			}
			m.nextField()
			return nil
		case tui.KeyTab:
			m.nextField()
			return nil
		case tui.KeyShiftTab:
			m.previousField()
			return nil
		}
	}

	if m.index < 0 || m.index >= len(m.fields) {
		return nil
	}
	if m.currentIsChoice() {
		return nil
	}
	var cmd tea.Cmd
	m.fields[m.index].input, cmd = m.fields[m.index].input.Update(msg)
	return cmd
}
