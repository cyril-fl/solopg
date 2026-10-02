package objects

import (
	"solopg/app/cmdrun/domain/card"
	"solopg/app/cmdrun/domain/card/attributes/objectcategory"
	"solopg/app/cmdrun/domain/card/attributes/rarity"
	"solopg/app/cmdrun/domain/card/attributes/stats"
	"solopg/app/cmdrun/domain/card/attributes/variety"
	"solopg/app/shared/services/i19n"
)

type Object struct {
	card.Card

	Category objectcategory.Category
	Effects  []stats.Effect
	Pod      int
}

type Template struct {
	Name        string
	Description string
	Rarity      rarity.Rarity
	Variety     variety.Variety
	Category    objectcategory.Category
	Effects     []stats.Effect
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

	if newCard == nil {
		// i18N -- register
		return nil, i19n.NewError("error.invalid:new", map[string]any{
			"Subject": i19n.Localize("item"),
			"Error":   err,
		})
	}

	// if newCard.Variety != attributes.ArticleCard && newCard.Variety != attributes.EquipmentCard {
	// 	return nil, fmt.Errorf("invalid card variety for item: %s", newCard.Variety)
	// }

	// Check Category
	if !params.Category.Validate() {
		// i18N -- register
		return nil, i19n.NewError("error.invalid", map[string]any{
			"Subject": i19n.Localize("category"),
			"Receive": params.Category,
		})
	}

	return &Object{
		Card:     *newCard,
		Category: params.Category,
		Effects:  params.Effects,
		Pod:      params.Pod,
	}, nil
}

func (o *Object) GetName() string {
	return o.Description.Name
}

func (o *Object) GetDescription() string {
	return o.Description.Description
}
