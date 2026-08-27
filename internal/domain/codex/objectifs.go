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
		if o.Entries[i].Title == "" {
			summaries = append(summaries, fmt.Sprintf("Objectif #%d", i+1))
			continue
		}
		summaries = append(summaries, fmt.Sprintf("%s — %s", o.Entries[i].Title, o.Entries[i].Description))
	}
	return summaries
}

type ObjectifsEntry struct {
	Timestamp   time.Time
	Title       string
	Description string
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
		Timestamp: time.Now().UTC(),
	})
}

func (o *ObjectifsTable) AddObjectif(title, description string) {
	o.ensureTable()
	o.Entries = append(o.Entries, ObjectifsEntry{
		Timestamp:   time.Now().UTC(),
		Title:       title,
		Description: description,
	})
}
