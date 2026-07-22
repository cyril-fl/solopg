package objects

import (
	"fmt"
	// "os"

	// "solopg/models/cards"
	// "solopg/models/cards/objects"
	"solopg/internal/domain/card/attributes"
	"solopg/internal/domain/card/effects"
	// "solopg/internal/domain/card/slots"
	// "solopg/utils"
	// "solopg/domain/stats"
	// "gopkg.in/yaml.v3"
)

type Gear struct {
	Object

	Attributes   Attribute
	EquipmentSlot EquipmentSlot
}

type EquipmentSlot string

const (
	Helmet     EquipmentSlot = "helmet"
	Chestplate EquipmentSlot = "chestplate"
	Gauntlets  EquipmentSlot = "gauntlets"
	Greaves    EquipmentSlot = "greaves"
	Boots      EquipmentSlot = "boots"
	RightHand  EquipmentSlot = "right_hand"
	LeftHand   EquipmentSlot = "left_hand"
)

func (s EquipmentSlot) Validate() error {
	switch s {
	case Helmet, Chestplate, Gauntlets, Greaves, Boots, RightHand, LeftHand:
		return nil
	default:
		return fmt.Errorf("invalid equipment slot: %s", s)
	}
}

type NewGearParams struct {
	Name         string
	Description  string
	Rarity       attributes.Rarity
	Variety      attributes.Variety
	Category     Category
	Attributes   Attribute
	Effects      []effects.Effect
	Pod          int
	EquipmentSlot EquipmentSlot
}

func NewGear(params NewGearParams) (*Gear, error) {
	// Check Card
	newItem, err := NewObject(NewObjectParams{
		Name:        params.Name,
		Description: params.Description,
		Rarity:      params.Rarity,
		Variety:     params.Variety,
		Category:    params.Category,
		Effects:     params.Effects,
		Pod:         params.Pod,
	})

	if err != nil {
		return nil, err
	}

	if newItem == nil {
		return nil, fmt.Errorf("failed to create new item for equipment gear")
	}

	// Check DestinedSlot
	if err := params.EquipmentSlot.Validate(); err != nil {
		return nil, err
	}

	return &Gear{
		Object:       *newItem,
		Attributes:   params.Attributes,
		EquipmentSlot: params.EquipmentSlot,
	}, nil
}

// func LoadFromFile(fileAddress string) (*Gear, error) {
// 	data, err := os.ReadFile(fileAddress)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		return nil, err
// 	}

// 	var gearParams NewGearParams
// 	if err := yaml.Unmarshal(data, &gearParams); err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		return nil, err
// 	}

// 	utils.JsonifiedLog(gearParams)

// 	gear, err := NewGear(gearParams)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		return nil, err
// 	}

// 	utils.JsonifiedLog(gear)

// 	return gear, nil
// }
