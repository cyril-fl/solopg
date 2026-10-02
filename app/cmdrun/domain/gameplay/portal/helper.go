package portal

import (
	"solopg/app/cmdrun/domain/card/locations"
	"solopg/app/cmdrun/domain/gameplay/dice"
	"solopg/app/shared/services/i19n"
)

// - Portal - //
func Teleport() (*locations.Location, error) {
	list := locations.List()

	if len(list) == 0 {
		// i18N -- register
		return nil, i19n.NewError("error.required", map[string]any{
			"Subject":  i19n.Localize("location"),
			"Property": i19n.Localize("list"),
		})
	}

	roll := dice.Roll(len(list))
	selectedLocation := list[roll-1]

	return &selectedLocation, nil
}
