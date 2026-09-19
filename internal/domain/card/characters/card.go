package characters

import (
	"fmt"
	"slices"
	"solopg/internal/domain/card"
	"solopg/internal/domain/card/attributes/rarity"
	"solopg/internal/domain/card/attributes/stats"
	"solopg/internal/domain/card/attributes/variety"
	"solopg/internal/domain/card/characters/classes"
	"solopg/internal/domain/card/characters/equipment"
	"solopg/internal/domain/card/characters/races"
	"solopg/internal/domain/card/characters/wallet"
	"solopg/internal/domain/card/objects"
	"solopg/internal/domain/card/objects/gears"
)

type Character struct {
	card.Card

	Class string
	Race  string
	Stats stats.Stats

	Equipment equipment.ArmorSet
	Inventory []objects.Object
	Wallet    wallet.Wallet
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
	Equipment   equipment.ArmorSet
	Inventory   []objects.Object
}

func New(params Template) (*Character, error) {
	if !params.Variety.Validate() {
		params.Variety = variety.MakeDefault("character_card")
	}

	// Check Card
	newCard, err := card.NewCard(card.NewCardParams{
		Name:        params.Name,
		Description: params.Description,
		Rarity:      params.Rarity,
		Variety:     params.Variety,
	})

	if newCard == nil {
		return nil, fmt.Errorf("failed to create new card for character %s: %v", params.Name, err)
	}

	if !classes.Assert(params.Class) {
		return nil, fmt.Errorf("invalid class for character: %s", params.Class)
	}

	// // Check Race
	if !slices.Contains(races.ListNames(), params.Race) {
		return nil, fmt.Errorf("invalid race for character: %s", params.Race)
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

func (character *Character) EquipGearList(gear []*gears.Gear) {
	if character == nil || gear == nil {
		return
	}

	armorSet := equipment.NewArmorSet(gear)

	character.Equipment = armorSet
}

func (character *Character) EquipGear(gear *gears.Gear) {
	if character == nil || gear == nil {
		return
	}

	switch gear.Slot {
	case gears.Helmet:
		character.Equipment.Helmet = gear
	case gears.Chestplate:
		character.Equipment.Chestplate = gear
	case gears.Gauntlets:
		character.Equipment.Gauntlets = gear
	case gears.Greaves:
		character.Equipment.Greaves = gear
	case gears.Boots:
		character.Equipment.Boots = gear
	case gears.RightHand:
		character.Equipment.RightHand = gear
	case gears.LeftHand:
		character.Equipment.LeftHand = gear
	default:
		fmt.Printf("Invalid equipment slot: %s\n", gear.Slot)
	}
}