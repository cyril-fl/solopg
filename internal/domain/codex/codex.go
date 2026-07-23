package codex

import (
	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/locations"
)

type Codex struct {
	NpcsTable      []*characters.Character
	MonstersTable  []*characters.Character
	LocationsTable []*locations.Location  
	ObjectsTable   []interface{}          
	ObjectifsTable []interface{}          
}
type CodexTemplate struct {
	NpcsTable      []*characters.Character 
	MonstersTable  []*characters.Character 
	LocationsTable []*locations.Location   
	ObjectsTable   []interface{}           
	ObjectifsTable []interface{}           
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
