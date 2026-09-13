package message

import tea "charm.land/bubbletea/v2"

type FormNextField struct{}
type FormPreviousField struct{}
type FormSubmit struct{}

type msg interface {
	FormPreviousField |
		FormNextField |
		FormSubmit
}

func SendFormMsg[T msg]() tea.Cmd {
	return func() tea.Msg {
		return T{}
	}
}
