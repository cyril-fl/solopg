package codex

import (
	"fmt"
	"time"
)

type ObjectifsTable struct {
	*Table[ObjectifsEntry]
}

func (o *ObjectifsTable) Summaries() []string {
	o.ensureTable()
	summaries := make([]string, 0, len(o.Entries))
	for i := range o.Entries {
		summaries = append(summaries, fmt.Sprintf("Objectif #%d", i+1))
	}
	return summaries
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
