package card

import (
	"solopg/internal/domain/card/attributes"
)

type Card struct {
	attributes.Description

	Rarity  attributes.Rarity
	Variety attributes.Variety
}

type NewCardParams struct {
	Name        string
	Description string
	Rarity      attributes.Rarity
	Variety     attributes.Variety
}

func NewCard(params NewCardParams) (*Card, error) {
	// Check Rarity
	if err := params.Rarity.Validate(); err != nil {
		return nil, err
	}

	// Check Variety
	if err := params.Variety.Validate(); err != nil {
		return nil, err
	}

	return &Card{
		Description: attributes.NewDescription(params.Name, params.Description),
		Rarity:      params.Rarity,
		Variety:     params.Variety,
	}, nil
}
