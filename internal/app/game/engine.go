package game

import (
	// "fmt"
	// "fmt"
	"solopg/internal/domain/campaign"
	"solopg/internal/domain/card/locations"
	"solopg/internal/domain/codex"
	"solopg/types/id"
)

// DEPRECATED

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

func (e *Engine) Initialize() {
	e.Log("Game engine initialized. Welcome to SoloPG!")
	e.DiscoverLocation(e.State.CurrentLocation)
}

func (e *Engine) UpdateLocation(newLocation *locations.Location) {
	e.State.CurrentLocation = newLocation
	e.DiscoverLocation(newLocation)
}

func (e *Engine) Log(message string) {
	e.State.Journal.AddEntry(message)
}

func (e *Engine) DiscoverLocation(newLocation *locations.Location) {
	if newLocation == nil {
		return
	}

	LocationsEntry := e.State.Codex.LocationsTable.FindEntryByName(newLocation.Name)
	if LocationsEntry != nil {
		e.Log("Player moved to " + newLocation.Name)
		return
	}

	e.State.Codex.LocationsTable.AddEntry(codex.LocationsEntryTemplate{
		Location: newLocation,
	})

	e.Log("New location discovered: " + newLocation.Name)
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
