package codexsidemenu

import (
	"fmt"
	"solopg/internal/app/tui/models/codexform"
	"solopg/internal/domain/card/attributes"
	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/characters/archetypes/classes"
	"solopg/internal/domain/card/characters/archetypes/races"
	"solopg/internal/domain/card/effects"
	"solopg/internal/domain/codex"
	"solopg/internal/infrastructure/t"
)

type link struct {
	id   id
	name string
}

type id uint8

const (
	CodexNPCs id = iota
	CodexMonsters
	CodexLocations
	CodexObjects
	CodexObjectifs
)

// TODO refactor
func (m *Model) addCodexEntry(result codexform.Result) error {
	values := result.Values
	var err error
	switch result.Kind {
	case codexform.NPCs, codexform.Monsters:
		err = addNPCToCodex(result, m.codex, values)
	case codexform.Locations:
		err = m.codex.AddLocation(values)
	case codexform.Objects:
		err = m.codex.AddObject(values)
	case codexform.Objectifs:
		err = m.codex.AddObjective(values)
	default:
		err = t.NewError("error.unknown_codex", map[string]any{"Kind": result.Kind})
	}

	if err != nil {
		return err
	}

	m.addJournalEntry(fmt.Sprintf("Nouvelle entrée ajoutée : %s", result.Kind))
	return nil
}

// TODO bouger quand le systeme de perso aura été refondu
func addNPCToCodex(result codexform.Result, codexData *codex.Codex, values map[string]string) error {
	if result.Kind == codexform.NPCs && classes.FindByName(values["class"]) == nil {
		return t.NewError("error.unknown_class", map[string]any{"Class": values["class"]})
	}
	if races.FindByName(values["race"]) == nil {
		return t.NewError("error.unknown_race", map[string]any{"Race": values["race"]})
	}
	if result.Kind == codexform.Monsters && !races.FindByName(values["race"]).IsMonster() {
		return t.NewError("error.non_monster_race", map[string]any{"Race": values["race"]})
	}
	character, err := characters.New(characters.Template{
		Name:        values["name"],
		Description: values["description"],
		Rarity:      attributes.F,
		Class:       values["class"],
		Race:        values["race"],
		Stats:       effects.BaseStats(),
	})
	if err != nil {
		return err
	}
	if result.Kind == codexform.NPCs {
		codexData.NpcsTable.AddNPC(character)
	} else {
		codexData.MonstersTable.AddMonster(character)
	}
	return nil
}
