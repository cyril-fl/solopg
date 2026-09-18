package codex

import (
	"fmt"
	"solopg/internal/domain/card/attributes"
	"solopg/internal/domain/card/attributes/rarity"
	"solopg/internal/domain/card/locations"
	"solopg/internal/infrastructure/t"
	"time"
)

// -- Table --//
type LocationsTable struct {
	*TableData[LocationsEntry]
}
type LocationsEntry struct {
	Timestamp time.Time           `bson:"timestamp"`
	Location  *locations.Location `bson:"location"`
}

var tablelocations = "codex.locations"

func NewLocationsTable(entries []LocationsEntry) *LocationsTable {
	return &LocationsTable{
		TableData: newTable(t.Localize(tablelocations), entries),
	}
}

// -- Methods --//
func (l *LocationsTable) Add(location *locations.Location) {
	l.Entries = append(l.Entries, LocationsEntry{
		Timestamp: time.Now().UTC(),
		Location:  location,
	})
}

func (l *LocationsTable) AddFromMappedValues(values map[string]string) error {
	if err := l.assertEntry(values); err != nil {
		return err
	}

	location, err := locations.New(locations.Template{
		Name:        values["name"],
		Description: values["description"],
		Rarity:      rarity.Default(),
		Variety:     attributes.LocationCard,
	})

	if err != nil {
		return err
	}

	l.Add(location)

	return nil
}

func (l *LocationsTable) Summaries() []string {
	summaries := make([]string, 0, len(l.Entries))
	for _, entry := range l.Entries {
		if entry.Location == nil {
			summaries = append(summaries, t.Localize("codex.unknown_location"))
			continue
		}
		summaries = append(summaries, fmt.Sprintf("%s — %s", entry.Location.Name, entry.Location.Description))
	}

	if len(summaries) == 0 {
		summaries = append(summaries, t.Localize("codex.no_locations"))
	}

	return summaries
}

func (l *LocationsTable) FindEntryByName(name string) *LocationsEntry {

	for i, entry := range l.Entries {
		if entry.Location != nil && entry.Location.Name == name {
			return &l.Entries[i]
		}
	}
	return nil
}

// -- Helper --//

func (l *LocationsTable) assertEntry(values map[string]string) error {
	var err []error

	if values["name"] == "" {
		err = append(err, fmt.Errorf("name is empty"))
	}
	if values["description"] == "" {
		err = append(err, fmt.Errorf("description is empty"))
	}

	if len(err) > 0 {
		return l.formatAssertErrors(err)
	}

	return nil
}

func (l *LocationsTable) ensure() {
	ensureEmbeddedTable(t.Localize(tablelocations), &l.TableData)
}
