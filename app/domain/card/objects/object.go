package objects

import (
	"fmt"
	"solopg/app/domain/card"
	"solopg/app/domain/card/attributes/objectcategory"
	"solopg/app/domain/card/attributes/rarity"
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/domain/card/attributes/variety"
)

type Object struct {
	card.Card

	Category objectcategory.Category
	Effects  []stats.Effect
	Pod      int
}

type Template struct {
	Name        string
	Description string
	Rarity      rarity.Rarity
	Variety     variety.Variety
	Category    objectcategory.Category
	Effects     []stats.Effect
	Pod         int
}

func New(params Template) (*Object, error) {
	// Check Card
	newCard, err := card.NewCard(card.NewCardParams{
		Name:        params.Name,
		Description: params.Description,
		Rarity:      params.Rarity,
		Variety:     params.Variety,
	})

	if newCard == nil {
		return nil, fmt.Errorf("failed to create new card for item %s: %v", params.Name, err)
	}

	// if newCard.Variety != attributes.ArticleCard && newCard.Variety != attributes.EquipmentCard {
	// 	return nil, fmt.Errorf("invalid card variety for item: %s", newCard.Variety)
	// }

	// Check Category
	if !params.Category.Validate() {
		return nil, fmt.Errorf("invalid category for item: %s", params.Category)
	}

	return &Object{
		Card:     *newCard,
		Category: params.Category,
		Effects:  params.Effects,
		Pod:      params.Pod,
	}, nil
}
