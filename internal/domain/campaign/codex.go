package campaign

import (
	"solopg/internal/domain/codex"

	"github.com/google/uuid"
)

type Codex struct {
	CampaignID uuid.UUID   `bson:"campaignId" json:"campaignId"`
	Codex      codex.Codex `bson:"codex" json:"codex"`
}
type CodexTemplate struct {
	CampaignID uuid.UUID
	Codex      codex.Codex
}

func NewCodex(params CodexTemplate) *Codex {
	return &Codex{
		CampaignID: params.CampaignID,
		Codex:      params.Codex,
	}
}
