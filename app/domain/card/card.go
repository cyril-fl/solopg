package card

import (
	"solopg/app/domain/card/attributes/description"
	"solopg/app/domain/card/attributes/rarity"
	"solopg/app/domain/card/attributes/variety"
	"solopg/app/services/i18n"
	"strings"
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
		// i18N -- register
		valid := make([]string, 0, len(rarity.List()))
		for _, r := range rarity.List() {
			valid = append(valid, i18n.Localize(r.String()))
		}

		return nil, i18n.NewError("error.unexpected:value", map[string]any{
			"Subject":  i18n.Localize("rarity"),
			"Expected": strings.Join(valid, ", "),
			"Received": params.Rarity,
		})
	}

	if !params.Variety.Validate() {
		// i18N -- register
		valid := make([]string, 0, len(rarity.List()))
		for _, v := range variety.List() {
			valid = append(valid, i18n.Localize(v.String()))
		}

		return nil, i18n.NewError("error.unexpected:value", map[string]any{
			"Subject":  i18n.Localize("variety"),
			"Expected": strings.Join(valid, ", "),
			"Received": params.Variety,
		})
	}

	return &Card{
		Description: description.New(params.Name, params.Description),
		Rarity:      params.Rarity,
		Variety:     params.Variety,
	}, nil
}
