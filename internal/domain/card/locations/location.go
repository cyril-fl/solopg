package locations

import (
	"fmt"

	"solopg/internal/domain/card"
	"solopg/internal/domain/card/attributes"
	"solopg/internal/domain/card/effects"
)

type Location struct {
	card.Card

	Effects []effects.Effect
}

type LocationTemplate struct {
	Name        string
	Description string
	Rarity      attributes.Rarity
	Variety     attributes.Variety
	Effects     []effects.Effect
}

func NewLocation(params LocationTemplate) (*Location, error) {
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
		return nil, fmt.Errorf("failed to create new card for location")
	}

	return &Location{
		Card:    *newCard,
		Effects: params.Effects,
	}, nil
}
