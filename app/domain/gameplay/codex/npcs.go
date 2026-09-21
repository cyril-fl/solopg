package codex

import (
	"fmt"
	"solopg/app/domain/card/attributes/rarity"
	"solopg/app/domain/card/characters"
	"solopg/app/domain/card/characters/classes"
	"solopg/app/domain/card/characters/races"
	"solopg/app/services/t"
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
		Name: values["name"],
		// TODO HIGH dans la description a la main entrée les details de l'equipement ect enfin le faire un peu en RP QUOI
		Description: values["description"],
		Rarity:      rarity.Default(),
		Race:        values["race"],
		Class:       values["class"],
		/*
			TODO HIGH
			Entrée les stats, s'inspirer de ce que je fais pour race
			entre les stat le calculer comme le jouer
			mais les details de l'equipemetn ect sont en text dans la description en mode lore.

			la puissance serait calcler / representer par un jet de des dans un horacle , genre un perso badass hyper puissant serait jouer avec l'oracle legendary et un perso faible avec l'oracle common.
		*/
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
