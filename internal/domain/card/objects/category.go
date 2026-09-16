package objects

import "fmt"

/*
	TODO upgrade to File driven approach if not needed in game logic
	MEDIUM
*/

type Category string

const (
	Weapon Category = "weapon"
	Armor  Category = "armor"
	Potion Category = "potion"
)

func (c Category) Validate() error {
	switch c {
	case Weapon, Armor, Potion:
		return nil
	default:
		return fmt.Errorf("invalid category: %s", c)
	}
}
