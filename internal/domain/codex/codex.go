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

// EnsureInitialized repairs a Codex loaded from storage where nested tables may be nil.
func EnsureInitialized(c *Codex) *Codex {
	if c == nil {
		return New()
	}

	if c.NpcsTable == nil {
		c.NpcsTable = NewNpcsTable([]NpcsEntry{})
	} else {
		c.NpcsTable.ensureTable()
	}

	if c.MonstersTable == nil {
		c.MonstersTable = NewMonstersTable([]MonstersEntry{})
	} else {
		c.MonstersTable.ensureTable()
	}

	if c.LocationsTable == nil {
		c.LocationsTable = NewLocationsTable([]LocationsEntry{})
	} else {
		c.LocationsTable.ensureTable()
	}

	if c.ObjectsTable == nil {
		c.ObjectsTable = NewObjectsTable([]ObjectsEntry{})
	} else {
		c.ObjectsTable.ensureTable()
	}

	if c.ObjectifsTable == nil {
		c.ObjectifsTable = NewObjectifsTable([]ObjectifsEntry{})
	} else {
		c.ObjectifsTable.ensureTable()
	}

	return c
}

func ensureEmbeddedTable[E any](table **Table[E]) {
	if *table == nil {
		*table = NewTable([]E{})
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
