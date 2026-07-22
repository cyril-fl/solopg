package save

import (
	// 	"go-compta/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open("save.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = database.AutoMigrate(
		&Save{},
	)
	if err != nil {
		return nil, err
	}

	return database, nil
}
