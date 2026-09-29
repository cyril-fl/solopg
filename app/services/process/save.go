package process

import (
	"solopg/app/services/game"
	"solopg/app/services/mongo"
	"solopg/app/services/mongo/repository"
)

func SaveGame(db *mongo.Mongo, engine *game.Engine) error {
	archives := engine.ExportArchives()
	archivesrepo := repository.NewArchivesRepository().SetDb(db)
	campaign := engine.ExportCampaign()
	campaignrepo := repository.NewCampaignRepository().SetDb(db)

	err := campaignrepo.RegisterCampaign(campaign)
	if err != nil {
		return err
	}

	return archivesrepo.RegisterArchives(archives)
}

func NewSaveFunc(db *mongo.Mongo, engine *game.Engine) func() error {
	return func() error {
		return SaveGame(db, engine)
	}
}
