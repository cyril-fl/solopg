package campaign

import (
	"fmt"
	"time"

	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/locations"
	"solopg/internal/infrastructure/t"
	"solopg/types/id"

	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"go.mongodb.org/mongo-driver/bson"
)

type Campaign struct {
	ID              id.ID                 `bson:"Id" json:"Id"`
	Player          *characters.Character `bson:"character" json:"character"`
	CurrentLocation *locations.Location   `bson:"location" json:"location"`
	CreatedAt       time.Time             `bson:"createdAt" json:"createdAt"`
	UpdatedAt       time.Time             `bson:"updatedAt" json:"updatedAt"`
}

type Template struct {
	Player          *characters.Character
	CurrentLocation *locations.Location
}

func New(params Template) *Campaign {
	now := time.Now().UTC()

	return &Campaign{
		ID:              id.New(),
		Player:          params.Player,
		CurrentLocation: params.CurrentLocation,
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
	return t.Localizer.MustLocalize(&goi18n.LocalizeConfig{
		MessageID: "campaign.description",
		TemplateData: map[string]any{
			"Location": c.CurrentLocation.Name,
			"Updated":  c.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func (c *Campaign) SetUpdatedAt(t time.Time) {
	c.UpdatedAt = t
}

func (c *Campaign) Filter() bson.M {
	return bson.M{
		"Id": c.ID,
	}
}
