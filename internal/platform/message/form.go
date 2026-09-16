package message

import (
	tea "charm.land/bubbletea/v2"
)

type FormNextField struct{}
type FormPreviousField struct{}
type FormSubmit struct{}
type FormPostSubmit struct{}
type FormError struct {
	err error
}

/*
REFACTOR
HIGH Implement a message to send form error
*/
func (e FormError) Error() error {
	return e.err
}

type msg interface {
	FormPreviousField |
		FormNextField |
		FormSubmit |
		FormPostSubmit
}

func SendFormMsg[T msg]() tea.Cmd {
	return func() tea.Msg {
		return T{}
	}
}

func SendErrorMsg(err error) tea.Cmd {
	return func() tea.Msg {
		return FormError{
			err: err,
		}
	}
}
