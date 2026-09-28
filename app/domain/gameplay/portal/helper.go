package portal

import (
	"solopg/app/domain/card/locations"
	"solopg/app/domain/gameplay/dice"
	"solopg/app/services/i18n"
)

// - Portal - //
func Teleport() (*locations.Location, error) {
	list := locations.List()

	if len(list) == 0 {
		// i18N -- register
		return nil, i18n.NewError("error.required", map[string]any{
			"Subject":  i18n.Localize("location"),
			"Property": i18n.Localize("list"),
		})
	}

	roll := dice.Roll(len(list))
	selectedLocation := list[roll-1]

	return &selectedLocation, nil
}
