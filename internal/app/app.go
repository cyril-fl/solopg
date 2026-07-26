package app

import (
	"solopg/internal/app/game"
	"solopg/internal/app/game/process"
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/bootui"
	"solopg/internal/app/tui/gameui"
	"solopg/internal/domain/card/effects"
	"solopg/internal/infrastructure/mongo"
	"solopg/internal/platform/jsonlog"
)

func Start() error {
	db, err := mongo.Connect()
	if err != nil {
		return err
	}

	defer mongo.Disconnect(db)

	campaignData, err := loadCampaignData(db)
	if err != nil {
		return err
	}

	engine, err := boot(db, campaignData)
	if err != nil {
		return err
	}

	gameUi := gameui.NewUi(gameui.UiParams{
		Engine: engine,
		OnSave: process.NewSaveFunc(db, engine),
	})

	gameUi.Start()

	return nil
}

func loadCampaignData(db *mongo.Mongo) (*game.CampaignData, error) {
	saves, err := db.LoadCampaign()
	if err != nil {
		return nil, err
	}

	bootUi := bootui.NewUi(saves)
	selectedSave, err := bootUi.SelectSave()
	if err != nil {
		return nil, tui.NormalizeError(err)
	}

	campaign, err := process.ResolveCampaign(selectedSave)
	if err != nil {
		return nil, tui.NormalizeError(err)
	}

	maybeArchives, err := db.LoadArchivesByCampaignID(campaign.ID)
	if err != nil {
		return nil, err
	}

	archives := process.ResolveArchives(campaign.ID, maybeArchives)

	return &game.CampaignData{
		Campaign: campaign,
		Archives: archives,
	}, nil
}

func boot(db *mongo.Mongo, data *game.CampaignData) (*game.Engine, error) {
	campaign := data.Campaign
	archives := data.Archives

	engine := game.NewEngine(campaign.ID, game.NewState(game.CampaignData{
		Campaign: campaign,
		Archives: archives,
	}))

	engine.Initialize()

	err := process.SaveGame(db, engine)
	if err != nil {
		return nil, err
	}

	return engine, nil
}

func Try() {
	test := effects.ListStats()

	jsonlog.JsonifiedLog(test)
}
