package codexsidemenu

import tea "charm.land/bubbletea/v2"

func (m *Model) SetMenuSize(width, height int) {
	m.menu.SetSize(width, height)
}

func (m *Model) SelectMenu(index int) { m.menu.Select(index) }

func (m *Model) UpdateMenu(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.menu, cmd = m.menu.Update(msg)
	return cmd
}
