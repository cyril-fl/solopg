package attributes

import (
	"solopg/domain/types"
)

type Stats struct {
	Health int
	Physical int
	Mental int
	Stamina int
	Social int
}

type Stat string

const (
    Health   Stat = "health" // HP
    Physical Stat = "physical" // Physical strength, dexterity, and endurance
    Mental   Stat = "mental" // Intelligence and wisdom
    Stamina  Stat = "stamina" // MP
    Social   Stat = "social" // Charisma, persuasion, and social skills
)

func (s Stat) Validate() bool {
	switch s {
	case Health, Physical, Mental, Stamina, Social:
		return true
	default:
		return false
	}
}

func (s Stat) String() string {
	return string(s)
}

type Attribute struct {
	Stat  Stat
	Value int
}

type Effect struct {
	types.Description
	Modifier Attribute
}


/* TODO:
	- Create stats
 */