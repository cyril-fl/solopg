package locations

import (
	"solopg/app/domain/card"
	"solopg/app/domain/card/attributes/rarity"
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/domain/card/attributes/variety"
	"solopg/app/services/i18n"
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
		// i18N -- register
		return nil, i18n.NewError("error.invalid:new", map[string]any{
			"Subject": i18n.Localize("location"),
			"Error":   err,
		})
	}

	return &Location{
		Card:    *newCard,
		Effects: params.Effects,
	}, nil
}

func (l *Location) GetName() string {
	return l.Description.Name
}

func (l *Location) GetDescription() string {
	return l.Description.Description
}
