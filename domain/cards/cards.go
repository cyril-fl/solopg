package cards

import (
	"solopg/domain/types"
)

type Card struct {
	types.Description

	Rarity Rarity
	Variety Variety
}	

type NewCardParams struct {
	Name string
	Description string
	Rarity Rarity
	Variety Variety
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
		Description: types.NewDescription(params.Name, params.Description),
		Rarity:  params.Rarity,
		Variety: params.Variety,
	}, nil
}