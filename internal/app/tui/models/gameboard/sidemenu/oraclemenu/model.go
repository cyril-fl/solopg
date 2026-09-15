package oraclemenu

import (
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/models/gameboard/sidemenu"
	"solopg/internal/domain/gameplay"
	"solopg/internal/infrastructure/t"
	"solopg/types/size"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

// -- OracleMenu -- //
// Menu
func NewSideMenu(size size.Size, focused bool) *OracleMenu {
	items := make([]list.Item, 0)
	for _, oracle := range gameplay.GetOracle() {
		if oracle.Visible {
			key := "oracle." + oracle.ID
			name := t.Localize(key)
			items = append(items, tui.NewItem(name, "", oracle))
		}
	}

	model := list.New(items, list.NewDefaultDelegate(), size.Width-4, size.Height)
	tui.ConfigureList(&model)
	tui.SetListFocus(&model, focused)

	return &OracleMenu{
		id:      "oracles",
		focused: focused,
		list:    model,
	}
}

type OracleMenu struct {
	id      string
	focused bool
	list    list.Model
}

// / Getters and Setters
func (m *OracleMenu) ID() string {
	return m.id
}
func (m *OracleMenu) IsOpen() bool {
	return false
}
func (m *OracleMenu) SetOpen(open bool) {
	// No action needed as it doesn't have an open state
}

func (m *OracleMenu) SetFocus(focused bool) {
	m.focused = focused
	tui.SetListFocus(&m.list, focused)
}

func (m *OracleMenu) GetList() *list.Model {
	return &m.list
}
func (m *OracleMenu) SetList(list list.Model) {
	m.list = list
}

// / Handlers
func (m *OracleMenu) handleKeyShiftEnter() tea.Cmd {
	selected, ok := m.list.SelectedItem().(tui.Item[*gameplay.Oracle])
	if !ok {
		return func() tea.Msg {
			return tui.ErrorMsg{Err: t.NewError("error.oracle_selection")}
		}
	}

	oracle := selected.Value()
	rollResult, err := gameplay.RollOracle[any](oracle)
	if err != nil {
		return func() tea.Msg {
			return tui.ErrorMsg{Err: err}
		}
	}

	return func() tea.Msg {
		return Msg{
			Result: rollResult,
		}
	}
}

func (m *OracleMenu) HandleUpdate(params sidemenu.UpdateParams) (tea.Model, tea.Cmd) {
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
	Result *gameplay.OracleResult[any]
}
