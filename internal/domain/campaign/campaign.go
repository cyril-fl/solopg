package campaign

import (
	"fmt"
	"time"

	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/locations"

	"github.com/google/uuid"
)

type Campaign struct {
	ID              uuid.UUID            `bson:"Id" json:"Id"`
	Player          characters.Character `bson:"character" json:"character"`
	CurrentLocation locations.Location   `bson:"location" json:"location"`
	CreatedAt       time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt       time.Time            `bson:"updatedAt" json:"updatedAt"`
}

type Template struct {
	Player          *characters.Character
	CurrentLocation *locations.Location
}

func New(params Template) *Campaign {
	now := time.Now().UTC()

	return &Campaign{
		ID:              uuid.New(),
		Player:          *params.Player,
		CurrentLocation: *params.CurrentLocation,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

func (c *Campaign) Title() string {
	return fmt.Sprintf("%s | %s | %s",
		c.Player.Name,
		c.Player.Race,
		c.Player.Class,
	)
}

func (c *Campaign) Description() string {
	return fmt.Sprintf(
		"Location: %s | Updated: %s",
		c.CurrentLocation.Name,
		c.UpdatedAt.Format("2006-01-02 15:04:05"),
	)
}
