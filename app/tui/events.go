package tui

import (
	"errors"
	"solopg/app/services/t"

	tea "charm.land/bubbletea/v2"
)

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

// Message
type ResolutionMsg struct {
	Completed bool
	Value     any
	Err       error
}

// Save
type SaveMsg struct {
	Err error
}

func SendSaveMsg(save func() error) tea.Cmd {
	return func() tea.Msg {
		if save == nil {
			return SaveMsg{Err: t.NewError("error.save_unconfigured")}
		}
		return SaveMsg{Err: save()}
	}
}

// Scroll
type ScrollMsg struct {
	top    int
	height int
}

func SendScrollMsg(top, height int) tea.Cmd {
	return func() tea.Msg {
		return ScrollMsg{
			top:    top,
			height: height,
		}
	}
}
func (sm ScrollMsg) GetTop() int {
	return sm.top
}
func (sm ScrollMsg) GetHeight() int {
	return sm.height
}
func (sm ScrollMsg) GetBottom() int {
	return sm.top + sm.height
}

// Refresh
type Refresh struct {
	Resize bool
}

func SendRefreshMsg(resize bool) tea.Cmd {
	return func() tea.Msg {
		return Refresh{Resize: resize}
	}
}

// Errors

type ErrorMsg struct {
	Err error
}

var ErrCreationCancelled = errors.New("character creation cancelled")
var ErrSelectionCancelled = errors.New("selection cancelled")

func IsCancelled(err error) bool {
	return errors.Is(err, ErrSelectionCancelled) ||
		errors.Is(err, ErrCreationCancelled)
}

func NormalizeError(err error) error {
	if IsCancelled(err) {
		return nil
	}
	return err
}

// keys
const (
	// Navigation keys
	KeyEnter     = "enter"
	KeyEsc       = "esc"
	KeyBackspace = "backspace"
	KeyTab       = "tab"
	KeyShiftTab  = "shift+tab"

	// Arrow keys
	KeyUp     = "up"
	KeyDown   = "down"
	KeyLeft   = "left"
	KeyRight  = "right"
	KeyPgUp   = "pgup"
	KeyPgDown = "pgdown"

	// Shortcut
	ShortcutCtrlC    = "ctrl+c"
	ShortcutCtrlQ    = "ctrl+q"
	ShortcutCtrlN    = "ctrl+n"
	ShortcutShiftTab = "shift+tab"

	// Command keys
	CmdShiftEnter = "shift+enter"
	CmdAltEnter   = "alt+enter"
	CmdCtrlS      = "ctrl+s"

	// Other keys
	EmptyKey = ""
)
