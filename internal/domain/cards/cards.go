package cards

type Card struct {
	Description

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

	if err := params.Rarity.Validate(); err != nil {
		return nil, err
	}

	if err := params.Variety.Validate(); err != nil {
		return nil, err
	}
	
	return &Card{
		Description: NewDescription(params.Name, params.Description),
		Rarity:  params.Rarity,
		Variety: params.Variety,
	}, nil
}