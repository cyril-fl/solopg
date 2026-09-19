package gears

import (
	"fmt"

	"solopg/internal/domain/card/attributes/rarity"
	"solopg/internal/domain/card/attributes/stats"
	"solopg/internal/domain/card/attributes/variety"
	"solopg/internal/domain/card/objects"
	"solopg/internal/infrastructure/yaml"
)

type Gear struct {
	objects.Object

	Attributes objects.Attribute
	Slot       Slot `yaml:"destinedslot"`
}

type Slot string

const (
	Helmet     Slot = "helmet"
	Chestplate Slot = "chestplate"
	Gauntlets  Slot = "gauntlets"
	Greaves    Slot = "greaves"
	Boots      Slot = "boots"
	RightHand  Slot = "right_hand"
	LeftHand   Slot = "left_hand"
)

func (s Slot) Validate() error {
	switch s {
	case Helmet, Chestplate, Gauntlets, Greaves, Boots, RightHand, LeftHand:
		return nil
	default:
		return fmt.Errorf("invalid equipment slot: %s", s)
	}
}

type Template struct {
	Name        string
	Description string
	Rarity      rarity.Rarity
	Variety     variety.Variety
	Category    objects.Category
	Attributes  objects.Attribute
	Effects     []stats.Effect
	Pod         int
	Slot        Slot `yaml:"destinedslot"`
}

func New(params Template) (*Gear, error) {
	// Check Card
	newItem, err := objects.New(objects.Template{
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

	// Check EquipmentSlot
	if err := params.Slot.Validate(); err != nil {
		return nil, err
	}

	return &Gear{
		Object:     *newItem,
		Attributes: params.Attributes,
		Slot:       params.Slot,
	}, nil
}

func FromFile(fileAddress string) (*Gear, error) {
	params, err := yaml.LoadFromFile[Template](fileAddress)
	if err != nil {
		return nil, err
	}

	return New(*params)
}
