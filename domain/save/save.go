package save

import (
	"github.com/google/uuid"

	"solopg/domain/cards"
	"solopg/domain/characters"
)
type Save struct {
	CampainID uuid.UUID
	Character characters.Character
	Location  cards.Location
}