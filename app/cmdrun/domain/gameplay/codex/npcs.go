package codex

import (
	"errors"
	"fmt"
	"os"
	"solopg/app/cmdrun/domain/card/attributes/rarity"
	"solopg/app/cmdrun/domain/card/characters"
	"solopg/app/cmdrun/domain/card/characters/classes"
	"solopg/app/cmdrun/domain/card/characters/races"
	"solopg/app/cmdrun/services/factory/millstats"
	"solopg/app/shared/services/i19n"
	"solopg/app/shared/services/logs"
	"strings"

	"time"
)

// - Table --//
type NpcsTable struct {
	*TableData[NpcsEntry]
}
type NpcsEntry struct {
	Timestamp time.Time             `bson:"timestamp"`
	Character *characters.Character `bson:"character"`
}

var tablenpcs = "codex.npcs"

func NewNpcsTable(entries []NpcsEntry) *NpcsTable {
	return &NpcsTable{
		TableData: newTable(i19n.Localize(tablenpcs), entries),
	}
}

// - Methods --//
func (tb *NpcsTable) Add(character *characters.Character) {
	tb.Entries = append(tb.Entries, NpcsEntry{
		Timestamp: time.Now().UTC(),
		Character: character,
	})
}

func (tb *NpcsTable) AddFromMappedValues(values map[string]string) error {
	if err := tb.assertEntry(values); err != nil {
		return err
	}

	generator := millstats.NewWithValues(values)
	generator.GenerateFromValues(millstats.FromValuesParams{
		Randomness: true,
	})

	if generator.HasError() {

		cwd, err := os.Getwd()
		return logs.NewError("error.unexpected", map[string]any{
			"Path":  cwd,
			"Error": errors.Join(err, generator.GetError()),
		})
	}

	/*
		NOTE Description should be a text field that describes the character in a RP way, and includes details
		about the character's equipment, abilities, and other relevant information. Power level should
		be calculated / represented by a dice roll in an oracle, for example:
		a powerful character stats would be chosen with the legendary.
	*/
	character, err := characters.New(characters.Template{
		Name:        values["name"],
		Description: values["description"],
		Rarity:      rarity.Default(),
		Race:        values["race"],
		Class:       values["class"],
		Stats:       generator.GetStats(),
	})

	if err != nil {

		return logs.NewError("error.invalid:new", map[string]any{
			"Subject": "npcs",
			"Error":   err,
		})
	}

	tb.Add(character)

	return nil
}

func (tb *NpcsTable) Summaries() []string {
	if len(tb.Entries) == 0 {
		return []string{i19n.Localize("codex.entry:empty")}
	}

	content := strings.Builder{}
	for _, entry := range tb.Entries {
		if entry.Character == nil {
			content.WriteString(i19n.Localize("codex.codex.unknown:entry"))
			continue
		} else {
			fmt.Fprintf(&content, "%s\n", i19n.Localize("codex.npcs:entry", map[string]any{
				"Name":  i19n.Localize(entry.Character.Name),
				"Race":  i19n.Localize(entry.Character.Race),
				"Class": i19n.Localize(entry.Character.Class),
			}))

			fmt.Fprintf(&content, "%s\n", i19n.Localize("stats.entry", entry.Character.Stats.MappedString()))
		}

		content.WriteString("\n\n")
	}

	return []string{content.String()}
}

// - Helper --//
func (tb *NpcsTable) assertEntry(entry map[string]string) error {
	var err []error
	if entry["name"] == "" {

		err = append(err, logs.NewError("error.required", map[string]any{
			"Subject":  "npcs",
			"Property": "property.name",
		}))
	}
	if entry["description"] == "" {

		err = append(err, logs.NewError("error.required", map[string]any{
			"Subject":  "npcs",
			"Property": "property.description",
		}))
	}
	if !races.Assert(entry["race"]) {

		err = append(err, logs.NewError("error.required", map[string]any{
			"Subject":  "npcs",
			"Property": "property.race",
		}))
	}
	if !classes.Assert(entry["class"]) {

		err = append(err, logs.NewError("error.required", map[string]any{
			"Subject":  "npcs",
			"Property": "property.class",
		}))
	}

	if len(err) > 0 {
		return tb.formatAssertErrors(err)
	}

	return nil
}

func (tb *NpcsTable) Ensure() {
	ensureEmbeddedTable(i19n.Localize(tablenpcs), &tb.TableData)
}
