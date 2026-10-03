package codex

import (
	"fmt"
	"solopg/app/cmdrun/domain/card/attributes/rarity"
	"solopg/app/cmdrun/domain/card/attributes/variety"
	"solopg/app/cmdrun/domain/card/locations"
	"solopg/app/shared/services/i19n"
	"solopg/app/shared/services/logs"
	"strings"
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
		TableData: newTable(i19n.Localize(tablelocations), entries),
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
	if len(tb.Entries) == 0 {
		return []string{i19n.Localize("codex.entry:empty")}
	}

	content := strings.Builder{}
	for _, entry := range tb.Entries {
		if entry.Location == nil {
			content.WriteString(i19n.Localize("codex.codex.unknown:entry"))
		} else {
			fmt.Fprintf(&content, "%s — %s", i19n.Localize(entry.Location.GetName()), i19n.Localize(entry.Location.GetDescription()))
		}
		content.WriteString("\n\n")
	}

	return []string{content.String()}
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

		err = append(err, logs.NewError("error.required", map[string]any{
			"Subject":  "location",
			"Property": "name",
		}))
	}
	if values["description"] == "" {

		err = append(err, logs.NewError("error.required", map[string]any{
			"Subject":  "location",
			"Property": "description",
		}))
	}

	if len(err) > 0 {
		return tb.formatAssertErrors(err)
	}

	return nil
}

func (tb *LocationsTable) Ensure() {
	if tb.name == "" {
		tb.name = i19n.Localize(tablelocations)
	}
	if tb.Entries == nil {
		tb.Entries = []LocationsEntry{}
	}
}
