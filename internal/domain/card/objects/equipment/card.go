package equipment

import (
	"fmt"
	"solopg/internal/domain/card/attributes/objectcategory"
	"solopg/internal/domain/card/attributes/potency"
	"solopg/internal/domain/card/attributes/rarity"
	"solopg/internal/domain/card/attributes/slot"
	"solopg/internal/domain/card/attributes/stats"
	"solopg/internal/domain/card/attributes/variety"
	"solopg/internal/domain/card/objects"
)

// - Equipment - //
// Set
type Equipment struct {
	List []Gear
}

func NewSet(gears []Gear) Equipment {
	set := Equipment{
		List: []Gear{},
	}

	for _, gear := range gears {
		set.AddGear(gear)
	}

	return set
}

func (set *Equipment) AddGear(gear Gear) {
	if !set.hasSlot(gear.Slot) {
		set.List = append(set.List, gear)
	}
}

func (set *Equipment) Gears() []Gear {
	return set.List
}

func (set *Equipment) hasSlot(slot slot.Slot) bool {
	for _, gear := range set.List {
		if gear.Slot == slot {
			return true
		}
	}
	return false
}

// Gear
type Gear struct {
	objects.Object

	Potency potency.Value
	Slot    slot.Slot
}

type Template struct {
	Name        string
	Description string
	Rarity      rarity.Rarity
	Variety     variety.Variety
	Category    objectcategory.Category
	Potency     potency.Value
	Effects     []stats.Effect
	Pod         int
	Slot        slot.Slot
}

func NewGear(params Template) (*Gear, error) {
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
	if !params.Slot.Validate() {
		return nil, fmt.Errorf("invalid equipment slot: %s", params.Slot)
	}

	return &Gear{
		Object:  *newItem,
		Potency: params.Potency,
		Slot:    params.Slot,
	}, nil
}
