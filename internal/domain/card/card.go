package card

import (
	"fmt"
	"solopg/internal/domain/card/attributes"
	"solopg/internal/domain/card/attributes/rarity"
	"solopg/internal/domain/card/attributes/variety"
)

type Card struct {
	attributes.Description

	Rarity  rarity.Rarity
	Variety variety.Variety
}

type NewCardParams struct {
	Name        string
	Description string
	Rarity      rarity.Rarity
	Variety     variety.Variety
}

func NewCard(params NewCardParams) (*Card, error) {
	// Check Rarity
	if !params.Rarity.Validate() {
		return nil, fmt.Errorf("invalid rarity: %s", params.Rarity)
	}

	// Check Variety
	if !params.Variety.Validate() {
		return nil, fmt.Errorf("invalid variety: %s", params.Variety)
	}

	return &Card{
		Description: attributes.NewDescription(params.Name, params.Description),
		Rarity:      params.Rarity,
		Variety:     params.Variety,
	}, nil
}
