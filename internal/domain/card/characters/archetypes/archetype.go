package archetypes

import "solopg/internal/domain/card/effects"

type Archetype interface {
	GetName() string
	IsPlayable() bool
	GetBonus() []effects.Modifier
}
