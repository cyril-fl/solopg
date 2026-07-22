package archetypes

import "fmt"

type Race string

const (
	Human Race = "Human"
	Dwarf Race = "Dwarf"

	Slime Race = "Slime"
)

func (r Race) Validate() error {
	switch r {
	case Human, Dwarf, Slime:
		return nil
	default:
		return fmt.Errorf("invalid race: %s", r)
	}
}
