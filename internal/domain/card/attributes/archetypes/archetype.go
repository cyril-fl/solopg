package archetypes

import "solopg/internal/domain/card/attributes/stats"

type Archetype interface {
	GetName() string
	IsPlayable() bool
	GetBonus() []stats.Modifier
}
