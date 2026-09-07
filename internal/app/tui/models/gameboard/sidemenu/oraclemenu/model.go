package oraclemenu

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/gameplay"
	"solopg/internal/infrastructure/t"
	"solopg/types/size"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

func NewModel(size size.Size) list.Model {
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
	return model
}

func NewMenuItem(size size.Size) *OracleMenuItem {
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
	return &OracleMenuItem{
		id:     "oracle",
		list:   model,
		isOpen: false,
	}
}

type OracleMenuItem struct {
	id     string
	list   list.Model
	isOpen bool
}

func (m *OracleMenuItem) ID() string {
	return m.id
}
func (m *OracleMenuItem) IsOpen() bool {
	return m.isOpen
}
func (m *OracleMenuItem) SetOpen(open bool) {
	m.isOpen = open
}

func (m *OracleMenuItem) GetFooter() []string {
	return []string{t.Localize("roll")}
}

func (m *OracleMenuItem) HandleKeyEnter(msg tea.Msg) error {
	return nil
}
func (m *OracleMenuItem) HandleKeyEsc(msg tea.Msg) error {
	return nil
}
func (m *OracleMenuItem) HandleKeyArrow(msg tea.Msg) error {
	return nil
}
func (m *OracleMenuItem) HandleCtrlN(msg tea.Msg) error {
	return nil
}
