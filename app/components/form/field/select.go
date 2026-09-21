package field

import (
	"fmt"
	"solopg/app/components/form"
	"solopg/app/tui"
	"solopg/app/tui/models"
	"solopg/types/direction"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// - Field - //
type selectField[T any] struct {
	field[T]
	options list.Model
}

type SelectTemplate[T any] struct {
	ID           string
	Label        string
	Options      []models.Item[T]
	Defaultvalue T
	Validator    func(T) error
	Required     bool
}

func SelectField[T any](template SelectTemplate[T]) *selectField[T] {
	items := make([]list.Item, 0, len(template.Options))
	for _, option := range template.Options {
		items = append(items, option)
	}
	// FIXME LOW Mettre des taille c'est juste un fix tempraire, ca le le rend pas "flex"
	options := list.New(items, list.NewDefaultDelegate(), 15, 10)
	models.ConfigureList(&options)
	models.SetListFocus(&options, false)

	return &selectField[T]{
		field:   newField(template.ID, template.Label, Select, template.Defaultvalue, template.Validator, template.Required),
		options: options,
	}
}

// - Field Implementation - //
func (f *selectField[T]) Focus() tea.Cmd {
	f.focus = true
	models.SetListFocus(&f.options, true)
	return nil
}

func (f *selectField[T]) Blur() {
	f.focus = false
	models.SetListFocus(&f.options, false)
}

func (f *selectField[T]) Value() any {
	return f.options.SelectedItem().(models.Item[T]).Value()
}

func (f *selectField[T]) GetValueAsString() string {
	if f.Value() == nil {
		return ""
	}
	return fmt.Sprintf("%v", f.Value())
}

func (f *selectField[T]) Reset() {
	f.ResetError()
	f.Blur()
	f.options.Select(0)
}

func (f *selectField[T]) Validate() error {
	if err := f.testRequireness(); err != nil {
		return err
	}

	if v, ok := f.Value().(T); ok {
		return f.validate(v)
	}

	return fmt.Errorf("invalid value type for field %s: expected %T, got %T", f.ID(), f.defaultvalue, f.Value())
}

func (f *selectField[T]) testRequireness() error {
	if f.required && f.Value() == nil {
		return fmt.Errorf("field %s is required", f.ID())
	}
	return nil
}

// - Tea Model Implementation - //
func (m *selectField[T]) Init() tea.Cmd {
	return nil
}

func (m *selectField[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.options.SetSize(msg.Width, msg.Height)
	case tea.KeyPressMsg:
		switch msg.String() {
		case tui.KeyUp, tui.KeyDown:
			return m.handleDirection(msg)
		case tui.KeyEnter:
			return m, form.SendMsg[form.NextField]()
		}
	default:
		newOptions, cmd := m.options.Update(msg)
		m.options = newOptions
		return m, cmd
	}
	return m, nil
}

func (m *selectField[T]) View() tea.View {
	title := lipgloss.NewStyle().
		Bold(true).
		Render(m.Label())

	options := m.options.View()

	err := ""
	if m.GetError() != nil {
		err = lipgloss.NewStyle().
			Foreground(lipgloss.Color("1")).
			Render(m.GetError().Error())
	}

	return tea.NewView(
		lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			err,
			options,
		),
	)
}

// - Helper - //
func (m *selectField[T]) handleDirection(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	updatedList, newdirection := direction.GetListDirection(&m.options, msg)
	m.options = updatedList
	switch newdirection {
	case direction.Next:
		return m, form.SendMsg[form.NextField]()
	case direction.Previous:
		return m, form.SendMsg[form.PreviousField]()
	}
	return m, nil
}
