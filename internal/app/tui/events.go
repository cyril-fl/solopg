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
	KeyCtrlC    = "ctrl+c"
	KeyEnter    = "enter"
	KeyQuit     = "ctrl+q"
	KeySettings = "s"
	KeyBack     = "b"
	KeyEsc      = "esc"
	KeyEmpty    = ""
)
