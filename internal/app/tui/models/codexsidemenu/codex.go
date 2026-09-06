package codexsidemenu

import (
	"fmt"
	"solopg/internal/app/game"
	"solopg/internal/app/tui/models/codexform"
	"solopg/internal/domain/card/attributes"
	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/characters/archetypes/classes"
	"solopg/internal/domain/card/characters/archetypes/races"
	"solopg/internal/domain/card/effects"
	"solopg/internal/domain/card/locations"
	"solopg/internal/domain/card/objects"
	"solopg/internal/domain/codex"
	"solopg/internal/infrastructure/t"
	"strings"
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
func addCodexEntry(engine *game.Engine, result codexform.Result) error {
	if engine == nil || engine.State == nil {
		return t.NewError("error.game_unavailable")
	}

	codexData := engine.State.Codex
	if codexData == nil {
		codexData = codex.New()
		engine.State.Codex = codexData
	}

	values := result.Values
	var err error
	switch result.Kind {
	case codexform.NPCs, codexform.Monsters:
		err = addNPCToCodex(result, codexData, values)
	case codexform.Locations:
		err = addLocotionToCodex(codexData, values)
	case codexform.Objects:
		err = addObjectToCodex(codexData, values)
	case codexform.Objectifs:
		err = addObjectiveToCodex(codexData, values)
	default:
		err = t.NewError("error.unknown_codex", map[string]any{"Kind": result.Kind})
	}

	if err != nil {
		return err
	}

	engine.AddJournalEntry("Codex", fmt.Sprintf("Nouvelle entrée ajoutée : %s", result.Kind))
	return nil
}

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

func addLocotionToCodex(codexData *codex.Codex, values map[string]string) error {
	location, err := locations.New(locations.Template{
		Name:        values["name"],
		Description: values["description"],
		Rarity:      attributes.F,
		Variety:     attributes.LocationCard,
	})
	if err != nil {
		return err
	}
	codexData.LocationsTable.AddEntry(codex.LocationsEntryTemplate{Location: location})
	return nil
}

func addObjectToCodex(codexData *codex.Codex, values map[string]string) error {
	category := objects.Category(strings.ToLower(values["category"]))
	if err := category.Validate(); err != nil {
		return err
	}
	object, err := objects.New(objects.Template{
		Name:        values["name"],
		Description: values["description"],
		Rarity:      attributes.F,
		Variety:     attributes.ArticleCard,
		Category:    category,
	})
	if err != nil {
		return err
	}
	codexData.ObjectsTable.AddObject(object)
	return nil
}

func addObjectiveToCodex(codexData *codex.Codex, values map[string]string) error {
	codexData.ObjectifsTable.AddObjectif(values["title"], values["description"])
	return nil
}
