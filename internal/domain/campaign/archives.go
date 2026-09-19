package campaign

import (
	"solopg/internal/domain/gameplay/codex"
	"solopg/types/id"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Archives struct {
	CampaignID id.ID            `bson:"campaignId" json:"campaignId"`
	Codex      *codex.Codex     `bson:"codex" json:"codex"`
	Journal    *Journal `bson:"journal" json:"journal"`
	Log        *Journal `bson:"log" json:"log"`
	CreatedAt  time.Time        `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time        `bson:"updatedAt" json:"updatedAt"`
}

type ArchivesTemplate struct {
	CampaignID id.ID
	Codex      *codex.Codex
	Journal    *Journal
	Log        *Journal
}

func NewArchives(params ArchivesTemplate) *Archives {
	now := time.Now().UTC()

	return &Archives{
		CampaignID: params.CampaignID,
		Codex:      params.Codex,
		Journal:    params.Journal,
		Log:        params.Log,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

func (a *Archives) SetUpdatedAt(t time.Time) {
	a.UpdatedAt = t
}

func (a *Archives) Filter() bson.M {
	return bson.M{
		"campaignId": a.CampaignID,
	}
}
