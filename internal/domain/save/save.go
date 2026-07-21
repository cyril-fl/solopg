package save

import (
	"github.com/google/uuid"

	"solopg/internal/domain/cards"
)
type Save struct {
	CampainID uuid.UUID
	Character cards.Character
	Location  cards.Location
}