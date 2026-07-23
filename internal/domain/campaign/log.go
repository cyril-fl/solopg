package campaign

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID         uuid.UUID `bson:"id" json:"id"`
	CampaignID uuid.UUID `bson:"campaignId" json:"campaignId"`
	Message    string    `bson:"message" json:"message"`
	Timestamp  time.Time `bson:"timestamp" json:"timestamp"`
}

type LogTemplate struct {
	CampaignID uuid.UUID
	Message    string
}

func NewLog(params LogTemplate) *Log {
	return &Log{
		ID:         uuid.New(),
		CampaignID: params.CampaignID,
		Message:    params.Message,
		Timestamp:  time.Now().UTC(),
	}
}
