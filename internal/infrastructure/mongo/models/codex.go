package mongo

import (
	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/locations"

	"github.com/google/uuid"
)

type Codex struct {
	CampaignID     uuid.UUID               `bson:"campaignId" json:"campaignId"`
	NpcsTable      []*characters.Character `bson:"characterTable" json:"characterTable"`
	MonstersTable  []*characters.Character `bson:"monsterTable" json:"monsterTable"`
	LocationsTable []*locations.Location   `bson:"locationTable" json:"locationTable"`
	ObjectsTable   []interface{}           `bson:"objectTable" json:"objectTable"`     // Todo implement a proper type for objects
	ObjectifsTable []interface{}           `bson:"objectifTable" json:"objectifTable"` // Todo implement a proper type for objectifs
}
type CodexTemplate struct {
	CampaignID     uuid.UUID               `bson:"campaignId" json:"campaignId"`
	NpcsTable      []*characters.Character `bson:"characterTable" json:"characterTable"`
	MonstersTable  []*characters.Character `bson:"monsterTable" json:"monsterTable"`
	LocationsTable []*locations.Location   `bson:"locationTable" json:"locationTable"`
	ObjectsTable   []interface{}           `bson:"objectTable" json:"objectTable"`     // Todo implement a proper type for objects
	ObjectifsTable []interface{}           `bson:"objectifTable" json:"objectifTable"` // Todo implement a proper type for objectifs
}

func NewCodex(params CodexTemplate) *Codex {
	return &Codex{
		NpcsTable:      params.NpcsTable,
		MonstersTable:  params.MonstersTable,
		LocationsTable: params.LocationsTable,
		ObjectsTable:   params.ObjectsTable,
		ObjectifsTable: params.ObjectifsTable,
	}
}
