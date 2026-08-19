package codex

import (
	"solopg/internal/domain/card/locations"
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
