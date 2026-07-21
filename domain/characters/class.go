package characters

import "fmt"

type Class string

const (
	Adventurer Class = "Adventurer"
)

func (c Class) Validate() error {
	switch c {
	case Adventurer:
		return nil
	default:
		return fmt.Errorf("invalid class: %s", c)
	}
}	