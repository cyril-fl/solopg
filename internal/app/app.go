package app

import (
	"fmt"
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/models/choosearchetype"
	"solopg/internal/app/tui/models/choosename"
	"solopg/internal/app/tui/models/loadsave"
	"solopg/internal/domain/campaign"
	"solopg/internal/domain/card/characters/archetypes/classes"
	"solopg/internal/domain/card/characters/archetypes/races"
	"solopg/internal/infrastructure/mongo"
)

func Start() error {
	return runTUI()
}

func runTUI() error {
	db, err := mongo.Connect()
	if err != nil {
		return err
	}

	defer mongo.Disconnect(db)

	return runBootstrap(db)
}

func runBootstrap(db *mongo.Mongo) error {
	saves, err := db.LoadCampaign()
	if err != nil {
		return err
	}

	models := []tui.Step{
		{
			Model: loadsave.NewModel(saves),
			Resolve: func(ctx *tui.Context, value any) error {
				if value == nil {
					ctx.SelectedSave = nil
					return nil
				}
				selected, ok := value.(*campaign.Campaign)
				if !ok {
					return fmt.Errorf("unexpected save value %T", value)
				}
				ctx.SelectedSave = selected
				return nil
			},
		},
		{
			Model: choosename.NewModel(),
			Resolve: func(ctx *tui.Context, value any) error {
				ctx.SelectedName, _ = value.(string)
				return nil
			},
		},
		{
			Model: choosearchetype.NewModel(races.List()),
			Resolve: func(ctx *tui.Context, value any) error {
				name, ok := value.(string)
				if !ok {
					return fmt.Errorf("unexpected race value %T", value)
				}
				ctx.SelectedRace = races.FindByName(name)
				return nil
			},
		},
		{
			Model: choosearchetype.NewModel(classes.List()),
			Resolve: func(ctx *tui.Context, value any) error {
				name, ok := value.(string)
				if !ok {
					return fmt.Errorf("unexpected class value %T", value)
				}
				ctx.SelectedClass = classes.FindByName(name)
				return nil
			},
		},
	}

	return tui.New(models).Run()
}

// func loadCampaignData(db *mongo.Mongo) (*game.CampaignData, error) {

// 	campaign, err := process.ResolveCampaign(selectedSave)
// 	if err != nil {
// 		return nil, tui.NormalizeError(err)
// 	}

// 	maybeArchives, err := db.LoadArchivesByCampaignID(campaign.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	archives := process.ResolveArchives(campaign.ID, maybeArchives)

// 	return &game.CampaignData{
// 		Campaign: campaign,
// 		Archives: archives,
// 	}, nil
// }

// func boot(db *mongo.Mongo, data *game.CampaignData) (*game.Engine, error) {
// 	campaign := data.Campaign
// 	archives := data.Archives

// 	engine := game.NewEngine(campaign.ID, game.NewState(game.CampaignData{
// 		Campaign: campaign,
// 		Archives: archives,
// 	}))

// 	engine.Initialize()

// 	err := process.SaveGame(db, engine)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return engine, nil
// }

func Try() error {
	db, err := mongo.Connect()
	if err != nil {
		return err
	}

	defer mongo.Disconnect(db)

	return runBootstrap(db)
}
