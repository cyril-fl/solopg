package equipment

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
/*
TODO
MEDIUM refactor avec cache
*/

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
