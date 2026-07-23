package campaign

import (
	"time"

	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/locations"

	"github.com/google/uuid"
)

type Save struct {
	CampaignID      uuid.UUID            `bson:"campaignId" json:"campaignId"`
	Player          characters.Character `bson:"character" json:"character"`
	CurrentLocation locations.Location   `bson:"location" json:"location"`
	CreatedAt       time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt       time.Time            `bson:"updatedAt" json:"updatedAt"`
}

type SaveTemplate struct {
	Player          *characters.Character
	CurrentLocation *locations.Location
}

func NewSave(params SaveTemplate) *Save {
	now := time.Now().UTC()

	return &Save{
		CampaignID:      uuid.New(),
		Player:          *params.Player,
		CurrentLocation: *params.CurrentLocation,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

// func LoadSave(campaignID string, character characters.Character, location locations.Location) (*Save, error) {
// 	parsedCampaignID, err := uuid.Parse(campaignID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Save{
// 		CampaignID: parsedCampaignID,
// 		Character:  character,
// 		Location:   location,
// 	}, nil
// }

// func (save *Save) Touch() {
// 	now := time.Now().UTC()

// 	if save.CreatedAt.IsZero() {
// 		save.CreatedAt = now
// 	}

// 	save.UpdatedAt = now
// }
