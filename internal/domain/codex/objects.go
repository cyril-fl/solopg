package codex

import (
	"solopg/internal/domain/card/objects"
	"time"
)

type ObjectsTable struct {
	*Table[ObjectsEntry]
}

type ObjectsEntry struct {
	timestamp time.Time
	object    *objects.Object
}

type ObjectsEntryTemplate struct {
	object *objects.Object
}

func NewObjectsTable(entries []ObjectsEntry) *ObjectsTable {
	return &ObjectsTable{
		Table: NewTable(entries),
	}
}

func (o *ObjectsTable) ensureTable() {
	ensureEmbeddedTable(&o.Table)
}

func (o *ObjectsTable) AddEntry(entry ObjectsEntryTemplate) {
	o.ensureTable()
	o.Entries = append(o.Entries, ObjectsEntry{
		timestamp: time.Now().UTC(),
		object:    entry.object,
	})
}
