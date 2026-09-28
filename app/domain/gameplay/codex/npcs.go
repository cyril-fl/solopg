package codex

import (
	"errors"
	"fmt"
	"os"
	"solopg/app/domain/card/attributes/rarity"
	"solopg/app/domain/card/characters"
	"solopg/app/domain/card/characters/classes"
	"solopg/app/domain/card/characters/races"
	"solopg/app/services/i18n"
	"solopg/app/services/process/generatestats"
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
		TableData: newTable(i18n.Localize(tablenpcs), entries),
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

	generator := generatestats.NewGenerator(values)
	generator.GenerateFromValues(generatestats.FromValuesParams{
		Randomness: true,
	})

	if generator.HasError() {
		// i18N -- register
		cwd, err := os.Getwd()
		return i18n.NewError("error.unexpected", map[string]any{
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
		// i18N -- register
		return i18n.NewError("error.invalid:new", map[string]any{
			"Subject": i18n.Localize("npcs"),
			"Error":   err,
		})
	}

	tb.Add(character)

	return nil
}

func (tb *NpcsTable) Summaries() []string {
	if len(tb.Entries) == 0 {
		return []string{i18n.Localize("codex.entry:empty")}
	}

	content := strings.Builder{}
	for _, entry := range tb.Entries {
		if entry.Character == nil {
			content.WriteString(i18n.Localize("codex.codex.unknown:entry"))
			continue
		} else {
			fmt.Fprintf(&content, "%s\n", i18n.Localize("codex.npcs:entry", map[string]any{
				"Name":  i18n.Localize(entry.Character.Name),
				"Race":  i18n.Localize(entry.Character.Race),
				"Class": i18n.Localize(entry.Character.Class),
			}))

			fmt.Fprintf(&content, "%s\n", i18n.Localize("stats.entry", entry.Character.Stats.MappedString()))
		}

		content.WriteString("\n\n")
	}

	return []string{content.String()}
}

// - Helper --//
func (tb *NpcsTable) assertEntry(entry map[string]string) error {
	var err []error
	if entry["name"] == "" {
		// i18N -- register
		err = append(err, i18n.NewError("error.required_property", map[string]any{
			"Subject":  i18n.Localize("npcs"),
			"Property": i18n.Localize("property.name"),
		}))
	}
	if entry["description"] == "" {
		// i18N -- register
		err = append(err, i18n.NewError("error.required_property", map[string]any{
			"Subject":  i18n.Localize("npcs"),
			"Property": i18n.Localize("property.description"),
		}))
	}
	if !races.Assert(entry["race"]) {
		// i18N -- register
		err = append(err, i18n.NewError("error.required_property", map[string]any{
			"Subject":  i18n.Localize("npcs"),
			"Property": i18n.Localize("property.race"),
		}))
	}
	if !classes.Assert(entry["class"]) {
		// i18N -- register
		err = append(err, i18n.NewError("error.required_property", map[string]any{
			"Subject":  i18n.Localize("npcs"),
			"Property": i18n.Localize("property.class"),
		}))
	}

	if len(err) > 0 {
		return tb.formatAssertErrors(err)
	}

	return nil
}

func (tb *NpcsTable) Ensure() {
	ensureEmbeddedTable(i18n.Localize(tablenpcs), &tb.TableData)
}
