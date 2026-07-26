package codex

/*
TODO: ajouter des méthodes pour chaque sous codex
*/
type Codex struct {
	NpcsTable      *NpcsTable
	MonstersTable  *MonstersTable
	LocationsTable *LocationsTable
	ObjectsTable   *ObjectsTable
	ObjectifsTable *ObjectifsTable
}

func New() *Codex {
	return &Codex{
		NpcsTable:      NewNpcsTable([]NpcsEntry{}),
		MonstersTable:  NewMonstersTable([]MonstersEntry{}),
		LocationsTable: NewLocationsTable([]LocationsEntry{}),
		ObjectsTable:   NewObjectsTable([]ObjectsEntry{}),
		ObjectifsTable: NewObjectifsTable([]ObjectifsEntry{}),
	}
}

type Table[E any] struct {
	Entries []E
}

func NewTable[E any](entries []E) *Table[E] {
	return &Table[E]{
		Entries: entries,
	}
}

func (t *Table[E]) AddEntry(entry E) {
	t.Entries = append(t.Entries, entry)
}
