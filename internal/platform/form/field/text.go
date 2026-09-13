package field

import (
	"solopg/internal/app/tui"
	"solopg/internal/platform/message"

	"strconv"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type textField[T stringOrInt] struct {
	field[T]
	input textinput.Model
}

type TextTemplate[T stringOrInt] struct {
	Label        string
	Input        textinput.Model
	Defaultvalue T
	Validator    func(T) error
}

func TextField[T stringOrInt](template TextTemplate[T]) *textField[T] {
	input := textinput.New()
	setValue(input, template.Defaultvalue)

	return &textField[T]{
		field: newField(template.Label, Input, template.Defaultvalue, template.Validator),
		input: input,
	}
}

// -- Field Implementation -- //
func (f *textField[T]) Focus() tea.Cmd {
	f.focus = true
	return f.input.Focus()
}

func (f *textField[T]) Blur() {
	f.focus = false
	f.input.Blur()
}

func (f *textField[T]) Value() any {
	return f.input.Value()
}

// -- Tea Model Implementation -- //
func (m *textField[T]) Init() tea.Cmd {
	return textinput.Blink
}

func (m *textField[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case tui.KeyEnter:
			return m, message.SendFormMsg[message.FormNextField]()
		case tui.KeyUp:
			return m, message.SendFormMsg[message.FormPreviousField]()
		case tui.KeyDown:
			return m, message.SendFormMsg[message.FormNextField]()
		default:
			newInput, cmd := m.input.Update(msg)
			m.input = newInput
			return m, cmd
		}
	}

	return m, nil
}

func (m *textField[T]) View() tea.View {
	title := lipgloss.NewStyle().
		Bold(true).
		Render(m.Label())

	input := m.input.View()

	return tea.NewView(
		lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			"",
			input,
		),
	)
}

// -- Helper -- //
type stringOrInt interface {
	~string | ~int
}

func setValue[T stringOrInt](input textinput.Model, value T) {
	switch v := any(value).(type) {
	case string:
		input.SetValue(v)
	case int:
		input.SetValue(strconv.Itoa(v))
	}
}
