package objects

import "fmt"

/*
	TODO
	MEDIUM upgrade to File driven approach if not needed in game logic
*/

type Category string

const (
	Objects Category = "object"
	Weapon  Category = "weapon"
	Armor   Category = "armor"
	Potion  Category = "potion"
)

func (c Category) Validate() error {
	switch c {
	case Objects, Weapon, Armor, Potion:
		return nil
	default:
		return fmt.Errorf("invalid category: %s", c)
	}
}
