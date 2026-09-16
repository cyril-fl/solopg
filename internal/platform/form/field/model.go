package field

import (
	"errors"
	"solopg/internal/infrastructure/t"

	tea "charm.land/bubbletea/v2"
)

// -- Field -- //
// Field represents a typed form field without exposing its concrete type.
type Field interface {
	ID() string
	Label() string
	Focus() tea.Cmd
	Blur()
	IsFocused() bool
	Value() any
	GetValueAsString() string
	Reset()
	Validate() error
	SetError(err error)
	GetError() error
	ResetError()
	Init() tea.Cmd
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() tea.View
}

type field[T any] struct {
	id           string
	kind         kind
	label        string
	focus        bool
	defaultvalue T
	validator    func(T) error
	required     bool
	err          []error
}

type kind string

var (
	// FieldTypeInput represents a text input field.
	Input kind = "input"
	// FieldTypeSelect represents a select field.
	Select kind = "select"
)

// newField creates a form field with an automatically generated identifier.
func newField[T any](id, label string, kind kind, value T, validator func(T) error, required bool) field[T] {
	return field[T]{
		id:           id,
		kind:         kind,
		label:        label,
		defaultvalue: value,
		validator:    validator,
		required:     required,
	}
}

// -- Methods --//
func (f *field[T]) ID() string {
	return f.id
}

func (f *field[T]) Label() string {
	return t.Localize(f.label)
}

func (f *field[T]) SetError(err error) {
	f.err = append(f.err, err)
}

func (f *field[T]) GetError() error {
	if len(f.err) > 0 {
		return errors.Join(f.err...)
	}

	return nil
}

func (f *field[T]) ResetError() {
	f.err = nil
}

// -- Helper --//
func (f *field[T]) validate(value T) error {
	if f.validator != nil {
		err := f.validator(value)

		f.SetError(err)
		return err
	}

	return nil
}

func (f *field[T]) IsFocused() bool {
	return f.focus
}
