package game

import (
	// "fmt"
	"solopg/internal/domain/campaign"
	"solopg/types/id"
)

type Engine struct {
	CampaignID id.ID
	State      *State
}

func NewEngine(CampaignID id.ID, State *State) *Engine {
	return &Engine{
		CampaignID: CampaignID,
		State:      State,
	}
}

func Boot(campaign *campaign.Campaign, archives *campaign.Archives) *Engine {
	return NewEngine(campaign.ID, NewState(StateTemplate{
		Campaign: campaign,
		Archives: archives,
	}))
}

func (e *Engine) ExportArchives() *campaign.Archives {
	return &campaign.Archives{
		CampaignID: e.CampaignID,
		Codex:      e.State.Codex,
		Journal:    e.State.Journal,
		CreatedAt:  e.State.Metadata.Archive.CreatedAt,
		UpdatedAt:  e.State.Metadata.Archive.UpdatedAt,
	}
}

func (e *Engine) ExportCampaign() *campaign.Campaign {
	return &campaign.Campaign{
		ID:              e.CampaignID,
		Player:          e.State.Player,
		CurrentLocation: e.State.CurrentLocation,
		CreatedAt:       e.State.Metadata.Campaign.CreatedAt,
		UpdatedAt:       e.State.Metadata.Campaign.UpdatedAt,
	}
}
