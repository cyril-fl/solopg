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

// TODO Add Ctrl Z ctrl Y pour s'il y a des fail de et miss click

// TODO add un menu rool de the dice
// Peut etre changer en ctrl maj 0 , 1; 2 ...

// TODO il y a deja les orcle, je devrais qjouter une liste de mot clé.
//  TODO gerer un systeme de combat
// Et un systeme pour gerer les degat
