package codex

import (
	"fmt"
	"solopg/internal/app/game"
	"solopg/internal/app/tui"
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

	"charm.land/bubbles/v2/list"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

// TODO refactor
type CodexLink struct {
	name string
	Kind CodexKind
}

type CodexKind uint8

const (
	CodexNPCs CodexKind = iota
	CodexMonsters
	CodexLocations
	CodexObjects
	CodexObjectifs
)

func MakeCodexListModel(width int, height int) list.Model {
	links := []CodexLink{
		{name: t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "codex.npcs"}), Kind: CodexNPCs},
		{name: t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "codex.monsters"}), Kind: CodexMonsters},
		{name: t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "codex.locations"}), Kind: CodexLocations},
		{name: t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "codex.objects"}), Kind: CodexObjects},
		{name: t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "codex.objectives"}), Kind: CodexObjectifs},
	}
	items := make([]list.Item, 0, len(links))
	for _, link := range links {
		items = append(items, tui.NewItem(link.name, "", link))
	}

	model := list.New(items, list.NewDefaultDelegate(), width, height)
	tui.ConfigureList(&model)

	return model
}

func AddCodexEntry(engine *game.Engine, result codexform.Result) error {
	if engine == nil || engine.State == nil {
		return t.NewError(&goi18n.LocalizeConfig{MessageID: "error.game_unavailable"})
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
		err = t.NewError(&goi18n.LocalizeConfig{MessageID: "error.unknown_codex", TemplateData: map[string]any{"Kind": result.Kind}})
	}

	if err != nil {
		return err
	}

	engine.AddJournalEntry("Codex", fmt.Sprintf("Nouvelle entrée ajoutée : %s", result.Kind))
	return nil
}

func addNPCToCodex(result codexform.Result, codexData *codex.Codex, values map[string]string) error {
	if result.Kind == codexform.NPCs && classes.FindByName(values["class"]) == nil {
		return t.NewError(&goi18n.LocalizeConfig{MessageID: "error.unknown_class", TemplateData: map[string]any{"Class": values["class"]}})
	}
	if races.FindByName(values["race"]) == nil {
		return t.NewError(&goi18n.LocalizeConfig{MessageID: "error.unknown_race", TemplateData: map[string]any{"Race": values["race"]}})
	}
	if result.Kind == codexform.Monsters && !races.FindByName(values["race"]).IsMonster() {
		return t.NewError(&goi18n.LocalizeConfig{MessageID: "error.non_monster_race", TemplateData: map[string]any{"Race": values["race"]}})
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

func RenderCodexPage(engine *game.Engine, link CodexLink) string {
	title := link.name
	var entries []string
	if engine != nil && engine.State != nil && engine.State.Codex != nil {
		codexData := engine.State.Codex.EnsureInitialized()
		switch link.Kind {
		case CodexNPCs:
			entries = codexData.NpcsTable.Summaries()
		case CodexMonsters:
			entries = codexData.MonstersTable.Summaries()
		case CodexLocations:
			entries = codexData.LocationsTable.Summaries()
		case CodexObjects:
			entries = codexData.ObjectsTable.Summaries()
		case CodexObjectifs:
			entries = codexData.ObjectifsTable.Summaries()
		}
	}
	if len(entries) == 0 {
		entries = []string{t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: "codex.empty"})}
	}
	return title + "\n\n" + strings.Join(entries, "\n\n")
}