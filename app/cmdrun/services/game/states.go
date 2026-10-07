package game

import (
	"solopg/app/cmdrun/domain/campaign"
	"solopg/app/cmdrun/domain/card/characters"
	"solopg/app/cmdrun/domain/card/locations"
	"solopg/app/cmdrun/domain/gameplay/codex"
	"solopg/app/cmdrun/services/timer"
	"time"
)

type State struct {
	Player          *characters.Character
	CurrentLocation *locations.Location
	Codex           *codex.Codex
	AdventureLog    *campaign.Journal
	TimeHistory     []timer.Record
	Timer           *timer.Timer
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
	return &State{
		Player:          data.Campaign.Player,
		CurrentLocation: data.Campaign.CurrentLocation,
		Codex:           data.Archives.Codex,
		AdventureLog:    data.Archives.Journal,
		TimeHistory:     data.Campaign.TimeHistory,
		Timer:           timer.New(),
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
