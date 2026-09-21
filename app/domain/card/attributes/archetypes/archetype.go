package archetypes

import "solopg/app/domain/card/attributes/stats"

type Archetype interface {
	GetName() string
	IsPlayable() bool
	GetBonus() []stats.Modifier
}
