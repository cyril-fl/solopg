package card

import (
	"fmt"
	"solopg/internal/domain/card/attributes"
	"solopg/internal/domain/card/attributes/rarity"
)

type Card struct {
	attributes.Description

	Rarity  rarity.Rarity
	Variety attributes.Variety
}

type NewCardParams struct {
	Name        string
	Description string
	Rarity      rarity.Rarity
	Variety     attributes.Variety
}

func NewCard(params NewCardParams) (*Card, error) {
	// Check Rarity
	if !params.Rarity.Validate() {
		return nil, fmt.Errorf("invalid rarity: %s", params.Rarity)
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
