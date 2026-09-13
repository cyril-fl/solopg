package field

type selectField[T any] struct {
	field[T]
}

type SelectTemplate[T any] struct {
	Label     string
	Options   []T
	Defaultvalue     T
	Validator func(T) error
}

func SelectField[T any](template SelectTemplate[T]) *selectField[T] {
	return &selectField[T]{
		field: newField(template.Label, Select, template.Defaultvalue, template.Validator),
	}
}

func (f *selectField[T]) Value() any {
	return f.defaultvalue
}
