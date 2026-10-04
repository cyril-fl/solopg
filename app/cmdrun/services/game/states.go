package game

import (
	"solopg/app/cmdrun/domain/campaign"
	"solopg/app/cmdrun/domain/card/characters"
	"solopg/app/cmdrun/domain/card/locations"
	"solopg/app/cmdrun/domain/gameplay/codex"
	"time"
)

type State struct {
	Player          *characters.Character
	CurrentLocation *locations.Location
	Codex           *codex.Codex
	/*
		AdventureLog         *campaign.AdventureLog
		OK, j'ai:
			- un Avdenturer Journl AJ =
			- un Log journal LJ = Log

		AJ me semble etre la chat romm

		c'est bien ca , c'est les log de debut de partie .

		log n'est jamais utiliser
	*/
	AdventureLog *campaign.Journal
	Log          *campaign.Journal
	Metadata     Metadata
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
		TODO LOW Add a timer logic to register the time spent in the game and update the state accordingly.
	*/
	return &State{
		Player:          data.Campaign.Player,
		CurrentLocation: data.Campaign.CurrentLocation,
		Codex:           data.Archives.Codex,
		AdventureLog:    data.Archives.Journal,
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
