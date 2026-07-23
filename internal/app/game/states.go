package game

import (
	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/locations"
	"solopg/internal/domain/codex"
)

type State struct {
	Player          *characters.Character
	CurrentLocation *locations.Location
	Codex           *codex.Codex
}

type StateTemplate struct {
	Player          *characters.Character
	CurrentLocation *locations.Location
	Codex           *codex.Codex
}

func NewState(params StateTemplate) *State {
	// TODO: Add a timer logic to register the time spent in the game and update the state accordingly.
	return &State{
		Player:          params.Player,
		CurrentLocation: params.CurrentLocation,
		Codex:           params.Codex,
	}
}
