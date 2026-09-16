package form

import (
	"errors"
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

// -- Methods --//
// Fields
func (m *Model) Add(fields ...field.Field) *Model {
	m.fields = append(m.fields, fields...)
	return m
}

func (m *Model) Focus() *Model {
	return m.focusOnIndex(0)
}

func (m *Model) Reset() *Model {
	m.resetErrors()

	for _, f := range m.fields {
		f.Reset()
		f.Blur()
	}

	return m.focusOnIndex(0)
}

// Input
func (m *Model) GetValuesAsMappedString() map[string]string {
	formValues := make(map[string]string)

	for _, f := range m.fields {
		formValues[f.ID()] = f.GetValueAsString()
	}

	return formValues
}

// Validation
func (m *Model) Validate() *Model {
	m.resetErrors()

	hasErrors := false
	for _, f := range m.fields {
		if err := f.Validate(); err != nil {
			f.SetError(err)
			hasErrors = true
		}
	}

	if hasErrors {
		m.SetError(errors.New("form validation failed"))
	}
	
	return m
}

func (m *Model) SetError(err error) {
	m.err = append(m.err, err)
}

func (m *Model) GetError() error {
	if len(m.err) > 0 {
		return errors.Join(m.err...)
	}

	return nil
}

func (m *Model) HasErrors() bool {
	return len(m.err) > 0
}

// Submit
func (m *Model) Submit() *Model {
	m.resetErrors()
	m.Validate()

	return m
}

func (m *Model) EnableAutoSubmit() *Model {
	return m.setAutoSubmit(true)
}

func (m *Model) DisableAutoSubmit() *Model {
	return m.setAutoSubmit(false)
}

func (m *Model) setAutoSubmit(autoSubmit bool) *Model {
	m.autoSubmit = autoSubmit
	return m
}

// -- Helper --//
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

func (m *Model) resetErrors() {
	m.err = nil
	for _, f := range m.fields {
		f.ResetError()
	}
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
