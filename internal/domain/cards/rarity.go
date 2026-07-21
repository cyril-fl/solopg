package cards

type Rarity string

const (
    F Rarity = "F"
    E Rarity = "E"
    D Rarity = "D"
    C Rarity = "C"
    B Rarity = "B"
    A Rarity = "A"
    S Rarity = "S"
    SS Rarity = "SS"
    SSS Rarity = "SSS"
)

func (r Rarity) IsValid() bool {
	switch r {
	case F, E, D, C, B, A, S, SS, SSS:
		return true
	default:
		return false
	}
}