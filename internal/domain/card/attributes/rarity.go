package attributes

import "fmt"

// TODO: upgrade to File driven approach if not needed in game logic
type Rarity string

const (
	F   Rarity = "F"
	E   Rarity = "E"
	D   Rarity = "D"
	C   Rarity = "C"
	B   Rarity = "B"
	A   Rarity = "A"
	S   Rarity = "S"
	SS  Rarity = "SS"
	SSS Rarity = "SSS"
)

func (r Rarity) Validate() error {
	switch r {
	case F, E, D, C, B, A, S, SS, SSS:
		return nil
	default:
		return fmt.Errorf("invalid rarity: %s", r)
	}
}
