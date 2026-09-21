package form

import (
	tea "charm.land/bubbletea/v2"
)

// Messages
type NextField struct{}
type PreviousField struct{}
type Validate struct{}
type Submit struct{}

type msg interface {
	PreviousField |
		NextField |
		Validate |
		Submit
}

func SendMsg[T msg]() tea.Cmd {
	return func() tea.Msg {
		return T{}
	}
}

// Scroll
type Scroll struct {
	Content tea.Msg
}

func SendScrollMsg(message tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return Scroll{Content: message}
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
