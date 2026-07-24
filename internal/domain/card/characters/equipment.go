package characters

import (
	"fmt"
	"solopg/internal/domain/card/objects/gears"
)

type Equipment struct {
	Helmet     *gears.Gear
	Chestplate *gears.Gear
	Gauntlets  *gears.Gear
	Greaves    *gears.Gear
	Boots      *gears.Gear
	RightHand  *gears.Gear
	LeftHand   *gears.Gear
}

func (character *Character) SetEquipment(gear []*gears.Gear) {
	if character == nil {
		return
	}
	for _, g := range gear {
		character.SetEquipmentSlot(g)
	}
}

func (character *Character) SetEquipmentSlot(gear *gears.Gear) {
	if character == nil || gear == nil {
		return
	}

	switch gear.EquipmentSlot {
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
		fmt.Printf("Invalid equipment slot: %s\n", gear.EquipmentSlot)
	}
}
