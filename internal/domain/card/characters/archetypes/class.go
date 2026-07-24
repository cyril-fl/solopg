package archetypes

import "fmt"

type OG_Class string

const (
	Adventurer OG_Class = "Adventurer"
	Craftsman  OG_Class = "Craftsman"
	Monster    OG_Class = "Monster"
)

func (c OG_Class) Validate() error {
	switch c {
	case Adventurer, Craftsman, Monster:
		return nil
	default:
		return fmt.Errorf("invalid class: %s", c)
	}
}

var PlayableClasses = []OG_Class{Adventurer}

func (c OG_Class) String() string {
	return string(c)
}

