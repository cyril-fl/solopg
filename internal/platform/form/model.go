package form

import "solopg/internal/platform/form/field"

// -- Form -- //
type Model struct {
	fields []field.Field
	err    []string
}

func New() *Model {
	return &Model{
		fields: make([]field.Field, 0),
	}
}

func NewForm(fields ...field.Field) *Model {
	return New().Add(fields...)
}

func (m *Model) Add(fields ...field.Field) *Model {
	m.fields = append(m.fields, fields...)
	return m
}

func (m *Model) Submit() {
	m.Validate()
	if len(m.err) > 0 {
		return
	}
}

func (m *Model) Validate() {
	for _, f := range m.fields {
		if err := f.Validate(); err != nil {
			m.SetError(err)
		}
	}
}

func (m *Model) SetError(err error) {
	m.err = append(m.err, err.Error())
}
