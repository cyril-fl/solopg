package characters

import (
	"fmt"
	"solopg/internal/domain/card/objects/gears"
	"solopg/internal/infrastructure/yaml"
)

type ArmorSet struct {
	Helmet     *gears.Gear
	Chestplate *gears.Gear
	Gauntlets  *gears.Gear
	Greaves    *gears.Gear
	Boots      *gears.Gear
	RightHand  *gears.Gear
	LeftHand   *gears.Gear
}

func NewArmorSet(set []*gears.Gear) ArmorSet {
	if len(set) == 0 {
		return ArmorSet{}
	}

	armorSet := ArmorSet{}

	for _, gear := range set {
		switch gear.Slot {
		case gears.Helmet:
			armorSet.Helmet = gear
		case gears.Chestplate:
			armorSet.Chestplate = gear
		case gears.Gauntlets:
			armorSet.Gauntlets = gear
		case gears.Greaves:
			armorSet.Greaves = gear
		case gears.Boots:
			armorSet.Boots = gear
		case gears.RightHand:
			armorSet.RightHand = gear
		case gears.LeftHand:
			armorSet.LeftHand = gear
		default:
			fmt.Printf("Invalid equipment slot: %s\n", gear.Slot)
		}
	}

	return armorSet
}

func (character *Character) EquipArmorSet(gear []*gears.Gear) {
	if character == nil || gear == nil {
		return
	}

	armorSet := NewArmorSet(gear)

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

const armorSetFilePath = "data/template/armor_sets"

func ListArmorSets() []string {
	armorSets, err := yaml.GetFolderDirectories(armorSetFilePath)
	if err != nil {
		fmt.Printf("Error reading armor set folder: %v\n", err)
		return nil
	}

	return armorSets
}

func FindArmorSetByName(name string) []*gears.Gear {
	armorSetFiles, err := yaml.GetFolderFiles(armorSetFilePath + "/" + name)

	if err != nil {
		fmt.Printf("Error reading equipment file: %v\n", err)
		return nil
	}

	if len(armorSetFiles) == 0 {
		fmt.Printf("No equipment found with name: %s\n", name)
		return nil
	}

	var armorSet []*gears.Gear
	for _, file := range armorSetFiles {
		gear, err := gears.FromFile(armorSetFilePath + "/" + name + "/" + file)
		if err != nil {
			fmt.Printf("Error loading equipment from file %s: %v\n", file, err)
			continue
		}
		armorSet = append(armorSet, gear)
	}

	return armorSet
}
