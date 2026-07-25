package effects

import (
	"solopg/internal/domain/card/attributes"
)

// TODO: upgrade to File driven approach if not needed in game logic
type Stats struct {
	Health   int // HP
	Physical int // Physical strength, dexterity, and endurance
	Mental   int // Intelligence and wisdom
	Stamina  int // MP
	Social   int // Charisma, persuasion, and social skills
}

type Stat string

// // TODO: upgrade to File driven approach if Stat are
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

type Modifier struct {
	Stat  Stat
	Value int
}

type Effect struct {
	attributes.Description
	Modifier Modifier
}

func (s *Stats) ApplyModifier(mod Modifier) {
	switch mod.Stat {
	case Health:
		s.Health += mod.Value
	case Physical:
		s.Physical += mod.Value
	case Mental:
		s.Mental += mod.Value
	case Stamina:
		s.Stamina += mod.Value
	case Social:
		s.Social += mod.Value
	}
}

func (s *Stats) ApplyModifiers(mods []Modifier) {
	for _, mod := range mods {
		s.ApplyModifier(mod)
	}
}
