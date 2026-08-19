package choosearchetype

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/characters/archetypes"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type step int

const (
	choiceStep step = iota
	stepConfirm
)

type model struct {
	Step step
	list list.Model
}

func NewModel[T archetypes.Archetype](data []T) model {
	items := make([]list.Item, 0, len(data))

	for _, i := range data {
		if !i.IsPlayable() {
			continue
		}

		items = append(items, tui.NewItem(i.GetName(), "", i))
	}

	listModel := list.New(items, list.NewDefaultDelegate(), 0, 0)
	tui.ConfigureList(&listModel)

	return model{
		Step: choiceStep,
		list: listModel,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		// The parent adds a three-line footer to the child view.
		m.list.SetSize(msg.Width, max(0, msg.Height-3))
		return m, nil
	}

	switch m.Step {
	case choiceStep:
		return m.updateList(msg)
	case stepConfirm:
		return m.updateConfirm(msg)
	default:
		return m, nil
	}
}

func (m model) updateList(msg tea.Msg) (tea.Model, tea.Cmd) {
	updatedList, cmd := m.list.Update(msg)
	m.list = updatedList

	if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == tui.KeyEnter {
		selectedItem := m.list.SelectedItem()
		if selectedItem == nil {
			return m, nil
		}

		return m, func() tea.Msg {
			return tui.ResolutionMsg{
				Completed: true,
				Value:     selectedItem.FilterValue(),
			}
		}
	}

	return m, cmd
}

func (m model) updateConfirm(msg tea.Msg) (tea.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == tui.KeyEnter {
		return m, tea.Quit
	}

	return m, nil
}
