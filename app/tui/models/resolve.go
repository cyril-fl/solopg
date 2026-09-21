package models

import (
	"solopg/app/tui"

	tea "charm.land/bubbletea/v2"
)

type resolveModel struct{}

func Resolve() tea.Model {
	return resolveModel{}
}

func (resolveModel) Init() tea.Cmd {
	return func() tea.Msg {
		return tui.ResolutionMsg{
			Completed: true,
		}
	}
}

func (resolveModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return resolveModel{}, nil
}

func (resolveModel) View() tea.View {
	return tea.NewView("")
}
