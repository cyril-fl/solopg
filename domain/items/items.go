package items

import (
	"fmt"
	"solopg/domain/cards"
	"solopg/domain/stats"
)

type Item struct {
	cards.Card

	Category Category
	Effects  []stats.Effect
	Pod      int
}

type Category string

const (
	Weapon Category = "weapon"
	Armor  Category = "armor"
	Potion Category = "potion"
)

func (c Category) Validate() error {
	switch c {
	case Weapon, Armor, Potion:
		return nil
	default:
		return fmt.Errorf("invalid category: %s", c)
	}
}

type NewItemParams struct {
	Name string
	Description string
	Rarity cards.Rarity
	Variety cards.Variety
	Category Category
	Effects []stats.Effect
	Pod int
}

func NewItem(params NewItemParams) (*Item, error) {
	// Check Card
	newCard, err := cards.NewCard(cards.NewCardParams{
		Name: params.Name,
		Description: params.Description,
		Rarity: params.Rarity,
		Variety: params.Variety,
	})
	
	if err != nil {
		return nil, err
	}

	if newCard == nil  {
		return nil, fmt.Errorf("failed to create new card for item")
	}	
	
	if newCard.Variety != cards.ArticleCard && newCard.Variety != cards.EquipmentCard {
		return nil, fmt.Errorf("invalid card variety for item: %s", newCard.Variety)
	}

	// Check Category
	if err := params.Category.Validate(); err != nil {
		return nil, err
	}

	return &Item{
		Card: *newCard,
		Category: params.Category,
		Effects: params.Effects,
		Pod: params.Pod,
	}, nil
}	