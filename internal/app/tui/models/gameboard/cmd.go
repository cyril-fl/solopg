package gameboard

import (
	"solopg/internal/app/tui"
	"solopg/internal/infrastructure/t"

	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"

	tea "charm.land/bubbletea/v2"
)

func saveCmd(save func() error) tea.Cmd {
	return func() tea.Msg {
		if save == nil {
			return tui.SaveMsg{Err: t.NewError(&goi18n.LocalizeConfig{MessageID: "error.save_unconfigured"})}
		}
		return tui.SaveMsg{Err: save()}
	}
}
