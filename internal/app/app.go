package app

import (
	"fmt"
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/gameui"
	"solopg/internal/app/tui/models/choosearchetype"
	"solopg/internal/app/tui/models/choosename"
	"solopg/internal/app/tui/models/loadsave"
	"solopg/internal/app/tui/portalui"
	"solopg/internal/domain/campaign"
	"solopg/internal/domain/card/characters/archetypes/classes"
	"solopg/internal/domain/card/characters/archetypes/races"
	"solopg/internal/domain/card/locations"
	"solopg/internal/infrastructure/mongo"
)

func Start() error {
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

	ui := tui.New()

	ui.
		Add(tui.Step{
			Submodel: loadsave.NewModel(saves),
			Resolve: func(ctx *tui.Context, value any) error {
				if value == nil {
					ctx.SelectedSave = nil

					ui.Add(
						tui.Step{
							Submodel: choosename.NewModel(),
							Skip: func(ctx *tui.Context) bool {
								return ctx.SelectedSave != nil
							},
							Resolve: func(ctx *tui.Context, value any) error {
								ctx.SelectedName, _ = value.(string)
								return nil
							},
						},
						tui.Step{
							Submodel: choosearchetype.NewModel(races.List()),
							Skip: func(ctx *tui.Context) bool {
								return ctx.SelectedSave != nil
							},
							Resolve: func(ctx *tui.Context, value any) error {
								name, ok := value.(string)
								if !ok {
									return fmt.Errorf("unexpected race value %T", value)
								}
								ctx.SelectedRace = races.FindByName(name)
								return nil
							},
						},
						tui.Step{
							Submodel: choosearchetype.NewModel(classes.List()),
							Skip: func(ctx *tui.Context) bool {
								return ctx.SelectedSave != nil
							},
							Resolve: func(ctx *tui.Context, value any) error {
								name, ok := value.(string)
								if !ok {
									return fmt.Errorf("unexpected class value %T", value)
								}
								ctx.SelectedClass = classes.FindByName(name)
								return nil
							},
						},
						tui.Step{
							Submodel: portalui.NewModel(),
							Skip: func(ctx *tui.Context) bool {
								return ctx.SelectedSave != nil
							},
							Resolve: func(ctx *tui.Context, value any) error {
								location, ok := value.(*locations.Location)
								if !ok {
									return fmt.Errorf("unexpected location value %T", value)
								}
								ctx.SelectedLocation = location
								return nil
							},
						},
					)

					return nil
				}
				selected, ok := value.(*campaign.Campaign)
				if !ok {
					return fmt.Errorf("unexpected save value %T", value)
				}
				ctx.SelectedSave = selected
				return nil
			},
		}).
		Add(tui.Step{
			Submodel: gameui.NewModel(gameui.UiParams{}),
		})

	return ui.Run()
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
