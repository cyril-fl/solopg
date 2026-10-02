package campaign

import (
	"solopg/app/cmdrun/domain/gameplay/codex"
	"solopg/app/cmdrun/types/id"
	"solopg/app/cmdrun/types/interfaces"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Archives struct {
	CampaignID id.ID        `bson:"campaignId" json:"campaignId"`
	Codex      *codex.Codex `bson:"codex" json:"codex"`
	Journal    *Journal     `bson:"journal" json:"journal"`
	Log        *Journal     `bson:"log" json:"log"`
	CreatedAt  time.Time    `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time    `bson:"updatedAt" json:"updatedAt"`
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

func (a *Archives) EnsureInitialized() {
	if a.Codex == nil {
		a.Codex = codex.New()
	}

	if a.Journal == nil {
		a.Journal = NewJournal([]Entry{})
	}

	if a.Log == nil {
		a.Log = NewJournal([]Entry{})
	}

	list := []interfaces.Initializable{
		a.Codex,
		a.Journal,
		a.Log,
	}

	for _, archive := range list {
		archive.EnsureInitialized()
	}
}
