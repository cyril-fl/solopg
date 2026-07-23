package bootui

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/campaign"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type model struct {
	list list.Model

	selected *campaign.Campaign

	cancelled bool
}

func newModel(saves []campaign.Campaign) model {
	saveItems := []list.Item{
		saveItem{title: "New game"},
	}

	for _, save := range saves {
		saveItems = append(saveItems, saveItem{
			title:       save.Title(),
			description: save.Description(),
			save:        &save,
		})
	}

	savesList := list.New(saveItems, list.NewDefaultDelegate(), 0, 0)
	configureList(&savesList)

	return model{list: savesList}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, max(0, msg.Height-1))
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case tui.KeyCtrlC, tui.KeyQuit:
			m.cancelled = true
			return m, tea.Quit

		case tui.KeyEnter:
			if selected, ok := m.list.SelectedItem().(saveItem); ok {
				m.selected = selected.save
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}
