package dicemenu

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/gameplay"
	"solopg/internal/infrastructure/t"
	"solopg/types/size"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

func NewModel(size size.Size) list.Model {
	options := gameplay.ListDices()

	items := make([]list.Item, 0, len(options))
	for _, option := range options {
		items = append(items, tui.NewItem(option.GetName(), "", option))
	}

	model := list.New(items, list.NewDefaultDelegate(), size.Width-4, size.Height)
	tui.ConfigureList(&model)
	return model
}

func NewMenuItem(size size.Size) *DiceMenuItem {
	options := gameplay.ListDices()

	items := make([]list.Item, 0, len(options))
	for _, option := range options {
		items = append(items, tui.NewItem(option.GetName(), "", option))
	}

	model := list.New(items, list.NewDefaultDelegate(), size.Width-4, size.Height)
	tui.ConfigureList(&model)
	return &DiceMenuItem{
		id:     "dice",
		list:   model,
		isOpen: false,
	}
}

type DiceMenuItem struct {
	id     string
	list   list.Model
	isOpen bool
}

func (m *DiceMenuItem) ID() string {
	return m.id
}
func (m *DiceMenuItem) IsOpen() bool {
	return m.isOpen
}
func (m *DiceMenuItem) SetOpen(open bool) {
	m.isOpen = open
}

func (m *DiceMenuItem) GetFooter() []string {
	return []string{t.Localize("roll")}
}

func (m *DiceMenuItem) HandleKeyEnter(msg tea.Msg) error {
	return nil
}
func (m *DiceMenuItem) HandleKeyEsc(msg tea.Msg) error {
	return nil
}
func (m *DiceMenuItem) HandleKeyArrow(msg tea.Msg) error {
	return nil
}
func (m *DiceMenuItem) HandleCtrlN(msg tea.Msg) error {
	return nil
}

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
