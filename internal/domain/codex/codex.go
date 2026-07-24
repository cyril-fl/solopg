package codex

import (
	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/locations"
)

/*
	TODO: Ameliorer le codex avec des vrai entrée pour chaque sous codex
	s'insperer de Journal
	ajouter des méthodes
*/
type Codex struct {
	NpcsTable      []*characters.Character
	MonstersTable  []*characters.Character
	LocationsTable []*locations.Location
	ObjectsTable   []interface{}
	ObjectifsTable []interface{}
}
type Template struct {
	NpcsTable      []*characters.Character
	MonstersTable  []*characters.Character
	LocationsTable []*locations.Location
	ObjectsTable   []interface{}
	ObjectifsTable []interface{}
}

func New(params Template) *Codex {
	return &Codex{
		NpcsTable:      params.NpcsTable,
		MonstersTable:  params.MonstersTable,
		LocationsTable: params.LocationsTable,
		ObjectsTable:   params.ObjectsTable,
		ObjectifsTable: params.ObjectifsTable,
	}
}
