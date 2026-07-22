package save

import (
	"time"

	"github.com/google/uuid"

	"solopg/domain/characters"
	"solopg/domain/locations"
)
type Save struct {
	CampainID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Character characters.Character
	Location  locations.Location
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}