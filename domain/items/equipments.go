package items

import (
	"fmt"
	"solopg/domain/attributes"
	"solopg/domain/cards"
)

type EquipmentSlot string

const (	
	Helmet EquipmentSlot = "helmet"
    Chestplate  EquipmentSlot = "chestplate"
    Gauntlets EquipmentSlot = "gauntlets"
    Greaves EquipmentSlot = "greaves"
    Boots  EquipmentSlot = "boots"
    RightHand EquipmentSlot = "right_hand"
	LeftHand  EquipmentSlot = "left_hand"
)

func (s EquipmentSlot) Validate() error {
	switch s {
	case Helmet, Chestplate, Gauntlets, Greaves, Boots, RightHand, LeftHand:
		return nil
	default:
		return fmt.Errorf("invalid equipment slot: %s", s)
	}
}	

type Equipment struct {
	Helmet   *EquipmentGear
	Chestplate  *EquipmentGear
	Gauntlets  *EquipmentGear
	Greaves   *EquipmentGear
	Boots   *EquipmentGear
	RightHand *EquipmentGear
	LeftHand  *EquipmentGear
}

type EquipmentGear struct {
	Item
	DestinedSlot EquipmentSlot
}

type NewEquipmentParams struct {
	Name string
	Description string
	Rarity cards.Rarity
	Variety cards.Variety
	Category Category
	Effects []attributes.Effect
	Pod int
	DestinedSlot EquipmentSlot
}

func NewEquipmentGear(params NewEquipmentParams) (*EquipmentGear, error) {
	// Check Card
	newItem, err := NewItem(NewItemParams{
		Name: params.Name,
		Description: params.Description,
		Rarity: params.Rarity,
		Variety: params.Variety,
		Category: params.Category,
		Effects: params.Effects,
		Pod: params.Pod,
	})
	
	if err != nil {
		return nil, err
	}

	if newItem == nil  {
		return nil, fmt.Errorf("failed to create new item for equipment gear")
	}	

	// Check DestinedSlot
	if err := params.DestinedSlot.Validate(); err != nil {
		return nil, err
	}	

	return &EquipmentGear{
		Item: *newItem,
		DestinedSlot: params.DestinedSlot,
	}, nil
}