package loadsave

import (
	"solopg/app/domain/campaign"
	"solopg/app/tui"
	"solopg/app/tui/models"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type model struct {
	list      list.Model
	selected  *campaign.Campaign
	cancelled bool
}

func NewModel(saves []campaign.Campaign) model {
	return model{list: makeModel(saves)}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		/*
			TODOLOW Standardizer les tea.WindowSizeMsg et tea.KeyMsg ainsi que leur retour
		*/
		m.list.SetSize(msg.Width, max(0, msg.Height-3))
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case tui.KeyEnter:
			if selected, ok := m.list.SelectedItem().(models.Item[*campaign.Campaign]); ok {
				m.selected = selected.Value()
			}
			return m, func() tea.Msg {
				return tui.ResolutionMsg{
					Completed: true,
					Value:     m.selected,
				}
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}
