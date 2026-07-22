package characters

import (
	// "fmt"
	"solopg/internal/domain/card/objects"
)

type Equipment struct {
	Helmet     *objects.Gear
	Chestplate *objects.Gear
	Gauntlets  *objects.Gear
	Greaves    *objects.Gear
	Boots      *objects.Gear
	RightHand  *objects.Gear
	LeftHand   *objects.Gear
}
// 	switch s {
// 	case Helmet, Chestplate, Gauntlets, Greaves, Boots, RightHand, LeftHand:
// 		return nil
// 	default:
// 		return fmt.Errorf("invalid equipment slot: %s", s)
// 	}
// }