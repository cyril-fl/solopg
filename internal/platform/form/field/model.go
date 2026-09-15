package field

import (
	"solopg/internal/infrastructure/t"
	"solopg/types/id"

	tea "charm.land/bubbletea/v2"
)

// -- Field -- //
// Field represents a typed form field without exposing its concrete type.
type Field interface {
	ID() id.ID
	Label() string
	Focus() tea.Cmd
	Blur()
	IsFocused() bool
	Value() any
	Validate() error
	Init() tea.Cmd
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() tea.View
}

type field[T any] struct {
	id           id.ID
	kind         kind
	label        string
	focus        bool
	defaultvalue T
	validator    func(T) error
	err		  error
}

type kind string

var (
	// FieldTypeInput represents a text input field.
	Input kind = "input"
	// FieldTypeSelect represents a select field.
	Select kind = "select"
)

// newField creates a form field with an automatically generated identifier.
func newField[T any](label string, kind kind, value T, validator func(T) error) field[T] {
	return field[T]{
		id:           id.New(),
		kind:         kind,
		label:        label,
		defaultvalue: value,
		validator:    validator,
	}
}

func (f *field[T]) ID() id.ID {
	return f.id
}

func (f *field[T]) Label() string {
	return t.Localize(f.label)
}

func (f *field[T]) Validate() error {
	if f.validator == nil {
		return nil
	}

	f.err = f.validator(f.defaultvalue)
	
	return f.err
}

func (f *field[T]) IsFocused() bool {
	return f.focus
}
