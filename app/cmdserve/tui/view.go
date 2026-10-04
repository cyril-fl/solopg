package cmdservetui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m *model[T]) View() tea.View {
	var view strings.Builder

	for _, event := range m.events {
		fmt.Fprintln(&view, event.String())
	}

	if m.err != nil {
		// i18N
		fmt.Fprintf(&view, "\nErreur: %v\n", m.err)
	}

	return tea.NewView(view.String())
}
