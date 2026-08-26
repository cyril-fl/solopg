package gameboard

import (
	"fmt"
	"solopg/internal/app/tui"

	tea "charm.land/bubbletea/v2"
)

func saveCmd(save func() error) tea.Cmd {
	return func() tea.Msg {
		if save == nil {
			return tui.SaveMsg{Err: fmt.Errorf("fonction de sauvegarde non configurée")}
		}
		return tui.SaveMsg{Err: save()}
	}
}
