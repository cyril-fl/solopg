package campaign

import (
	"solopg/internal/domain/codex"
	"solopg/internal/domain/journal"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Archives struct {
	CampaignID uuid.UUID        `bson:"campaignId" json:"campaignId"`
	Codex      *codex.Codex     `bson:"codex" json:"codex"`
	Journal    *journal.Journal `bson:"journal" json:"journal"`
	CreatedAt  time.Time        `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time        `bson:"updatedAt" json:"updatedAt"`
}

type ArchivesTemplate struct {
	CampaignID uuid.UUID
	Codex      *codex.Codex
	Journal    *journal.Journal
}

func NewArchives(params ArchivesTemplate) *Archives {
	now := time.Now().UTC()

	return &Archives{
		CampaignID: params.CampaignID,
		Codex:      params.Codex,
		Journal:    params.Journal,
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
