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

	err = db.RegisterArchives(archives)
	if err != nil {
		return err
	}
	return nil
}
