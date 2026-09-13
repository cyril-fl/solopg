package field

import (
	"solopg/internal/infrastructure/t"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type inputField[T any] struct {
	field[T]
	input textinput.Model
}

type InputTemplate[T any] struct {
	Label     string
	Input     textinput.Model
	Defaultvalue     T
	Validator func(T) error
}

func InputField[T any](template InputTemplate[T]) *inputField[T] {
	input := textinput.New()
	// input.Focus()

	return &inputField[T]{
		field: newField(template.Label, Input, template.Defaultvalue, template.Validator),
		input: input,
	}
}


// -- Field Implementation -- //
func (f *inputField[T]) Value() any {
	return f.defaultvalue
}

// -- Tea Model Implementation -- //
func (m *inputField[T]) Init() tea.Cmd {
	return textinput.Blink
}

func (m *inputField[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	newInput, cmd := m.input.Update(msg)
	m.input = newInput

	if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "enter" {
		m.defaultvalue = any(newInput.Value()).(T)
	}

	return m, cmd
}	

func (m *inputField[T]) View() tea.View {
	title := lipgloss.NewStyle().
		Bold(true).
		Render(t.Localize("create_character"))

	input := m.input.View()

	return tea.NewView(
		lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			"",
			t.Localize("name"),
			input,
		),
	)
}