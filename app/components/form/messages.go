package form

import (
	tea "charm.land/bubbletea/v2"
)

// Messages
type NextField struct{}
type PreviousField struct{}
type Validate struct{}
type Scroll struct{}
type Submit struct{}

type msg interface {
	PreviousField |
		NextField |
		Validate |
		Scroll |
		Submit
}

func SendMsg[T msg]() tea.Cmd {
	return func() tea.Msg {
		return T{}
	}
}

// Error
type Error struct {
	err error
}

func SendErrorMsg(err error) tea.Cmd {
	return func() tea.Msg {
		return Error{
			err: err,
		}
	}
}
