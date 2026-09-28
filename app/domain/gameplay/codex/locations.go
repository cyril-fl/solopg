package codex

import (
	"fmt"
	"solopg/app/domain/card/attributes/rarity"
	"solopg/app/domain/card/attributes/variety"
	"solopg/app/domain/card/locations"
	"solopg/app/services/i18n"
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
		TableData: newTable(i18n.Localize(tablelocations), entries),
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
		return []string{i18n.Localize("codex.entry:empty")}
	}

	content := strings.Builder{}
	for _, entry := range tb.Entries {
		if entry.Location == nil {
			content.WriteString(i18n.Localize("codex.codex.unknown:entry"))
		} else {
			fmt.Fprintf(&content, "%s — %s", i18n.Localize(entry.Location.GetName()), i18n.Localize(entry.Location.GetDescription()))
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
		// i18N -- register
		err = append(err, i18n.NewError("error.required_property", map[string]any{
			"Subject":  i18n.Localize("location"),
			"Property": i18n.Localize("name"),
		}))
	}
	if values["description"] == "" {
		// i18N -- register
		err = append(err, i18n.NewError("error.required_property", map[string]any{
			"Subject":  i18n.Localize("location"),
			"Property": i18n.Localize("description"),
		}))
	}

	if len(err) > 0 {
		return tb.formatAssertErrors(err)
	}

	return nil
}

func (tb *LocationsTable) Ensure() {
	if tb.name == "" {
		tb.name = i18n.Localize(tablelocations)
	}
	if tb.Entries == nil {
		tb.Entries = []LocationsEntry{}
	}
}
