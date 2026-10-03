package loadsave

import (
	"solopg/app/cmdrun/domain/campaign"
	cmdruntui "solopg/app/cmdrun/tui"
	"solopg/app/cmdrun/tui/models"
	sharedtui "solopg/app/shared/tui"

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
			TODO LOW Standardizer les tea.WindowSizeMsg et tea.KeyMsg ainsi que leur retour
		*/
		m.list.SetSize(msg.Width, max(0, msg.Height-3))
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case sharedtui.KEY_ENTER:
			if selected, ok := m.list.SelectedItem().(models.Item[*campaign.Campaign]); ok {
				m.selected = selected.Value()
			}
			return m, func() tea.Msg {
				return cmdruntui.ResolutionMsg{
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
