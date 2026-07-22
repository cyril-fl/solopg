package equipments

import (
	"fmt"
	"os"
	"solopg/domain/cards"
	"solopg/domain/items"
	"solopg/domain/stats"

	"solopg/domain/utils"

	"gopkg.in/yaml.v3"
)

type Slot string

const (	
	Helmet Slot = "helmet"
    Chestplate  Slot = "chestplate"
    Gauntlets Slot = "gauntlets"
    Greaves Slot = "greaves"
    Boots  Slot = "boots"
    RightHand Slot = "right_hand"
	LeftHand  Slot = "left_hand"
)

func (s Slot) Validate() error {
	switch s {
	case Helmet, Chestplate, Gauntlets, Greaves, Boots, RightHand, LeftHand:
		return nil
	default:
		return fmt.Errorf("invalid equipment slot: %s", s)
	}
}	

type Equipment struct {
	Helmet   *Gear
	Chestplate  *Gear
	Gauntlets  *Gear
	Greaves   *Gear
	Boots   *Gear
	RightHand *Gear
	LeftHand  *Gear
}

type Gear struct {
	items.Item

	Attributes stats.Attribute
	DestinedSlot Slot
}

type NewGearParams struct {
	Name string
	Description string
	Rarity cards.Rarity
	Variety cards.Variety
	Category items.Category
	Attributes stats.Attribute
	Effects []stats.Effect
	Pod int
	DestinedSlot Slot
}

func NewGear(params NewGearParams) (*Gear, error) {
	// Check Card
	newItem, err := items.NewItem(items.NewItemParams{
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

	return &Gear{
		Item: *newItem,
		Attributes: params.Attributes,
		DestinedSlot: params.DestinedSlot,
	}, nil
}

func LoadFromFile(fileAddress string) (*Gear, error) {
	data, err := os.ReadFile(fileAddress)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}
	
	var gearParams NewGearParams
	if err := yaml.Unmarshal(data, &gearParams); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	utils.JsonifiedLog(gearParams)

	gear, err := NewGear(gearParams)	
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	utils.JsonifiedLog(gear)

	return gear, nil
}		