package codex

import (
	"errors"
	"fmt"
	"solopg/internal/domain/card/attributes/stats"
	"solopg/internal/domain/card/characters/races"
	"solopg/internal/domain/gameplay/oracle"
	"solopg/internal/infrastructure/t"
	"strconv"
	"strings"
	"time"
)

// - Table --//
type BeastiaryTable struct {
	*TableData[BeastiaryEntry]
}

type BeastiaryEntry struct {
	Timestamp time.Time  `bson:"timestamp"`
	Race      races.Race `bson:"race"`
}

var tablebeast = "codex.beastiary"

func NewBeastiaryTable(entries []BeastiaryEntry) *BeastiaryTable {
	entries = superChargeEntries(entries)

	return &BeastiaryTable{
		TableData: newTable(t.Localize(tablebeast), entries),
	}
}

// - Methods --//
func (tb *BeastiaryTable) Add(race races.Race) {
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

	raw := races.Template{
		Name:        values["name"],
		Description: values["description"],
		Playable:    false,
		Bonus:       tb.makeBonusFromValues(values),
	}

	race := races.New(raw)
	if tb.HasEntry(race) {
		return fmt.Errorf("race %s already exists in the beastiary", race.GetName())
	}

	if err := races.AddInConfig(raw); err != nil {
		return fmt.Errorf("failed to save race: %w", err)
	}

	tb.Add(race)

	return nil
}

/*
TODO HIGH Modifier ca que ce soit un peu joli....
Creer un genre de "page dans codex avec une methode de rendu
description
viewport (mettre la liste et prevoir une fonction de la redu pour une enum ou un truc du genre)
*/
func (tb *BeastiaryTable) Summaries() []string {
	content := strings.Builder{}

	if len(tb.Entries) == 0 {
		content.WriteString(t.Localize("codex.no_beasts"))
		return []string{content.String()}
	}

	for _, entry := range tb.Entries {
		content.WriteString(entry.Race.GetName())
		content.WriteString("\n")
		content.WriteString(entry.Race.GetDescription())
		content.WriteString("\n")
		content.WriteString(stats.New(entry.Race.GetBonus()).String())
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
	ensureEmbeddedTable(t.Localize(tablebeast), &tb.TableData)
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
			Race:      entry,
		}
	}

	slice := make([]BeastiaryEntry, 0, len(m))
	for value := range m {
		slice = append(slice, m[value])
	}

	return slice
}

func (tb BeastiaryTable) makeBonusFromValues(values map[string]string) []stats.Modifier {
	modifier := []stats.Modifier{}

	oracle, err := oracle.GetByID(values["encounter"])
	if err != nil {
		fmt.Printf("Error getting oracle: %v\n", err)
		return modifier
	}

	for _, stat := range stats.List() {
		modifier = append(modifier, tb.handleStatValues(statValuesTemplate{
			stat:     stat,
			value:    values[stat.String()],
			fallback: *oracle,
		}))
	}

	return modifier
}

type statValuesTemplate struct {
	stat     stats.Stat
	value    string
	fallback oracle.Oracle
}

func (tb BeastiaryTable) handleStatValues(params statValuesTemplate) stats.Modifier {
	if bonus, err := strconv.Atoi(params.value); err == nil && params.value != "" {
		return stats.Modifier{
			Stat:  params.stat,
			Value: bonus,
		}
	}

	roll, err := oracle.Roll[int](params.fallback)
	if err != nil {
		fmt.Printf("Error rolling for stat %s: %v\n", params.stat.String(), err)
		return stats.Modifier{
			Stat:  params.stat,
			Value: 0,
		}
	}

	return stats.Modifier{
		Stat:  params.stat,
		Value: roll.Result,
	}
}
