package locations

import (
	"fmt"

	"solopg/internal/domain/card"
	"solopg/internal/domain/card/attributes/rarity"
	"solopg/internal/domain/card/attributes/stats"
	"solopg/internal/domain/card/attributes/variety"
)

type Location struct {
	card.Card

	Effects []stats.Effect
}

type Template struct {
	Name        string
	Description string
	Rarity      rarity.Rarity
	Variety     variety.Variety
	Effects     []stats.Effect
}

func New(params Template) (*Location, error) {
	newCard, err := card.NewCard(card.NewCardParams{
		Name:        params.Name,
		Description: params.Description,
		Rarity:      params.Rarity,
		Variety:     params.Variety,
	})

	if newCard == nil {
		return nil, fmt.Errorf("failed to create new card for location %s: %v", params.Name, err)
	}

	return &Location{
		Card:    *newCard,
		Effects: params.Effects,
	}, nil
}
