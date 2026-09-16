package form

import (
	"solopg/internal/platform/form/field"
	"solopg/internal/platform/message"

	tea "charm.land/bubbletea/v2"
)

// -- Form -- //
type Model struct {
	fields     []field.Field
	autoSubmit bool
	err        []error
}

func New() *Model {
	return &Model{
		fields:     make([]field.Field, 0),
		autoSubmit: true,
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

func (m *Model) GetValuesAsString() string {
	/*
		TODO Implement a method to get the form values as a string representation
		HIGH
	*/
	return "Form Values:..."
}

func (m *Model) Validate() {
	for _, f := range m.fields {
		if err := f.Validate(); err != nil {
			m.SetError(err)
		}
	}
}

func (m *Model) SetError(err error) {
	m.err = append(m.err, err)
}

func (m *Model) GetErrors() []error {
	return m.err
}

func (m *Model) HasErrors() bool {
	return len(m.err) > 0
}

func (m *Model) EnableAutoSubmit() *Model {
	return m.SetAutoSubmit(true)
}

func (m *Model) DisableAutoSubmit() *Model {
	return m.SetAutoSubmit(false)
}

func (m *Model) SetAutoSubmit(autoSubmit bool) *Model {
	m.autoSubmit = autoSubmit
	return m
}

func (m *Model) Focus() *Model {
	return m.focusOnIndex(0)
}

func (m *Model) focusOnIndex(index int) *Model {
	for _, f := range m.fields {
		f.Blur()
	}
	if index >= 0 && index < len(m.fields) {
		m.fields[index].Focus()
	} else if len(m.fields) > 0 {
		m.fields[0].Focus()
	}

	return m
}

func (m *Model) getCurrentField() field.Field {
	if index := m.getCurrentFieldIndex(); index >= 0 && index < len(m.fields) {
		return m.fields[index]
	}
	return nil
}

func (m *Model) getCurrentFieldIndex() int {
	for i, f := range m.fields {
		if f.IsFocused() {
			return i
		}
	}
	return -1
}

func (m *Model) isLastField() bool {
	currentIndex := m.getCurrentFieldIndex()
	return currentIndex == len(m.fields)-1
}

// -- Tea Model Implementation --//
func (m *Model) Init() tea.Cmd {
	cmds := make([]tea.Cmd, 0, len(m.fields))
	for _, f := range m.fields {
		cmds = append(cmds, f.Init())
	}
	return tea.Batch(cmds...)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	currentIndex := m.getCurrentFieldIndex()
	if currentIndex == -1 {
		return m, nil
	}

	switch msg.(type) {
	case message.FormNextField:
		if m.isLastField() && m.autoSubmit {
			return m, message.SendFormMsg[message.FormSubmit]()
		}
		return m.focusOnIndex(currentIndex + 1), nil
	case message.FormPreviousField:
		return m.focusOnIndex(currentIndex - 1), nil
	}

	return m.updateFocusedField(msg)
}

func (m *Model) updateFocusedField(msg tea.Msg) (tea.Model, tea.Cmd) {
	index := m.getCurrentFieldIndex()
	if index < 0 {
		return m, nil
	}

	newField, cmd := m.fields[index].Update(msg)
	m.fields[index] = newField.(field.Field)

	return m, cmd
}
