package codex

import (
	"fmt"
	"solopg/app/domain/card/attributes/rarity"
	"solopg/app/domain/card/characters"
	"solopg/app/domain/card/characters/classes"
	"solopg/app/domain/card/characters/races"
	"solopg/app/services/process/generatestats"
	"solopg/app/services/t"
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

	generator := generatestats.NewGenerator(values)
	generator.GenerateFromValues(generatestats.FromValuesParams{
		Randomness: true,
	})

	if generator.HasError() {
		return fmt.Errorf("failed to generate stats from values: %w", generator.GetError())
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
		return fmt.Errorf("failed to create character from mapped values: %w", err)
	}

	n.Add(character)

	return nil
}

func (n *NpcsTable) Summaries() []string {
	content := strings.Builder{}

	if len(n.Entries) == 0 {
		content.WriteString(t.Localize("codex.no_npcs"))
		return []string{content.String()}
	}

	for _, entry := range n.Entries {
		if entry.Character == nil {
			content.WriteString(t.Localize("codex.unknown_npc"))
			content.WriteString("\n\n")
			continue
		}
		content.WriteString(entry.Character.Name)
		content.WriteString(" - ")
		content.WriteString(entry.Character.Race)
		content.WriteString(" - ")
		content.WriteString(entry.Character.Class)
		content.WriteString("\n")
		content.WriteString(entry.Character.Stats.String())
		content.WriteString("\n\n")
	}

	return []string{content.String()}
}

// - Helper --//
func (n *NpcsTable) assertEntry(entry map[string]string) error {
	var err []error
	if entry["name"] == "" {
		err = append(err, fmt.Errorf("name is empty"))
	}
	if entry["description"] == "" {
		err = append(err, fmt.Errorf("description is empty"))
	}
	if !races.Assert(entry["race"]) {
		err = append(err, fmt.Errorf("race is empty"))
	}
	if !classes.Assert(entry["class"]) {
		err = append(err, fmt.Errorf("class is empty"))
	}

	if len(err) > 0 {
		return n.formatAssertErrors(err)
	}

	return nil
}

func (n *NpcsTable) Ensure() {
	ensureEmbeddedTable(t.Localize(tablenpcs), &n.TableData)
}
