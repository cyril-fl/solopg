package field

import (
	"solopg/internal/app/tui"
	"solopg/internal/platform/message"
	"solopg/types/direction"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// -- Field -- //
type selectField[T any] struct {
	field[T]
	options list.Model
}

type SelectTemplate[T any] struct {
	Label        string
	Options      []tui.Item[T]
	Defaultvalue T
	Validator    func(T) error
}

func SelectField[T any](template SelectTemplate[T]) *selectField[T] {
	items := make([]list.Item, 0, len(template.Options))
	for _, option := range template.Options {
		items = append(items, option)
	}

	// FIXME: mettre des taille c'est juste un fix tempraire, ca le le rend pas "flex"
	options := list.New(items, list.NewDefaultDelegate(), 15, 10)
	tui.ConfigureList(&options)
	tui.SetListFocus(&options, false)

	return &selectField[T]{
		field:   newField(template.Label, Select, template.Defaultvalue, template.Validator),
		options: options,
	}
}

// -- Field Implementation -- //
func (f *selectField[T]) Focus() tea.Cmd {
	f.focus = true
	tui.SetListFocus(&f.options,true)
	return nil
}

func (f *selectField[T]) Blur() {
	f.focus = false
	tui.SetListFocus(&f.options, false)
}

func (f *selectField[T]) Value() any {
	return f.defaultvalue
}

// -- Tea Model Implementation -- //
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
					return m, message.SendFormMsg[message.FormNextField]()
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
	if m.err != nil {
		err = lipgloss.NewStyle().
			Foreground(lipgloss.Color("1")).
			Render(m.err.Error())
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

// -- Helper -- //
func (m *selectField[T]) handleDirection(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	updatedList, newdirection := direction.GetListDirection(&m.options, msg)
	m.options = updatedList
	switch newdirection {
		case direction.Next:
			return m, message.SendFormMsg[message.FormNextField]()
		case direction.Previous:
			return m, message.SendFormMsg[message.FormPreviousField]()
	}
	return m, nil
}