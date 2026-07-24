package app

import (
	"fmt"
	"solopg/internal/domain/card/characters"
	"solopg/internal/platform/jsonlog"
)

// import (
// 	"solopg/internal/app/game"
// 	"solopg/internal/app/tui"
// 	"solopg/internal/app/tui/bootui"
// 	"solopg/internal/app/tui/forgeui"
// 	"solopg/internal/app/tui/portalui"
// 	"solopg/internal/domain/campaign"
// 	"solopg/internal/infrastructure/mongo"
// )

// func Start() error {
// 	db, err := mongo.Connect()
// 	if err != nil {
// 		return err
// 	}

// 	defer mongo.Disconnect(db)

// 	saves, err := db.LoadSaves()
// 	if err != nil {
// 		return err
// 	}

// 	bootUi := bootui.NewUi(saves)
// 	selectedCampaign, err := bootUi.SelectSave()
// 	if err != nil {
// 		// TODO: mettre une loop ici
// 		if err == tui.ErrSelectionCancelled {
// 			return nil
// 		}
// 		return err
// 	}

// 	var currentCampaign *campaign.Campaign
// 	if selectedCampaign == nil {
// 		forgeUi := forgeui.New()
// 		player, err := forgeUi.CreateCharacter()
// 		if err != nil {
// 			if err == tui.ErrCreationCancelled {
// 				return nil
// 			}
// 			return err
// 		}

// 		portalUi := portalui.NewUi()
// 		location, err := portalUi.SelectLocation()
// 		if err != nil {
// 			if err == tui.ErrSelectionCancelled {
// 				return nil
// 			}
// 			return err
// 		}

// 		currentCampaign = campaign.New(campaign.Template{
// 			Player:          player,
// 			CurrentLocation: location,
// 		})

// 		db.SaveCampaign(currentCampaign)
// 	} else {
// 		currentCampaign = selectedCampaign
// 	}

// 	engine := game.Boot(currentCampaign)

// 	_ = engine // Use the engine as needed
// 	// gameUi := gameui.NewUi(engine)

// 	// return gameUi.Start()
// 	return nil
// }

func Start() error {
	// obj,err := articles.FromFile("data/template/articles/potion/heal_lvl1.yaml")
	// if err != nil {
	// 	return err
	// }

	monster, err := characters.FromFile("data/template/characters/monsters/slime.yaml")
	if err != nil {
		fmt.Println("Error loading monster:", err)
		return err
	}

	jsonlog.JsonifiedLog(monster)

	npc, err := characters.FromFile("data/template/characters/npcs/blacksmith.yaml")
	if err != nil {
		fmt.Println("Error loading npc:", err)
		return err
	}

	jsonlog.JsonifiedLog(npc)

	// location,err := locations.FromFile("data/template/locations/tavern.yaml")
	// if err != nil {
	// 	return err
	// }

	// jsonlog.JsonifiedLog(obj)
	// jsonlog.JsonifiedLog(location)

	return nil
}
