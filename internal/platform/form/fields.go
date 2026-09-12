package form

import "solopg/types/id"

type field[T any] struct {
	id        id.ID
	label     string
	value     T
	validator func(T) error
}


type InputField struct {
	field[string]
}

type SelectField struct {
	field[string]
}