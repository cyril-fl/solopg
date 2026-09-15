package dicemenu

import (
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/models/gameboard/sidemenu"
	"solopg/internal/domain/gameplay"
	"solopg/internal/infrastructure/t"
	"solopg/types/size"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

// -- DiceMenu -- //
// Menu
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

// / Getters and Setters
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

func (m *DiceMenu) GetList() *list.Model {
	return &m.list
}
func (m *DiceMenu) SetList(list list.Model) {
	m.list = list
}

// / Handlers
func (m *DiceMenu) handleKeyShiftEnter() tea.Cmd {
	selected, ok := m.list.SelectedItem().(tui.Item[gameplay.Dice])
	if !ok {
		return func() tea.Msg {
			return tui.ErrorMsg{Err: t.NewError("error.dice_selection")}
		}
	}

	dice := selected.Value()
	return func() tea.Msg {
		return Msg{
			Dice:  dice.GetName(),
			Value: dice.Roll(),
		}
	}
}

func (m *DiceMenu) HandleUpdate(params sidemenu.UpdateParams) (tea.Model, tea.Cmd) {
	switch msg := params.Msg.(type) {
				case tea.KeyPressMsg:
		switch msg.String() {
			case tui.CmdShiftEnter, tui.CmdAltEnter:
				return params.Model, m.handleKeyShiftEnter()
		default:
			return params.Delegate(msg)
		}
	default:
		return params.Delegate(msg)
	}
	return params.Model, nil
}

// Messages
type Msg struct {
	Dice  string
	Value int
}
