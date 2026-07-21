package cards

type Variety string

const (
	CharacterCard Variety = "character_card"
	EquipmentCard Variety = "equipment_card"
	EventCard     Variety = "event_card"
	LocationCard  Variety = "location_card"
	ItemCard      Variety = "item_card"
)

func (c Variety) IsValid() bool {
	switch c {
	case CharacterCard, EquipmentCard, EventCard, LocationCard, ItemCard:
		return true
	default:
		return false
	}
}
