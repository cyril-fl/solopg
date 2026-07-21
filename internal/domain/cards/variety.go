package cards

import (
	"fmt"
)

type Variety string

const (
	CharacterCard Variety = "character_card"
	EquipmentCard Variety = "equipment_card"
	EventCard     Variety = "event_card"
	LocationCard  Variety = "location_card"
	ItemCard      Variety = "item_card"
)

func (c Variety) Validate() error {
	switch c {
	case CharacterCard, EquipmentCard, EventCard, LocationCard, ItemCard:
		return nil
	default:
		return fmt.Errorf("invalid variety: %s", c)
	}
}

/* TODO:
	- Create variety
 */