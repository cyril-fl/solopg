package field

import (
	"solopg/internal/app/tui"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

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

	return &selectField[T]{
		field:   newField(template.Label, Select, template.Defaultvalue, template.Validator),
		options: options,
	}
}

// -- Field Implementation -- //
func (f *selectField[T]) Focus() tea.Cmd {
	f.focus = true
	return nil
}

func (f *selectField[T]) Blur() {
	f.focus = false
}

func (f *selectField[T]) Value() any {
	return f.defaultvalue
}

// -- Tea Model Implementation -- //
func (m *selectField[T]) Init() tea.Cmd {
	return nil
}

func (m *selectField[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.options.SetSize(size.Width, size.Height)
	}

	var cmd tea.Cmd
	m.options, cmd = m.options.Update(msg)

	return m, cmd
}

func (m *selectField[T]) View() tea.View {
	title := lipgloss.NewStyle().
		Bold(true).
		Render(m.Label())

	options := m.options.View()

	return tea.NewView(
		lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			"",
			options,
		),
	)
}
