package app

// import "solopg/internal/domain/gameplay"

import (
	"solopg/internal/app/game"
	"solopg/internal/app/tui"
	"solopg/internal/domain/campaign"
	"solopg/internal/infrastructure/mongo"

	"solopg/internal/app/tui/bootui"
	"solopg/internal/app/tui/forgeui"
	"solopg/internal/app/tui/portalui"
)

func Start() error {
	db, err := mongo.Connect()
	if err != nil {
		return err
	}

	defer mongo.Disconnect(db)

	saves, err := db.LoadSaves()
	if err != nil {
		return err
	}

	bootUi := bootui.NewUi(saves)
	selectedCampaign, err := bootUi.SelectSave()
	if err != nil {
		// TODO: mettre une loop ici
		if err == tui.ErrSelectionCancelled {
			return nil
		}
		return err
	}

	var currentCampaign *campaign.Campaign
	if selectedCampaign == nil {
		forgeUi := forgeui.New()
		player, err := forgeUi.CreateCharacter()
		if err != nil {
			if err == tui.ErrCreationCancelled {
				return nil
			}
			return err
		}

		portalUi := portalui.NewUi()
		location, err := portalUi.SelectLocation()
		if err != nil {
			if err == tui.ErrSelectionCancelled {
				return nil
			}
			return err
		}

		currentCampaign = campaign.New(campaign.Template{
			Player:          player,
			CurrentLocation: location,
		})

		db.SaveCampaign(currentCampaign)
	} else {
		currentCampaign = selectedCampaign
	}

	engine := game.Boot(currentCampaign)

	_ = engine // Use the engine as needed
	// gameUi := gameui.NewUi(engine)

	// return gameUi.Start()
	return nil
}
