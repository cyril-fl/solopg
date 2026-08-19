package codex

import "time"

type ObjectifsTable struct {
	*Table[ObjectifsEntry]
}

type ObjectifsEntry struct {
	timestamp time.Time
}

type ObjectifsEntryTemplate struct {
}

func NewObjectifsTable(entries []ObjectifsEntry) *ObjectifsTable {
	return &ObjectifsTable{
		Table: NewTable(entries),
	}
}

func (o *ObjectifsTable) ensureTable() {
	ensureEmbeddedTable(&o.Table)
}

func (o *ObjectifsTable) AddEntry(entry ObjectifsEntryTemplate) {
	o.ensureTable()
	o.Entries = append(o.Entries, ObjectifsEntry{
		timestamp: time.Now().UTC(),
	})
}
