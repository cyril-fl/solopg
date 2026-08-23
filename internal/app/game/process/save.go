package process

import (
	"solopg/internal/app/game"
	"solopg/internal/infrastructure/mongo"
)

// DEPRECATED

func SaveGame(db *mongo.Mongo, engine *game.Engine) error {
	archives := engine.ExportArchives()
	campaign := engine.ExportCampaign()

	err := db.RegisterCampaign(campaign)
	if err != nil {
		return err
	}

	err = db.RegisterArchives(archives)
	if err != nil {
		return err
	}
	return nil
}

func NewSaveFunc(db *mongo.Mongo, engine *game.Engine) func() error {
	return func() error {
		return SaveGame(db, engine)
	}
}
