package game

import (
	"fmt"
	"solopg/internal/domain/campaign"
	"solopg/internal/platform/jsonlog"

	"github.com/google/uuid"
)

type Engine struct {
	CampaignID uuid.UUID
	State      *State
}

func NewEngine(CampaignID uuid.UUID, State *State) *Engine {
	return &Engine{
		CampaignID: CampaignID,
		State:      State,
	}
}

func Boot(save *campaign.Campaign) *Engine {
	// TODO: Implement logic to load the game state from the selected save or create a new game state if no saves are available.

	if save == nil {
		fmt.Println("No save selected, creating new game.")
		return NewEngine(uuid.Nil, nil)
	}
	jsonlog.JsonifiedLog(save)
	return NewEngine(uuid.Nil, nil)
}
