package codex

import (
	"fmt"
	"solopg/internal/infrastructure/t"
	"time"
)

type ObjectifsTable struct {
	*table[ObjectifsEntry]
}

func (o *ObjectifsTable) Summaries() []string {
	o.ensureTable()
	summaries := make([]string, 0, len(o.Entries))
	for i := range o.Entries {
		if o.Entries[i].Title == "" {
			summaries = append(summaries, t.Localize("objective.number", map[string]any{"Number": i + 1}))
			continue
		}
		summaries = append(summaries, fmt.Sprintf("%s — %s", o.Entries[i].Title, o.Entries[i].Description))
	}

	if len(summaries) == 0 {
		summaries = append(summaries, t.Localize("codex.no_objectives"))
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
		table: NewTable(entries),
	}
}

func (o *ObjectifsTable) ensureTable() {
	ensureEmbeddedTable(&o.table)
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
