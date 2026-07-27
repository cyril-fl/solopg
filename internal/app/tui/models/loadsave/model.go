package loadsave

import (
	"fmt"
	"solopg/internal/app/tui"
	"solopg/internal/domain/campaign"

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
			fmt.Println("LOADSAVE SIZE", msg.Width, msg.Height)
		m.list.SetSize(msg.Width, max(0, msg.Height-3))
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case tui.KeyEnter:
			if selected, ok := m.list.SelectedItem().(tui.Item[*campaign.Campaign]); ok {
				m.selected = selected.Value()
			}
			return m, func() tea.Msg {
				return tui.ResolutionMsg{
					Completed: true,
					Value:     m.selected,
					Err:       nil,
				}
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}
