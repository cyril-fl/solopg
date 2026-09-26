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
		return nil, i18n.NewError("error.locations.empty", nil)
	}

	roll := dice.Roll(len(list))
	selectedLocation := list[roll-1]

	return &selectedLocation, nil
}
