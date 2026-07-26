package app

import (
	"solopg/internal/app/game"
	"solopg/internal/app/game/process"
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/bootui"
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

	saves, err := db.LoadCampaign()
	if err != nil {
		return err
	}

	bootUi := bootui.NewUi(saves)
	selectedSave, err := bootUi.SelectSave()
	if err != nil {
		return tui.NormalizeError(err)
	}

	campaign, err := process.ResolveCampaign(selectedSave)
	if err != nil {
		return tui.NormalizeError(err)
	}

	maybeArchives, err := db.LoadArchivesByCampaignID(campaign.ID)
	if err != nil {
		return err
	}

	archives := process.ResolveArchives(campaign.ID, maybeArchives)

	engine := game.Boot(campaign, archives)
	engine.Initialize()

	jsonlog.JsonifiedLog(engine)

	err = process.SaveGame(db, engine)
	if err != nil {
		return err
	}

	// gameUi := gameui.NewUi(engine)

	//  gameUi.Start()
	return nil
}

func Try() {
	test := effects.ListStats()

	jsonlog.JsonifiedLog(test)
}
