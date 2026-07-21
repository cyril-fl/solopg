package cards

type Stat string

const (
    Health   Stat = "health" // HP
    Physical Stat = "physical" // Physical strength, dexterity, and endurance
    Mental   Stat = "mental" // Intelligence and wisdom
    Stamina  Stat = "stamina" // MP
    Social   Stat = "social" // Charisma, persuasion, and social skills
)

func (s Stat) IsValid() bool {
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
	Description
	Modifier Attribute
}

