package process

import (
	"solopg/internal/app/game"
	"solopg/internal/infrastructure/mongo"
)

func SaveGame(db *mongo.Mongo, engine *game.Engine) error {
	archives := engine.ExportArchives()
	campaign := engine.ExportCampaign()

	err := db.RegisterCampaign(campaign)
	if err != nil {
		return err
	}

	return db.RegisterArchives(archives)
}

func NewSaveFunc(db *mongo.Mongo, engine *game.Engine) func() error {
	return func() error {
		return SaveGame(db, engine)
	}
}
