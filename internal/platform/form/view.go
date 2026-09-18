package form

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m *Form) View() tea.View {
	contents := make([]string, 0, len(m.fields))

	for _, f := range m.fields {
		contents = append(contents, f.View().Content)
	}

	err := m.GetError()
	if err != nil {
		contents = append(contents, "\n"+err.Error())
	}

	return tea.NewView(
		lipgloss.JoinVertical(lipgloss.Left, contents...),
	)
}
