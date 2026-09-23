package form

import (
	"charm.land/bubbles/v2/viewport"
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

// SCROLL
// DISABLED
func (m *Form) EnsureFocusedFieldVisible(viewport *viewport.Model) {
	/*
	 REFACTOR MEDIUM 
	 Faire en sorte que le forme EMETTE un event avec position ect qui soir recu par un e view port et qui gere le scroll
	 plustot que l'etat actuel ou le form gere le scroll du viewport en le recevant en param 

	 ... connard de GPT
	*/
	fieldsTop, fieldsHight := m.getFocusedFieldPosition(*viewport)
	fieldsBottom := fieldsTop + fieldsHight

	viewportHeight := viewport.Height()
	viewportOffset := viewport.YOffset()

	if fieldsTop < 0 || viewportHeight <= 0 {
		return
	}

	if fieldsTop < viewportOffset {
		viewportOffset = fieldsTop
	}

	if fieldsBottom > viewportOffset+viewportHeight {
		viewportOffset = fieldsTop + fieldsHight - viewportHeight
	}

	viewport.SetYOffset(viewportOffset)
}

// SCROLL
// DISABLED
func (m *Form) getFocusedFieldPosition(viewport viewport.Model) (int, int) {
	top := 0

	for _, field := range m.fields {
		content := lipgloss.NewStyle().Width(viewport.Width()).Render(field.View().Content)
		height := lipgloss.Height(content)

		top += height

		if field.IsFocused() {
			return top, height
		}
	}

	return -1, 0
}
