package codex

import (
	"errors"
	"fmt"
)

// - Codex - //
type Codex struct {
	NpcsTable      *NpcsTable
	BeastsTable    *BeastsTable
	LocationsTable *LocationsTable
	ObjectsTable   *ObjectsTable
	ObjectifsTable *ObjectivesTable
}

func New() *Codex {
	return &Codex{
		NpcsTable:      NewNpcsTable([]NpcsEntry{}),
		BeastsTable:    NewBeastsTable([]BeastEntry{}),
		LocationsTable: NewLocationsTable([]LocationsEntry{}),
		ObjectsTable:   NewObjectsTable([]ObjectEntry{}),
		ObjectifsTable: NewObjectivesTable([]ObjectifEntry{}),
	}
}

// EnsureInitialized repairs a Codex loaded from storage where nested tables may be nil.
func (c *Codex) EnsureInitialized() *Codex {
	if c.NpcsTable == nil {
		c.NpcsTable = NewNpcsTable(nil)
	}
	if c.BeastsTable == nil {
		c.BeastsTable = NewBeastsTable(nil)
	}
	if c.LocationsTable == nil {
		c.LocationsTable = NewLocationsTable(nil)
	}
	if c.ObjectsTable == nil {
		c.ObjectsTable = NewObjectsTable(nil)
	}
	if c.ObjectifsTable == nil {
		c.ObjectifsTable = NewObjectivesTable(nil)
	}

	c.NpcsTable.Ensure()
	c.BeastsTable.Ensure()
	c.LocationsTable.Ensure()
	c.ObjectsTable.Ensure()
	c.ObjectifsTable.Ensure()

	return c
}

// - Codex table - //
type Table interface {
	Summaries() []string
	AddFromMappedValues(values map[string]string) error
	Ensure()
}

type TableData[E any] struct {
	name    string
	Entries []E `bson:"entries"`
}

func newTable[E any](name string, entries []E) *TableData[E] {
	return &TableData[E]{
		name:    name,
		Entries: entries,
	}
}

func ensureEmbeddedTable[E any](name string, table **TableData[E]) {
	if *table == nil {
		*table = newTable(name, []E{})
	}
}

func (t TableData[any]) formatAssertErrors(err []error) error {
	return errors.Join(
		fmt.Errorf("invalid %s:", t.name),
		errors.Join(err...),
	)
}
