package field

import "solopg/types/id"

// -- Field -- //
// Field represents a typed form field without exposing its concrete type.
type Field interface {
	ID() id.ID
	Label() string
	Value() any
	Validate() error
}

type field[T any] struct {
	id        id.ID
	kind      kind
	label     string
	defaultvalue     T
	validator func(T) error
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
		id:        id.New(),
		kind:      kind,
		label:     label,
		defaultvalue:     value,
		validator: validator,
	}
}

func (f *field[T]) ID() id.ID {
	return f.id
}

func (f *field[T]) Label() string {
	return f.label
}

func (f *field[T]) Validate() error {
	if f.validator == nil {
		return nil
	}

	return f.validator(f.defaultvalue)
}
