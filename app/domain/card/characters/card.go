package characters

import (
	"solopg/app/domain/card"
	"solopg/app/domain/card/attributes/rarity"
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/domain/card/attributes/variety"
	"solopg/app/domain/card/characters/classes"
	"solopg/app/domain/card/characters/races"
	"solopg/app/domain/card/characters/wallet"
	"solopg/app/domain/card/objects"
	"solopg/app/domain/card/objects/equipment"
	"solopg/app/services/i18n"
)

type Character struct {
	card.Card

	Class string
	Race  string
	Stats stats.Stats

	Equipment equipment.Equipment
	Inventory []objects.Object
	Wallet    wallet.Wallet

	/*
		TODO LOW Ajouter des "effet" au personnage, ce ne serait pas des effet comme une arme qui fais x ou y mais plus des stats comme un empoissenememnt ex lier a une potion au autre
		ce serais aussi utiliser dans le formulaire de creation de perso pour simulier les effet d'un equipment qu'on a pas ecrit piece par piece.
	*/
}

type Template struct {
	Name        string
	Description string
	Rarity      rarity.Rarity
	Variety     variety.Variety
	Class       string
	Race        string
	Stats       stats.Stats
	Wallet      wallet.Wallet
	Equipment   equipment.Equipment
	Inventory   []objects.Object
}

func New(params Template) (*Character, error) {
	if !params.Variety.Validate() {
		params.Variety = variety.MakeDefault("character_card")
	}

	if !params.Rarity.Validate() {
		params.Rarity = rarity.MakeDefault("")
	}

	newCard, err := card.NewCard(card.NewCardParams{
		Name:        params.Name,
		Description: params.Description,
		Rarity:      params.Rarity,
		Variety:     params.Variety,
	})

	if newCard == nil {
		// i18N -- register
		return nil, i18n.NewError("error.invalid:new", map[string]any{
			"Subject": params.Name,
			"Error":   err,
		})
	}

	if !classes.Assert(params.Class) {
		// i18N -- register
		return nil, i18n.NewError("error.invalid", map[string]any{
			"Subject": i18n.Localize("class"),
			"Receive": params.Class,
		})
	}

	if !races.Assert(params.Race) {
		// i18N -- register
		return nil, i18n.NewError("error.invalid", map[string]any{
			"Subject": i18n.Localize("race"),
			"Receive": params.Race,
		})
	}

	return &Character{
		Card:      *newCard,
		Class:     params.Class,
		Race:      params.Race,
		Stats:     params.Stats,
		Equipment: params.Equipment,
		Inventory: params.Inventory,
		Wallet:    params.Wallet,
	}, nil
}

func (character *Character) EquipEquipement(gears []equipment.Gear) {
	for _, gear := range gears {
		character.Equipment.AddGear(gear)
	}
}

func (character *Character) EquipGear(gear equipment.Gear) {
	character.Equipment.AddGear(gear)
}
