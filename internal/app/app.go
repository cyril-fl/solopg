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
	db, err := mongo.Connect()
	if err != nil {
		return err
	}

	defer mongo.Disconnect(db)

	// campaignData, err := loadCampaignData(db)
	// if err != nil {
	// 	return err
	// }

	// engine, err := boot(db, campaignData)
	// if err != nil {
	// 	return err
	// }

	// gameUi := gameui.NewUi(gameui.UiParams{
	// 	Engine: engine,
	// 	OnSave: process.NewSaveFunc(db, engine),
	// })

	// gameUi.Start()

	return nil
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

	saves, err := db.LoadCampaign()
	if err != nil {
		return err
	}

	races := races.List()
	classes := classes.List()

	models := []tui.Step{
		{
			Model: loadsave.NewModel(saves),
			Resolve: func(ctx *tui.Context, value any) error {
				if value == nil {
					ctx.SelectedSave = nil
					return nil
				}
				ctx.SelectedSave = value.(*campaign.Campaign)
				return nil
			},
		},

				{
			Model: loadsave.NewModel(saves),
			Resolve: func(ctx *tui.Context, value any) error {
				if value == nil {
					ctx.SelectedSave = nil
					return nil
				}
				ctx.SelectedSave = value.(*campaign.Campaign)
				return nil
			},
		},

				{
			Model: loadsave.NewModel(saves),
			Resolve: func(ctx *tui.Context, value any) error {
				if value == nil {
					ctx.SelectedSave = nil
					return nil
				}
				ctx.SelectedSave = value.(*campaign.Campaign)
				return nil
			},
		},

				{
			Model: loadsave.NewModel(saves),
			Resolve: func(ctx *tui.Context, value any) error {
				if value == nil {
					ctx.SelectedSave = nil
					return nil
				}
				ctx.SelectedSave = value.(*campaign.Campaign)
				return nil
			},
		},
		{
			Model: choosename.NewModel(),
			Resolve: func(ctx *tui.Context, value any) error {
				if value == nil {
					ctx.SelectedName = ""
					return nil
				}
				ctx.SelectedName = value.(string)
				return nil
			},
		},
		{
			Model: choosearchetype.NewModel(races),
			Resolve: func(ctx *tui.Context, value any) error {

				if value == nil {
					ctx.SelectedRace = nil
					return nil
				}
				
				fmt.Printf("%+v\n", ctx)
				fmt.Printf("%+v\n", value)
				
				return nil
			},
		},
				{
			Model: choosearchetype.NewModel(classes),
			Resolve: func(ctx *tui.Context, value any) error {

				if value == nil {
					ctx.SelectedClass = nil
					return nil
				}
				fmt.Printf("%+v\n", ctx)
				fmt.Printf("%+v\n", value)
				
				return nil
			},
		},
	}
	tui.New(models).Run()

	return nil
}
