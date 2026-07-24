package app

import (
	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/locations"
	"solopg/internal/domain/card/objects"
	"solopg/internal/platform/jsonlog"
)

// import "solopg/internal/domain/gameplay"

/* import (
	"solopg/internal/app/game"
	"solopg/internal/app/tui"
	"solopg/internal/domain/campaign"
	"solopg/internal/infrastructure/mongo"

	"solopg/internal/app/tui/bootui"
	"solopg/internal/app/tui/forgeui"
	"solopg/internal/app/tui/portalui"
)
*/
/* func Start() error {
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
} */

func Start() error {
	// race, err := yaml.ListFromFile[archetypes.Race]("data/template/systems/archetypes/races.yaml")
	// if err != nil {
	// 	return err
	// }

	// class, err := yaml.ListFromFile[archetypes.Class]("data/template/systems/archetypes/class.yaml")
	// if err != nil {
	// 	return err
	// }

	// jsonlog.JsonifiedLog(race)
	// jsonlog.JsonifiedLog(class)


	// load.ArticleFromFile("data/template/articles/potion/heal_lvl1.yaml")
	// load.CharacterFromFile("data/template/characters/monsters/slime.yaml")
	// load.CharacterFromFile("data/template/characters/npcs/blacksmith.yaml")
	// load.LocationFromFile("data/template/locations/tavern.yaml")

	obj,err := objects.ArticleFromFile("data/template/articles/potion/heal_lvl1.yaml")
	if err != nil {
		return err
	}

	monster,err := characters.CharacterFromFile("data/template/characters/monsters/slime.yaml")
	if err != nil {
		return err
	}

	npc,err := characters.CharacterFromFile("data/template/characters/npcs/blacksmith.yaml")
	if err != nil {
		return err
	}

	location,err := locations.LocationFromFile("data/template/locations/tavern.yaml")
	if err != nil {
		return err
	}

	jsonlog.JsonifiedLog(obj)
	jsonlog.JsonifiedLog(monster)
	jsonlog.JsonifiedLog(npc)
	jsonlog.JsonifiedLog(location)	

	return nil
}

