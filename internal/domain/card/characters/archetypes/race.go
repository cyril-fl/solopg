package archetypes

import "fmt"

type Race string

const (
	Elf   Race = "Elf"
	Human Race = "Human"
	Dwarf Race = "Dwarf"

	Slime Race = "Slime"
)

func (r Race) Validate() error {
	switch r {
	case Elf, Human, Dwarf, Slime:
		return nil
	default:
		return fmt.Errorf("invalid race: %s", r)
	}
}

var PlayableRaces = []Race{Elf, Human, Dwarf}

func (r Race) String() string {
	return string(r)
}
