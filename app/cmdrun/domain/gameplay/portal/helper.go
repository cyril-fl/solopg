package portal

import (
	"solopg/app/cmdrun/domain/card/locations"
	"solopg/app/cmdrun/domain/gameplay/dice"
	"solopg/app/shared/services/logs"
)

// - Portal - //
func Teleport() (*locations.Location, error) {
	list := locations.List()

	if len(list) == 0 {

		return nil, logs.Error("error.required", map[string]any{
			"Subject":  "location",
			"Property": "list",
		})
	}

	roll := dice.Roll(len(list))
	selectedLocation := list[roll-1]

	return &selectedLocation, nil
}
