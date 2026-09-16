package codex

import (
	"fmt"
	"solopg/internal/domain/card/attributes"
	"solopg/internal/domain/card/characters"
	"solopg/internal/infrastructure/t"
	"time"
)

// -- Table --//
/*
	TODO
	LOW renommer en beast
*/
type MonstersTable struct {
	*TableData[MonsterEntry]
}

type MonsterEntry struct {
	Timestamp time.Time `bson:"timestamp"`
	Character *characters.Character `bson:"character"`
}

var tablemonster = "codex.monsters"

func NewMonstersTable(entries []MonsterEntry) *MonstersTable {
	if entries == nil {
		entries = []MonsterEntry{}
	}

	return &MonstersTable{
		TableData: newTable(t.Localize(tablemonster), entries),
	}
}

// -- Methods --//
func (m *MonstersTable) Add(character *characters.Character) {
	m.Entries = append(m.Entries, MonsterEntry{
		Timestamp: time.Now().UTC(),
		Character: character,
	})
}

func (m *MonstersTable) AddFromMappedValues(values map[string]string) error {
	if err := m.assertEntry(values); err != nil {
		return err
	}

	character, err := characters.New(characters.Template{
		Name:        values["name"],
		Description: values["description"],
		/*
		TODO 
		LOW Changer ca avec la File driven, peu etre la jouer au Dés avec un oracle ?
		*/
		Rarity:      attributes.A, 
		Race:        values["race"],
		// Class:       values["class"],
	})

	if err != nil {
		return fmt.Errorf("failed to create character from mapped values: %w", err)
	}

	m.Add(character)

	return nil
}

func (m *MonstersTable) Summaries() []string {
	summaries := make([]string, 0, len(m.Entries))
	for _, entry := range m.Entries {
		if entry.Character == nil {
			summaries = append(summaries, t.Localize("codex.unknown_monster"))
			continue
		}
		summaries = append(summaries, fmt.Sprintf("%s — %s", entry.Character.Name, entry.Character.Description))
	}

	if len(summaries) == 0 {
		summaries = append(summaries, t.Localize("codex.no_monsters"))
	}

	return summaries
}

// -- Helper --//
func (m *MonstersTable) assertEntry(entry map[string]string) error {
	var err []error

	if entry["name"] == "" {
		err = append(err, fmt.Errorf("missing name for monster"))
	}

	if entry["description"] == "" {
		err = append(err, fmt.Errorf("missing description for monster"))
	}

	/*
		TODO
		HIGH Implmenter ça dans le assert
		if result.Kind == codexform.NPCs && classes.FindByName(values["class"]) == nil {
			return t.NewError("error.unknown_class", map[string]any{"Class": values["class"]})
		}
		if races.FindByName(values["race"]) == nil {
			return t.NewError("error.unknown_race", map[string]any{"Race": values["race"]})
		}
		if result.Kind == codexform.Monsters && !races.FindByName(values["race"]).IsMonster() {
			return t.NewError("error.non_monster_race", map[string]any{"Race": values["race"]})
		}
	*/

	if entry["race"] == "" {
		err = append(err, fmt.Errorf("missing race for monster"))
	}

	// if entry["class"] !== "" {
	// 	err = append(err, fmt.Errorf("missing class for monster"))
	// }

	if len(err) > 0 {
		return m.formatAssertErrors(err)
	}

	return nil
}

func (m *MonstersTable) ensure() {
	ensureEmbeddedTable(t.Localize(tablemonster), &m.TableData)
}
