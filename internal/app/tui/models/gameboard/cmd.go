package gameboard

import (
	"solopg/internal/app/tui"
	"solopg/internal/infrastructure/t"

	tea "charm.land/bubbletea/v2"
)

func saveCmd(save func() error) tea.Cmd {
	return func() tea.Msg {
		if save == nil {
			return tui.SaveMsg{Err: t.NewError("error.save_unconfigured")}
		}
		return tui.SaveMsg{Err: save()}
	}
}

/*
	TODO
	LOW Add Ctrl Z ctrl Y pour s'il y a des fail de et miss click

/*
	TODO
	MEDIUM il y a deja les orcle, je devrais qjouter une liste de mot clé.

/*
	TODO
	HIGH gerer un systeme de combat et un systeme pour gerer les degat

*/
