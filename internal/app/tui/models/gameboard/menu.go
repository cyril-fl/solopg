package gameboard

import tea "charm.land/bubbletea/v2"

type MenuItem interface {
	ID() string
	IsOpen() bool
	SetOpen(bool)

	GetView() string
	GetFooter() []string

	HandleKeyEnter(msg tea.Msg) error
	HandleKeyEsc(msg tea.Msg) error
	HandleKeyArrow(msg tea.Msg) error
	HandleCtrlN(msg tea.Msg) error
}
