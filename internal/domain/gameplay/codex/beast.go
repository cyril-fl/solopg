package codex

import (
	"fmt"
	"solopg/internal/domain/card/attributes/rarity"
	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/characters/classes"
	"solopg/internal/domain/card/characters/races"
	"solopg/internal/infrastructure/t"
	"time"
)

// - Table --//
type BeastsTable struct {
	*TableData[BeastEntry]
}

type BeastEntry struct {
	Timestamp time.Time             `bson:"timestamp"`
	Character *characters.Character `bson:"character"`
}

var tablebeast = "codex.beasts"

func NewBeastsTable(entries []BeastEntry) *BeastsTable {
	if entries == nil {
		entries = []BeastEntry{}
	}

	return &BeastsTable{
		TableData: newTable(t.Localize(tablebeast), entries),
	}
}

// - Methods --//
func (tb *BeastsTable) Add(character *characters.Character) {
	tb.Entries = append(tb.Entries, BeastEntry{
		Timestamp: time.Now().UTC(),
		Character: character,
	})
}

func (tb *BeastsTable) AddFromMappedValues(values map[string]string) error {
	if err := tb.assertEntry(values); err != nil {
		return err
	}

	character, err := characters.New(characters.Template{
		Name:        values["name"],
		Description: values["description"],
		/*
			TODO
			HIGH Changer ca avec la File driven, peu etre la jouer au Dés avec un oracle ?
		*/
		Rarity: rarity.Default(),
		Race:   values["race"],
	})

	if err != nil {
		return fmt.Errorf("failed to create character from mapped values: %w", err)
	}

	tb.Add(character)

	return nil
}

func (tb *BeastsTable) Summaries() []string {
	summaries := make([]string, 0, len(tb.Entries))
	for _, entry := range tb.Entries {
		if entry.Character == nil {
			summaries = append(summaries, t.Localize("codex.unknown_beast"))
			continue
		}
		summaries = append(summaries, fmt.Sprintf("%s — %s", entry.Character.Name, entry.Character.Description))
	}

	if len(summaries) == 0 {
		summaries = append(summaries, t.Localize("codex.no_beasts"))
	}

	return summaries
}

// - Helper --//
func (tb *BeastsTable) assertEntry(entry map[string]string) error {
	var err []error

	if entry["name"] == "" {
		err = append(err, fmt.Errorf("missing name for beast"))
	}

	if entry["description"] == "" {
		err = append(err, fmt.Errorf("missing description for beast"))
	}

	/*
		TODO
		HIGH Implmenter ça dans le assert
		if result.Kind == codexform.Beasts && !races.FindByName(values["race"]).IsBeast) {
			return t.NewError("error.non_beast_race", map[string]any{"Race": values["race"]})
		}
	*/

	if !races.Assert(entry["race"]) {
		err = append(err, fmt.Errorf("missing race for beast"))
	}
	if !classes.Assert(entry["class"]) {
		err = append(err, fmt.Errorf("missing class for beast"))
	}

	if len(err) > 0 {
		return tb.formatAssertErrors(err)
	}

	return nil
}

func (tb *BeastsTable) Ensure() {
	ensureEmbeddedTable(t.Localize(tablebeast), &tb.TableData)
}
