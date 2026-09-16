package game

import (
	"solopg/internal/domain/campaign"
	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/locations"
	"solopg/internal/domain/codex"
	"solopg/internal/domain/journal"
	"time"
)

type State struct {
	Player          *characters.Character
	CurrentLocation *locations.Location
	Codex           *codex.Codex
	Journal         *journal.Journal
	Log             *journal.Journal
	Metadata        Metadata
}
type Metadata struct {
	Archive  Timestamps
	Campaign Timestamps
}
type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
type CampaignData struct {
	Campaign *campaign.Campaign
	Archives *campaign.Archives
}

func NewState(data CampaignData) *State {
	/*
		TODO
		LOW Add a timer logic to register the time spent in the game and update the state accordingly.
	*/
	return &State{
		Player:          data.Campaign.Player,
		CurrentLocation: data.Campaign.CurrentLocation,
		Codex:           data.Archives.Codex,
		Journal:         data.Archives.Journal,
		Log:             data.Archives.Log,
		Metadata: Metadata{
			Archive: Timestamps{
				CreatedAt: data.Archives.CreatedAt,
				UpdatedAt: data.Archives.UpdatedAt,
			},
			Campaign: Timestamps{
				CreatedAt: data.Campaign.CreatedAt,
				UpdatedAt: data.Campaign.UpdatedAt,
			},
		},
	}
}
