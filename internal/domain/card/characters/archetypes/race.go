package archetypes

import "fmt"

type OG_Race string

const (
	Elf   OG_Race = "Elf"
	Human OG_Race = "Human"
	Dwarf OG_Race = "Dwarf"

	Slime OG_Race = "Slime"
)

func (r OG_Race) Validate() error {
	switch r {
	case Elf, Human, Dwarf, Slime:
		return nil
	default:
		return fmt.Errorf("invalid race: %s", r)
	}
}

var PlayableRaces = []OG_Race{Elf, Human, Dwarf}

func (r OG_Race) String() string {
	return string(r)
}

