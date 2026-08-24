package app

import (
	"fmt"
	"solopg/internal/app/game"
	"solopg/internal/app/game/process"
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/models"
	"solopg/internal/app/tui/models/choosearchetype"
	"solopg/internal/app/tui/models/choosename"
	"solopg/internal/app/tui/models/gameui"
	"solopg/internal/app/tui/models/loadsave"
	"solopg/internal/app/tui/models/portal"
	"solopg/internal/domain/campaign"
	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/characters/archetypes/classes"
	"solopg/internal/domain/card/characters/archetypes/races"
	"solopg/internal/domain/card/locations"
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/mongo"
	"solopg/internal/platform/jsonlog"
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
				selected, ok := value.(*campaign.Campaign)
				if !ok {
					return fmt.Errorf("unexpected save value %T", value)
				}

				ctx.SelectedSave = selected
				if selected == nil {
					ui.Insert(
						tui.Step{
							Submodel: choosename.NewModel(),
							Resolve: func(ctx *tui.Context, value any) error {
								ctx.SelectedName, _ = value.(string)
								return nil
							},
						},
						tui.Step{
							Submodel: choosearchetype.NewModel(races.List()),
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
							Submodel: portal.NewModel(),
							Resolve: func(ctx *tui.Context, value any) error {
								location, ok := value.(*locations.Location)
								if !ok {
									return fmt.Errorf("unexpected location value %T", value)
								}
								ctx.SelectedLocation = location
								return nil
							},
						},
						tui.Step{
							Submodel: models.Resolve(),
							Resolve: func(ctx *tui.Context, value any) error {
								engine, err := buildEngine(db, ctx)
								if err != nil {
									return err
								}

								ui.Insert(tui.Step{
									Submodel: gameui.NewModel(gameui.UiParams{
										Engine: engine,
										OnSave: process.NewSaveFunc(db, engine),
									}),
								})

								return nil
							},
						},
					)
				}

				return nil
			},
		})

	return ui.Run()
}

func Try() error {
	c := config.Load()

	jsonlog.JsonifiedLog(c)

	return nil
}



// ⚠️⚠️⚠️ 

// TODO Refacto comme ResolveArchives


func buildArchives(db *mongo.Mongo, selectedCampaign *campaign.Campaign) (*campaign.Archives, error) {
	maybeArchives, err := db.LoadArchivesByCampaignID(selectedCampaign.ID)
	if err != nil {
		return nil, err
	}

	// TODO Refacto
	archives := process.ResolveArchives(selectedCampaign.ID, maybeArchives)
	return archives, nil
}


func buildEngine(db *mongo.Mongo, ctx *tui.Context) (*game.Engine, error) {
	selectedCampaign, err := ResolveCampaignFromContext(ctx)
	if err != nil {
		return nil, err
	}

	archives, err := buildArchives(db, selectedCampaign)
	if err != nil {
		return nil, err
	}	

	// set engine
	engine := game.NewEngine(selectedCampaign.ID, game.NewState(game.CampaignData{
		Campaign: selectedCampaign,
		Archives: archives,
	}))

	engine.Initialize()

	// Save
	err = process.SaveGame(db, engine)
	if err != nil {
		return nil, err
	}

	return engine, nil
}

func ResolveCampaignFromContext(ctx *tui.Context) (*campaign.Campaign, error) {
	selectedCampaign := ctx.SelectedSave
	// set compaingn
	if selectedCampaign == nil {
		if ctx.SelectedRace == nil || ctx.SelectedClass == nil || ctx.SelectedLocation == nil {
			return nil, fmt.Errorf("incomplete character creation context")
		}

		player, err := characters.New(characters.Template{
			Name:  ctx.SelectedName,
			Race:  ctx.SelectedRace.GetName(),
			Class: ctx.SelectedClass.GetName(),
		})
		if err != nil {
			return nil, err
		}

		selectedCampaign = campaign.New(campaign.Template{
			Player:          player,
			CurrentLocation: ctx.SelectedLocation,
		})
	}
	return selectedCampaign, nil
}