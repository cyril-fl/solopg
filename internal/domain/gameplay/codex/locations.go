package codex

import (
	"fmt"
	"solopg/internal/domain/card/attributes/rarity"
	"solopg/internal/domain/card/attributes/variety"
	"solopg/internal/domain/card/locations"
	"solopg/internal/infrastructure/t"
	"time"
)

// - Table --//
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

// - Methods --//
func (tb *LocationsTable) Add(location *locations.Location) {
	tb.Entries = append(tb.Entries, LocationsEntry{
		Timestamp: time.Now().UTC(),
		Location:  location,
	})
}

func (tb *LocationsTable) AddFromMappedValues(values map[string]string) error {
	if err := tb.assertEntry(values); err != nil {
		return err
	}

	location, err := locations.New(locations.Template{
		Name:        values["name"],
		Description: values["description"],
		Rarity:      rarity.Default(),
		Variety:     variety.MakeDefault("location_card"),
	})

	if err != nil {
		return err
	}

	tb.Add(location)

	return nil
}

func (tb *LocationsTable) Summaries() []string {
	summaries := make([]string, 0, len(tb.Entries))
	for _, entry := range tb.Entries {
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

func (tb *LocationsTable) FindEntryByName(name string) *LocationsEntry {

	for i, entry := range tb.Entries {
		if entry.Location != nil && entry.Location.Name == name {
			return &tb.Entries[i]
		}
	}
	return nil
}

// - Helper --//

func (tb *LocationsTable) assertEntry(values map[string]string) error {
	var err []error

	if values["name"] == "" {
		err = append(err, fmt.Errorf("name is empty"))
	}
	if values["description"] == "" {
		err = append(err, fmt.Errorf("description is empty"))
	}

	if len(err) > 0 {
		return tb.formatAssertErrors(err)
	}

	return nil
}

func (tb *LocationsTable) Ensure() {
	if tb.name == "" {
		tb.name = t.Localize(tablelocations)
	}
	if tb.Entries == nil {
		tb.Entries = []LocationsEntry{}
	}
}
