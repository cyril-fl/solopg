package cmdservetui

import (
	"solopg/app/shared/services/factory/millfilterlog"
	"solopg/app/shared/services/logs"

	tea "charm.land/bubbletea/v2"
)

func (m *model[T]) View() tea.View {
	mill := millfilterlog.New(millfilterlog.Template{
		LogList: logs.ConvertStringablesToList(m.events),
		Flags:   m.flags,
	})

	mill.Filter()

	if m.HasErr() {
		m.handleStreamError(m.GetErr())
	}

	return tea.NewView(mill.GetJoinedStringList())
}

func (m *model[T]) makeLogList() []T {
	list := make([]T, 0, len(m.events))
	for _, event := range m.events {
		list = append(list, event)
	}
	return list
}
