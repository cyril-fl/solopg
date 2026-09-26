package codex

import (
	"errors"
	"fmt"
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/domain/card/characters/races"
	"solopg/app/services/i18n"
	"solopg/app/services/process/generatestats"
	"strings"
	"time"
)

// - Table --//
type BeastiaryTable struct {
	*TableData[BeastiaryEntry]
}

type BeastiaryEntry struct {
	Timestamp time.Time   `bson:"timestamp"`
	Race      *races.Race `bson:"race"`
}

var tablebeast = "codex.beastiary"

func NewBeastiaryTable(entries []BeastiaryEntry) *BeastiaryTable {
	entries = superChargeEntries(entries)

	return &BeastiaryTable{
		TableData: newTable(i18n.Localize(tablebeast), entries),
	}
}

// - Methods --//
func (tb *BeastiaryTable) Add(race *races.Race) {
	tb.Entries = append(tb.Entries, BeastiaryEntry{
		Timestamp: time.Now().UTC(),
		Race:      race,
	})
}

func (tb *BeastiaryTable) HasEntry(race races.Race) bool {
	for _, entry := range tb.Entries {
		if entry.Race.GetHash() == race.GetHash() {
			return true
		}
	}

	return false
}

func (tb *BeastiaryTable) AddFromMappedValues(values map[string]string) error {
	if err := tb.assertEntry(values); err != nil {
		return err
	}

	generator := generatestats.NewGenerator(values)
	generator.GenerateFromValues(generatestats.FromValuesParams{
		Randomness: false,
	})

	if generator.HasError() {
		return fmt.Errorf("failed to generate stats from values: %w", generator.GetError())
	}

	/*
		NOTE: Generated bonuses may be negative.
		First of all, this is intentional and follows the following logic:
		`a Troll is stupid and therefore receives an Intelligence penalty.``

		Then, this should not be handled by the process itself, but by the template
		selected with values["encounter"].

		These are configurable in the /data/systems/rules/oracles/encounter directory.
	*/
	raw := races.Template{
		Name:        values["name"],
		Description: values["description"],
		Playable:    false,
		Bonus:       generator.GetModifiers(),
	}

	race := races.New(raw)
	if tb.HasEntry(race) {
		return fmt.Errorf("race %s already exists in the beastiary", race.GetName())
	}

	if err := races.AddInConfig(raw); err != nil {
		return fmt.Errorf("failed to save race: %w", err)
	}

	tb.Add(&race)

	return nil
}

/*
TODO LOW Modifier ca que ce soit un peu joli....
Creer un genre de "page dans codex avec une methode de rendu
description
viewport (mettre la liste et prevoir une fonction de la redu pour une enum ou un truc du genre)
*/
func (tb *BeastiaryTable) Summaries() []string {
	if len(tb.Entries) == 0 {
		return []string{i18n.Localize("codex.entry:empty")}
	}

	content := strings.Builder{}
	for _, entry := range tb.Entries {
		if entry.Race == nil {
			content.WriteString(i18n.Localize("codex.codex.unknown:entry"))
		} else {
			fmt.Fprintf(&content, "%s\n", i18n.Localize("codex.beastiary:entry", map[string]any{
				"Name":        i18n.Localize(entry.Race.GetName()),
				"Description": i18n.Localize(entry.Race.GetDescription()),
			}))

			if len(entry.Race.GetBonus()) > 0 {
				modifiersToStats := stats.New(entry.Race.GetBonus())
				fmt.Fprintf(&content, "%s\n", i18n.Localize("codex.entry:stats", modifiersToStats.MappedString()))
			}
		}
		content.WriteString("\n\n")
	}

	return []string{content.String()}
}

// - Helper --//
func (tb *BeastiaryTable) assertEntry(entry map[string]string) error {
	var err []error

	if entry["name"] == "" {
		err = append(err, fmt.Errorf("missing name for beast"))
	}

	if entry["description"] == "" {
		err = append(err, fmt.Errorf("missing description for beast"))
	}

	if len(err) > 0 {
		return tb.formatAssertErrors(err)
	}

	return nil
}

func (tb *BeastiaryTable) Ensure() {
	ensureEmbeddedTable(i18n.Localize(tablebeast), &tb.TableData)
	tb.syncWithConfig()
}

func (tb *BeastiaryTable) syncWithConfig() error {
	var errs []error

	for _, race := range tb.Entries {
		if err := races.AddInConfig(races.Template{
			Name:        race.Race.GetName(),
			Description: race.Race.GetDescription(),
			Playable:    race.Race.IsPlayable(),
			Bonus:       race.Race.GetBonus(),
		}); err != nil {
			errs = append(errs,
				fmt.Errorf(
					"failed to save race %s: %w",
					race.Race.GetName(),
					err,
				),
			)
		}
	}

	return errors.Join(errs...)
}

func superChargeEntries(entries []BeastiaryEntry) []BeastiaryEntry {
	preloaded := getPreloadBeastiaryEntries()

	return append(preloaded, entries...)
}

func getPreloadBeastiaryEntries() []BeastiaryEntry {
	m := make(map[string]BeastiaryEntry)

	for _, entry := range races.List() {
		m[entry.Name] = BeastiaryEntry{
			Timestamp: time.Now().UTC(),
			Race:      &entry,
		}
	}

	slice := make([]BeastiaryEntry, 0, len(m))
	for value := range m {
		slice = append(slice, m[value])
	}

	return slice
}
