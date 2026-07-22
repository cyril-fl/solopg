package stats

import (
	"solopg/domain/types"
)

type Stats struct {
	Health int // HP
	Physical int // Physical strength, dexterity, and endurance
	Mental int // Intelligence and wisdom
	Stamina int // MP
	Social int // Charisma, persuasion, and social skills
}

type Stat string

const (
    Health   Stat = "health"
    Physical Stat = "physical"
    Mental   Stat = "mental"
    Stamina  Stat = "stamina"
    Social   Stat = "social"
)

func (s Stat) Validate() bool {
	switch s {
	case Health, Physical, Mental, Stamina, Social:
		return true
	default:
		return false
	}
}

type Attribute struct {
	Attack int
	Defense int
}

type Modifier struct {
	Stat  Stat
	Value int
}

type Effect struct {
	types.Description
	Modifier Modifier
}