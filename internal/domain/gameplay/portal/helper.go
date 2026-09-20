package portal

import (
	"solopg/internal/domain/card/locations"
	"solopg/internal/domain/gameplay/dice"
	"solopg/internal/infrastructure/t"
)

// TODO LOW voir si on peu pas rendre ca fn() *Location ou juste Location
// - Portal - //
func Teleport() (*locations.Location, error) {
	list := locations.List()

	if len(list) == 0 {
		return nil, t.NewError("error.locations.empty", nil)
	}

	roll := dice.Roll(len(list))
	selectedLocation := list[roll-1]

	return &selectedLocation, nil
}
