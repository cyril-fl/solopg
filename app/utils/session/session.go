package session

import (
	"os"
	"os/exec"
	"solopg/app/domain/campaign"
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/domain/card/characters/classes"
	"solopg/app/domain/card/characters/races"
	"solopg/app/domain/card/locations"
	"solopg/app/services/game"
	"solopg/app/services/mongo"
	"solopg/app/services/process"
	"solopg/app/services/process/resolvearchives"
	"solopg/app/services/process/resolvecampaign"
	"solopg/app/services/t"
	"solopg/app/tui"
	"solopg/app/tui/models"
	"solopg/app/tui/view/gameboard"
	"solopg/app/tui/view/loadsave"
	"solopg/app/tui/view/onboarding/onboardarchetype"
	"solopg/app/tui/view/onboarding/onboardename"
	"solopg/app/tui/view/onboarding/onboardforgecharacter"
	"solopg/app/tui/view/onboarding/onboardlocation"
)

func ClearTui() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func RunSession(db *mongo.Mongo) error {
	saves, err := db.LoadCampaign()
	if err != nil {
		return err
	}

	ui := tui.New()

	return ui.
		Add(tui.Step{
			Submodel: loadsave.NewModel(saves),
			Resolve: func(ctx *tui.Context, value any) error {
				selected, ok := value.(*campaign.Campaign)
				if !ok {
					return t.NewError("error.unexpected_save_value", map[string]any{"Type": value})
				}
				ctx.SelectedSave = selected

				if ctx.SelectedSave == nil {
					onboardingSteps := onboardingSteps()
					ui.Insert(onboardingSteps...)
				}

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
						Submodel: gameboard.NewModel(gameboard.UiParams{
							Engine: engine,
							OnSave: process.NewSaveFunc(db, engine),
						}),
					})

					return nil
				},
			}).
		Run()
}

func buildEngine(db *mongo.Mongo, ctx *tui.Context) (*game.Engine, error) {
	resolvedCampaign, err := resolvecampaign.ResolveCampaignFromContext(ctx)
	if err != nil {
		return nil, err
	}

	loadedArchives, err := resolvearchives.LoadArchivesFromDbByCampaignID(db, resolvedCampaign.ID)
	if err != nil {
		return nil, err
	}

	newEngine := game.NewEngine(resolvedCampaign.ID, game.NewState(game.CampaignData{
		Campaign: resolvedCampaign,
		Archives: loadedArchives,
	}))

	isNewCampaign := ctx.SelectedSave == nil
	if isNewCampaign {
		newEngine.Initialize()

		err = process.SaveGame(db, newEngine)
		if err != nil {
			return nil, err
		}
	}

	return newEngine, nil
}

func onboardingSteps() []tui.Step {
	return []tui.Step{
		{
			Submodel: onboardename.NewModel(),
			Resolve: func(ctx *tui.Context, value any) error {
				name, ok := value.(string)
				if !ok {
					return t.NewError("error.unexpected_name_value", map[string]any{"Type": value})
				}
				ctx.SelectedName = name
				return nil
			},
		},
		{
			Submodel: onboardarchetype.NewModel(races.List()),
			Resolve: func(ctx *tui.Context, value any) error {
				race, ok := value.(string)
				if !ok {
					return t.NewError("error.unexpected_race_value", map[string]any{"Type": value})
				}
				ctx.SelectedRace = races.FindByName(race)
				return nil
			},
		},
		{
			Submodel: onboardarchetype.NewModel(classes.List()),
			Resolve: func(ctx *tui.Context, value any) error {
				class, ok := value.(string)
				if !ok {
					return t.NewError("error.unexpected_class_value", map[string]any{"Type": value})
				}
				ctx.SelectedClass = classes.FindByName(class)
				return nil
			},
		},
		{
			Submodel: onboardforgecharacter.NewModel(),
			Resolve: func(ctx *tui.Context, value any) error {
				build, ok := value.([]stats.Modifier)
				if !ok {
					return t.NewError("error.unexpected_build_value", map[string]any{"Type": value})
				}
				ctx.SelectedBuild = build
				return nil
			},
		},
		{
			Submodel: onboardlocation.NewModel(),
			Resolve: func(ctx *tui.Context, value any) error {
				location, ok := value.(*locations.Location)
				if !ok {
					return t.NewError("error.unexpected_location_value", map[string]any{"Type": value})
				}
				ctx.SelectedLocation = location
				return nil
			},
		},
	}
}
