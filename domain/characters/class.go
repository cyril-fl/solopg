package characters

import "fmt"

type Class string

const (
	Adventurer Class = "Adventurer"
	Craftsman  Class = "Craftsman"
	Monster    Class = "Monster"
)

func (c Class) Validate() error {
	switch c {
	case Adventurer, Craftsman, Monster:
		return nil
	default:
		return fmt.Errorf("invalid class: %s", c)
	}
}	