package gameboard

import (
	"solopg/app/services/t"
	"solopg/app/tui"

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

// TODO HIGH il y a deja les orcle, je devrais qjouter une liste de mot clé.
// BACKLOG gerer un systeme de combat et un systeme pour gerer les degat

/*
BACKLOG Add Ctrl Z ctrl Y pour s'il y a des fail de et miss click
- Faire une queue Passed event
- Faire une stack future event
le tout = History {
	Passed
	Future
	Curent
}

faire une evnt undo / redo en en fonction de tui.CtrlZ / CtrlY

si on entre un nouvelle element utilser une methode dumpFuture()

gere les element en tyant les string
- Player entry
- System entry

seul le Player entry sont undoable / redoable, si on tomber sur un System entry on stop la navigation
(eviter de refaire de jeux de des, d'optimser ect)

le ctrl Z / Y ne concerne que la vue, pas les codex ect .

*/
