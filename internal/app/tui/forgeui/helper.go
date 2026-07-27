package forgeui

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/attributes"
	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/characters/archetypes"
	"solopg/internal/domain/card/characters/archetypes/classes"
	"solopg/internal/domain/card/characters/archetypes/races"
	"solopg/internal/domain/card/effects"
	"solopg/internal/domain/card/objects"

	"charm.land/bubbles/v2/list"
)

func (m model) buildCharacter() (*characters.Character, error) {
	race := races.FindByName(m.selectedRace)
	class := classes.FindByName(m.selectedClass)

	baseStats := effects.BaseStats()

	// TODO: Applys somme randomness with dice roll
	if race != nil {
		baseStats.ApplyModifiers(race.GetBonus())
	}
	if class != nil {
		baseStats.ApplyModifiers(class.GetBonus())
	}

	armorClassSet := characters.FindArmorSetByName(class.ArmorSet)
	armorSet := characters.NewArmorSet(armorClassSet)

	return characters.New(characters.Template{
		Name:        m.selectedName,
		Description: "Your character",
		Rarity:      attributes.F,
		Class:       m.selectedClass,
		Race:        m.selectedRace,
		Stats:       baseStats,
		Equipment:   armorSet,
		Inventory:   []objects.Object{},
		Wallet:      characters.Wallet{Gold: 0, Silver: 0, Copper: 0},
	})
}

func makeItems[T archetypes.Archetype](data []T) []list.Item {
	items := make([]list.Item, 0, len(data))
	for _, i := range data {
		if !i.IsPlayable() {
			continue
		}

		newItem := tui.NewItem(i.GetName(), "", &i)
		items = append(items, newItem)
	}

	return items
}

func makeModel[T archetypes.Archetype](data []T) list.Model {
	items := makeItems(data)
	model := list.New(items, list.NewDefaultDelegate(), 0, 0)
	tui.ConfigureList(&model)

	return model
}
