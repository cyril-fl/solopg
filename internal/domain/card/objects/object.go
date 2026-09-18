package objects

import (
	"fmt"

	"solopg/internal/domain/card"
	"solopg/internal/domain/card/attributes"
	"solopg/internal/domain/card/attributes/rarity"
	"solopg/internal/domain/card/effects"
)

type Object struct {
	card.Card

	Category Category
	Effects  []effects.Effect
	Pod      int
}

type Template struct {
	Name        string
	Description string
	Rarity      rarity.Rarity
	Variety     attributes.Variety
	Category    Category
	Effects     []effects.Effect
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

	if err != nil {
		return nil, err
	}

	if newCard == nil {
		return nil, fmt.Errorf("failed to create new card for item")
	}

	if newCard.Variety != attributes.ArticleCard && newCard.Variety != attributes.EquipmentCard {
		return nil, fmt.Errorf("invalid card variety for item: %s", newCard.Variety)
	}

	// Check Category
	if err := params.Category.Validate(); err != nil {
		return nil, err
	}

	return &Object{
		Card:     *newCard,
		Category: params.Category,
		Effects:  params.Effects,
		Pod:      params.Pod,
	}, nil
}
