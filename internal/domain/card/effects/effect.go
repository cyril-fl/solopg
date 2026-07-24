package effects

import "solopg/internal/domain/card/attributes"

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
