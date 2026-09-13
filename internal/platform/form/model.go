package form

import (
	"solopg/internal/platform/form/field"
	"solopg/internal/platform/message"

	tea "charm.land/bubbletea/v2"
)

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

func (m *Model) Focus() *Model {
	return m.FocusOnIndex(0)
}

func (m *Model) FocusOnIndex(index int) *Model {
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

func (m *Model) GetFields() []field.Field {
	return m.fields
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
			return m.FocusOnIndex(currentIndex + 1), nil
		case message.FormPreviousField:
			return m.FocusOnIndex(currentIndex - 1), nil
		case message.FormSubmit:
			// fmt.Println("SubmitFormMsg received in codexMenu.HandleUpdate")
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
