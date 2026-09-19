package card

import (
	"fmt"
	"solopg/internal/domain/card/attributes/description"
	"solopg/internal/domain/card/attributes/rarity"
	"solopg/internal/domain/card/attributes/variety"
)

type Card struct {
	description.Description

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
	if !params.Rarity.Validate() {
		return nil, fmt.Errorf("invalid rarity: %s", params.Rarity)
	}

	if !params.Variety.Validate() {
		return nil, fmt.Errorf("invalid variety: %s", params.Variety)
	}

	return &Card{
		Description: description.New(params.Name, params.Description),
		Rarity:      params.Rarity,
		Variety:     params.Variety,
	}, nil
}
