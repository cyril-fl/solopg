package codex

import (
	"solopg/internal/domain/card/attributes"
	"solopg/internal/domain/card/locations"
	"solopg/internal/domain/card/objects"
	"strings"
)

/*
TODO: ajouter des méthodes pour chaque sous codex
*/
// -- Codex -- //
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
func (c *Codex) EnsureInitialized() *Codex {
	if c.NpcsTable == nil {
		c.NpcsTable = NewNpcsTable(nil)
	}
	if c.MonstersTable == nil {
		c.MonstersTable = NewMonstersTable(nil)
	}
	if c.LocationsTable == nil {
		c.LocationsTable = NewLocationsTable(nil)
	}
	if c.ObjectsTable == nil {
		c.ObjectsTable = NewObjectsTable(nil)
	}
	if c.ObjectifsTable == nil {
		c.ObjectifsTable = NewObjectifsTable(nil)
	}

	c.NpcsTable.ensureTable()
	c.MonstersTable.ensureTable()
	c.LocationsTable.ensureTable()
	c.ObjectsTable.ensureTable()
	c.ObjectifsTable.ensureTable()

	return c
}

func ensureEmbeddedTable[E any](table **Table[E]) {
	if *table == nil {
		*table = NewTable([]E{})
	}
}

// TODO ajouter un reel assert
func (codexData *Codex) AddLocation(values map[string]string) error {
	location, err := locations.New(locations.Template{
		Name:        values["name"],
		Description: values["description"],
		Rarity:      attributes.F,
		Variety:     attributes.LocationCard,
	})
	if err != nil {
		return err
	}
	codexData.LocationsTable.AddEntry(LocationsEntryTemplate{Location: location})
	return nil
}

func (codexData *Codex) AddObject(values map[string]string) error {
	category := objects.Category(strings.ToLower(values["category"]))
	if err := category.Validate(); err != nil {
		return err
	}
	object, err := objects.New(objects.Template{
		Name:        values["name"],
		Description: values["description"],
		Rarity:      attributes.F,
		Variety:     attributes.ArticleCard,
		Category:    category,
	})
	if err != nil {
		return err
	}
	codexData.ObjectsTable.AddObject(object)
	return nil
}

func (codexData *Codex) AddObjective(values map[string]string) error {
	codexData.ObjectifsTable.AddObjectif(values["title"], values["description"])
	return nil
}

// -- Codex table -- //
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
