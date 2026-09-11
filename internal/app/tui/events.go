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
	// Control keys
	KeyCtrlC      = "ctrl+c"
	KeyQuit       = "ctrl+q"
	KeySave       = "ctrl+s"
	KeyShiftTab   = "shift+tab"
	KeyCtrlN      = "ctrl+n"
	KeyShiftEnter = "shift+enter"

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

	// Other keys
	EmptyKey = ""
)
