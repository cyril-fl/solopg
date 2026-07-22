package characters

import (
	"fmt"
	"solopg/internal/domain/card/objects"
)

type Equipment struct {
	Helmet     *objects.Gear
	Chestplate *objects.Gear
	Gauntlets  *objects.Gear
	Greaves    *objects.Gear
	Boots      *objects.Gear
	RightHand  *objects.Gear
	LeftHand   *objects.Gear
}

func (character *Character) SetEquipment(gear []*objects.Gear) {
	if character == nil {
		return
	}
	for _, g := range gear {
		character.SetEquipmentSlot(g)
	}
}

func (character *Character) SetEquipmentSlot(gear *objects.Gear) {
	if character == nil || gear == nil {
		return
	}

	switch gear.EquipmentSlot {
	case objects.Helmet:
		character.Equipment.Helmet = gear
	case objects.Chestplate:
		character.Equipment.Chestplate = gear
	case objects.Gauntlets:
		character.Equipment.Gauntlets = gear
	case objects.Greaves:
		character.Equipment.Greaves = gear
	case objects.Boots:
		character.Equipment.Boots = gear
	case objects.RightHand:
		character.Equipment.RightHand = gear
	case objects.LeftHand:
		character.Equipment.LeftHand = gear
	default:
		fmt.Printf("Invalid equipment slot: %s\n", gear.EquipmentSlot)
	}
}
