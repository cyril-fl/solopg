package game

import (
	"solopg/app/cmdrun/domain/campaign"
	"solopg/app/cmdrun/domain/card/locations"
	"solopg/app/cmdrun/types/id"
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

func (e *Engine) Initialize() {
	// TODO i18N -- log system
	// STEP 3 -- Log journal
	e.Log("Game engine initialized. Welcome to SoloPG!")
	e.DiscoverLocation(e.State.CurrentLocation)
}

func (e *Engine) UpdateLocation(newLocation *locations.Location) {
	e.State.CurrentLocation = newLocation
	e.DiscoverLocation(newLocation)
}

/*
REFACTOR NOTE
Ce systeme de log la est le log:campaign / journal d'aventurier.
Doit être étendu

logger dedans ->
- changement de location quand implementer
- systeme day night si implementer
- rencontre avec PNJ / mob
- details de combat :
  - dégat subit
  - resultats

- object urilisé / equipement enfilé / recu /acheter / vendu si implementer ect
- quetes recus / terminé
*/
func (e *Engine) Log(message string) {
	// TODO MEDIUM check

	// STEP 1
	/*
		Utiles la fonction AddEntry du "log" journal
	*/
	e.State.Log.AddEntry("System", message)
}

func (e *Engine) GetAdventureJournal() *campaign.Journal {
	return e.State.AdventureLog
}

func (e *Engine) DiscoverLocation(newLocation *locations.Location) {
	if newLocation == nil {
		return
	}

	LocationsEntry := e.State.Codex.LocationsTable.FindEntryByName(newLocation.Name)
	// i18N + log
	if LocationsEntry != nil {
		// STEP 3 -- Log journal
		e.Log("Player moved to " + newLocation.Name)
		return
	}

	e.State.Codex.LocationsTable.Add(newLocation)

	// i18N + log
	// STEP 3 -- Log journal
	e.Log("New location discovered: " + newLocation.Name)
}

func (e *Engine) ExportArchives() *campaign.Archives {
	return &campaign.Archives{
		CampaignID: e.CampaignID,
		Codex:      e.State.Codex,
		Journal:    e.State.AdventureLog,
		Log:        e.State.Log,
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
