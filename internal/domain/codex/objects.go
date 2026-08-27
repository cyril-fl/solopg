package codex

import (
	"fmt"
	"solopg/internal/domain/card/objects"
	"time"
)

type ObjectsTable struct {
	*Table[ObjectsEntry]
}

func (o *ObjectsTable) Summaries() []string {
	o.ensureTable()
	summaries := make([]string, 0, len(o.Entries))
	for _, entry := range o.Entries {
		if entry.object == nil {
			summaries = append(summaries, "Objet inconnu")
			continue
		}
		summaries = append(summaries, fmt.Sprintf("%s — %s", entry.object.Name, entry.object.Description))
	}
	return summaries
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

func (o *ObjectsTable) AddObject(object *objects.Object) {
	o.AddEntry(ObjectsEntryTemplate{object: object})
}
