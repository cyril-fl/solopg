package mongo

import (
	"time"

	"github.com/google/uuid"

	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/locations"
)

type Save struct {
	CampaignID uuid.UUID            `bson:"campaignId" json:"campaignId"`
	Character  characters.Character `bson:"character" json:"character"`
	Location   locations.Location   `bson:"location" json:"location"`
	CreatedAt  time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time            `bson:"updatedAt" json:"updatedAt"`
}

// func NewSave(character cards.Character, location cards.Location) *Save {
// 	now := time.Now().UTC()

// 	return &Save{
// 		CampaignID: uuid.New(),
// 		Character:  character,
// 		Location:   location,
// 		CreatedAt:  now,
// 		UpdatedAt:  now,
// 	}
// }

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
