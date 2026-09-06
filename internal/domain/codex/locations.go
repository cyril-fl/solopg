package codex

import (
	"fmt"
	"solopg/internal/domain/card/locations"
	"solopg/internal/infrastructure/t"
	"time"
)

type LocationsTable struct {
	*Table[LocationsEntry]
}
type LocationsEntry struct {
	Timestamp time.Time
	Location  *locations.Location
}

type LocationsEntryTemplate struct {
	Location *locations.Location
}

func NewLocationsTable(entries []LocationsEntry) *LocationsTable {
	return &LocationsTable{
		Table: NewTable(entries),
	}
}

func (l *LocationsTable) ensureTable() {
	ensureEmbeddedTable(&l.Table)
}

func (l *LocationsTable) Summaries() []string {
	l.ensureTable()
	summaries := make([]string, 0, len(l.Entries))
	for _, entry := range l.Entries {
		if entry.Location == nil {
			summaries = append(summaries, t.Localize("codex.unknown_location"))
			continue
		}
		summaries = append(summaries, fmt.Sprintf("%s — %s", entry.Location.Name, entry.Location.Description))
	}
	return summaries
}

func (l *LocationsTable) AddEntry(entry LocationsEntryTemplate) {
	l.ensureTable()
	l.Entries = append(l.Entries, LocationsEntry{
		Timestamp: time.Now().UTC(),
		Location:  entry.Location,
	})
}

func (l *LocationsTable) FindEntryByName(name string) *LocationsEntry {
	l.ensureTable()
	for i, entry := range l.Entries {
		if entry.Location != nil && entry.Location.Name == name {
			return &l.Entries[i]
		}
	}
	return nil
}
