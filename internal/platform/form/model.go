package form

import (
	"solopg/types/id"
)

// -- Form -- //
type Model struct {
	fields []Field
	err    []string
}

func New() *Model {
	return &Model{
		fields: make([]Field, 0),
		err:    make([]string, 0),
	}
}

func NewForm(fields ...Field) *Model {
	return New().Add(fields...)
}

func (m *Model) Add(fields ...Field) *Model {
	m.fields = append(m.fields, fields...)
	return m
}

func (m *Model) Validate()          {}
func (m *Model) Submit()            {}
func (m *Model) SetError(err error) {}


// func (m *Model) NextField() tea.Cmd { return m.moveFocus(1) }

// func (m *Model) PreviousField() tea.Cmd { return m.moveFocus(-1) }

// func (m *Model) moveFocus(step int) tea.Cmd {
// 	if len(m.fields) == 0 {
// 		return nil
// 	}
// 	// m.fields[m.index].Blur()xzzz
// 	// for range m.fields {
// 	// 	m.index = (m.index + step + len(m.fields)) % len(m.fields)
// 	// 	if m.fields[m.index].Enabled() {
// 	// 		return m.fields[m.index].Focus()
// 	// 	}
// 	// }
// 	return nil
// }


// func (m *Model) Update(msg tea.Msg) tea.Cmd {
// 	return nil
// }


// -- Field -- //
// Field represents a typed form field without exposing its concrete type.
type Field interface {
	ID() id.ID
	Label() string
	Value() any
	Validate() error
}

// NewField creates a form field with an automatically generated identifier.
func NewField[T any](label string, value T, validator func(T) error) Field {
	return &field[T]{
		id:        id.New(),
		label:     label,
		value:     value,
		validator: validator,
	}
}

func (f *field[T]) ID() id.ID {
	return f.id
}

func (f *field[T]) Label() string {
	return f.label
}

func (f *field[T]) Value() any {
	return f.value
}

func (f *field[T]) Validate() error {
	if f.validator == nil {
		return nil
	}

	return f.validator(f.value)
}
