package attributes

import "fmt"

/*
	TODO upgrade to File driven approach if not needed il game logic
	MEDIUM
*/
type Variety string

const (
	CharacterCard Variety = "character_card"
	EquipmentCard Variety = "equipment_card"
	EventCard     Variety = "event_card"
	LocationCard  Variety = "location_card"
	ArticleCard   Variety = "article_card"
)

func (c Variety) Validate() error {
	switch c {
	case CharacterCard, EquipmentCard, EventCard, LocationCard, ArticleCard:
		return nil
	default:
		return fmt.Errorf("invalid variety: %s", c)
	}
}
