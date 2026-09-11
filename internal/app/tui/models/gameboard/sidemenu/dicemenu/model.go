package dicemenu

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/gameplay"
	"solopg/types/size"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

func NewSideMenu(size size.Size, focused bool) *DiceMenu {
	options := gameplay.ListDices()

	items := make([]list.Item, 0, len(options))
	for _, option := range options {
		items = append(items, tui.NewItem(option.GetName(), "", option))
	}

	model := list.New(items, list.NewDefaultDelegate(), size.Width-4, size.Height)
	tui.ConfigureList(&model)
	tui.SetListFocus(&model, focused)

	return &DiceMenu{
		id:      "dice",
		focused: focused,
		list:    model,
	}
}

type DiceMenu struct {
	id      string
	focused bool
	list    list.Model
}

func (m *DiceMenu) ID() string {
	return m.id
}
func (m *DiceMenu) IsOpen() bool {
	return false
}
func (m *DiceMenu) SetOpen(open bool) {
	// No action needed as it doesn't have an open state
}

func (m *DiceMenu) SetFocus(focused bool) {
	m.focused = focused
	tui.SetListFocus(&m.list, focused)
}

func (m *DiceMenu) GetList() list.Model {
	return m.list
}
func (m *DiceMenu) SetList(list list.Model) {
	m.list = list
}

func (m *DiceMenu) HandleKeyEnter(msg tea.Msg) error {
// func handleDiceRoll(m model) (model, tea.Cmd) {
// 	selected, ok := m.diceMenu.SelectedItem().(tui.Item[gameplay.Dice])
// 	if !ok {
// 		return m, nil
// 	}

// 	dice := selected.Value()

// 	result := dice.Roll()

// 	message := fmt.Sprintf("%s : %d", dice.GetName(), result)

// 	m.engine.AddJournalEntry("Dice", message)
// 	m.messages = append(m.messages, message)

// 	m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width()).Render(strings.Join(m.messages, "\n")))
// 	m.viewport.GotoBottom()

// 	return m, nil
// }
	return nil
}
func (m *DiceMenu) HandleKeyEsc(msg tea.Msg) error {
	return nil
}
func (m *DiceMenu) HandleCtrlN(msg tea.Msg) error {
	return nil
}

