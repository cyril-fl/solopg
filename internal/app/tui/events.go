package tui

import (
	"errors"
)

// Message
type ResolutionMsg struct {
	Completed bool
	Value     any
	Err       error
}

type SaveMsg struct {
	Err error
}

type ErrorMsg struct {
	Err error
}

// Errors
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

	// Arrow keys
	KeyUp    = "up"
	KeyDown  = "down"
	KeyLeft  = "left"
	KeyRight = "right"

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
