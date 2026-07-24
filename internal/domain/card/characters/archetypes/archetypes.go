package archetypes

type Archetype struct {
	Name     string
	Playable bool
}

func (a Archetype) String() string {
	return a.Name
}

type Race Archetype

type Class Archetype
