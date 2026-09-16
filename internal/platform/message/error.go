package message

// import tea "charm.land/bubbletea/v2"

type ErrorMsg interface {
	SetError(err error)
	GetError() error
}

// func SendErrorMsg[E ErrorMsg](err error) tea.Cmd {
// 	return func() tea.Msg {
// 		newMsg := new(E)
// 		newMsg.SetError(err)
// 		return *newMsg
// 	}
// }

// func SendErrorMsg[E ErrorMsg](err error) tea.Cmd {
// 	return func() tea.Msg {
// 		var msg E
// 		msg.SetError(err)
// 		return msg
// 	}
// }