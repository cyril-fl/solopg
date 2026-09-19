package codex

import (
	"fmt"
	"solopg/internal/domain/card/attributes/rarity"
	"solopg/internal/domain/card/characters"
	"solopg/internal/infrastructure/t"
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
		TableData: newTable(t.Localize(tablenpcs), entries),
	}
}

// - Methods --//
func (n *NpcsTable) Add(character *characters.Character) {
	n.Entries = append(n.Entries, NpcsEntry{
		Timestamp: time.Now().UTC(),
		Character: character,
	})
}

func (n *NpcsTable) AddFromMappedValues(values map[string]string) error {
	if err := n.assertEntry(values); err != nil {
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
		Class:  values["class"],
	})

	if err != nil {
		return fmt.Errorf("failed to create character from mapped values: %w", err)
	}

	n.Add(character)

	return nil
}

func (n *NpcsTable) Summaries() []string {
	summaries := make([]string, 0, len(n.Entries))
	for _, entry := range n.Entries {
		if entry.Character == nil {
			summaries = append(summaries, t.Localize("codex.unknown_npc"))
			continue
		}
		summaries = append(summaries, fmt.Sprintf("%s — %s - %s - %s", entry.Character.Name, entry.Character.Race, entry.Character.Class, entry.Character.Description))
	}

	if len(summaries) == 0 {
		summaries = append(summaries, t.Localize("codex.no_npcs"))
	}

	return summaries
}

// - Helper --//
func (n *NpcsTable) assertEntry(entry map[string]string) error {
	var err []error
	if entry["name"] == "" {
		err = append(err, fmt.Errorf("name is empty"))
	}
	if entry["race"] == "" {
		err = append(err, fmt.Errorf("race is empty"))
	}
	if entry["class"] == "" {
		err = append(err, fmt.Errorf("class is empty"))
	}
	if entry["description"] == "" {
		err = append(err, fmt.Errorf("description is empty"))
	}

	if len(err) > 0 {
		return n.formatAssertErrors(err)
	}

	return nil
}

func (n *NpcsTable) Ensure() {
	ensureEmbeddedTable(t.Localize(tablenpcs), &n.TableData)
}
