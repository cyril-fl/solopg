package cmdrunrunsession

import (
	"solopg/app/cmdrun/domain/campaign"
	"solopg/app/cmdrun/domain/card/attributes/stats"
	"solopg/app/cmdrun/domain/card/characters/classes"
	"solopg/app/cmdrun/domain/card/characters/races"
	"solopg/app/cmdrun/domain/card/locations"
	"solopg/app/cmdrun/services/process/generatengine"
	cmdruntui "solopg/app/cmdrun/tui"
	"solopg/app/shared/services/logs"
	"solopg/app/shared/services/mango"
	"solopg/app/shared/services/mango/repository"
	"solopg/app/shared/services/process"
	"solopg/app/shared/services/process/generatesteplist"

	"solopg/app/cmdrun/services/process/savestate"
	"solopg/app/cmdrun/tui/models"
	"solopg/app/cmdrun/tui/view/gameboard"
	"solopg/app/cmdrun/tui/view/loadsave"
	"solopg/app/cmdrun/tui/view/onboarding/onboardarchetype"
	"solopg/app/cmdrun/tui/view/onboarding/onboardename"
	"solopg/app/cmdrun/tui/view/onboarding/onboardforgecharacter"
	"solopg/app/cmdrun/tui/view/onboarding/onboardlocation"
)

type runner struct {
	process.Process
	cache

	db  *mango.Mongo
	ctx *cmdruntui.Context
}

type cache struct {
	view *generatesteplist.Generator[cmdruntui.Context]
}

func Process(db *mango.Mongo) *runner {
	return &runner{
		db: db,
	}
}

func (p *runner) Run() {
	p.createTuiView()
	p.makeSelectsaveStep()
	p.makeResolveStep()
	p.runTuiView()
}

func (p *runner) GetResult() {
	// TODO HIGH Remplacer par un warning
	logs.NewError("error:not_implemented", map[string]any{
		"Function": "GetResult",
		"Subject":  "session",
	})
}

// Methods
func (p *runner) createTuiView() {
	p.cache.view = generatesteplist.Process[cmdruntui.Context]()
}

func (p *runner) makeSelectsaveStep() {
	if p.HasErr() {
		return
	}

	repository := repository.Campaign().SetDb(p.db)
	campaigns, err := repository.Load()
	if err != nil { // i18N -- register
		err := logs.NewError("error.unexpected:action", map[string]any{
			"Action":   "unexpected:action.load_campaigns",
			"Received": err.Error(),
		})

		p.SetErr(err)

		return
	}

	p.cache.view.Add(
		cmdruntui.Step{
			Submodel: loadsave.NewModel(campaigns),
			Resolve: func(ctx *cmdruntui.Context, value any) error {

				selected, ok := value.(*campaign.Campaign)
				if !ok { // i18N -- register
					return logs.NewError("error.unexpected:value", map[string]any{
						"Subject":  "campaign",
						"Expected": "*campaign.Campaign",
						"Received": value,
					})
				}

				ctx.SelectedSave = selected
				if ctx.SelectedSave == nil {
					onboardingSteps := makeOnboardingSteps()
					p.cache.view.Insert(onboardingSteps...)
				}

				return nil
			},
		})
}

func (p *runner) makeResolveStep() {
	if p.HasErr() {
		return
	}

	p.cache.view.Add(
		cmdruntui.Step{
			Submodel: models.Resolve(),
			Resolve: func(ctx *cmdruntui.Context, value any) error {

				process := generatengine.Process(p.db, ctx)
				process.Run()

				if process.HasErr() {
					return logs.NewError("error.unexpected:action", map[string]any{
						"Action": "unexpected:action.build_engine",
						"Error":  process.GetErr().Error(),
					})
				}

				p.cache.view.Insert(cmdruntui.Step{
					Submodel: gameboard.NewModel(gameboard.UiParams{
						Engine: process.GetResult(),
						OnSave: savestate.Save(p.db, process.GetResult()),
					}),
				})

				return nil
			},
		})
}

func (p *runner) runTuiView() {
	model := cmdruntui.NewModel(p.cache.view.GetSteps())

	p.cache.view.Run(&model)

	if p.cache.view.HasErr() {
		p.SetErr(p.cache.view.GetErr())
	}
}

// Helpers
func makeOnboardingSteps() []cmdruntui.Step {
	return []cmdruntui.Step{
		{
			Submodel: onboardename.NewModel(),
			Resolve: func(ctx *cmdruntui.Context, value any) error {
				name, ok := value.(string)
				if !ok {

					return logs.NewError("error.unexpected:value", map[string]any{
						"Subject":  "name",
						"Expected": "string",
						"Received": value,
					})
				}
				ctx.SelectedName = name
				return nil
			},
		},
		{
			Submodel: onboardarchetype.NewModel(races.List()),
			Resolve: func(ctx *cmdruntui.Context, value any) error {
				selected, ok := value.(models.Item[races.Race])
				if !ok {

					return logs.NewError("error.unexpected:value", map[string]any{
						"Subject":  "race",
						"Expected": "models.Item[races.Race]",
						"Received": value,
					})
				}
				selectedRace := selected.Value()
				ctx.SelectedRace = &selectedRace
				return nil
			},
		},
		{
			Submodel: onboardarchetype.NewModel(classes.List()),
			Resolve: func(ctx *cmdruntui.Context, value any) error {
				selected, ok := value.(models.Item[classes.Class])
				if !ok {

					return logs.NewError("error.unexpected:value", map[string]any{
						"Subject":  "class",
						"Expected": "models.Item[classes.Class]",
						"Received": value,
					})
				}
				class := selected.Value()
				ctx.SelectedClass = &class
				return nil
			},
		},
		{
			Submodel: onboardforgecharacter.NewModel(),
			Resolve: func(ctx *cmdruntui.Context, value any) error {
				build, ok := value.([]stats.Modifier)
				if !ok {

					return logs.NewError("error.unexpected:value", map[string]any{
						"Subject":  "build",
						"Expected": "[]stats.Modifier",
						"Received": value,
					})
				}
				ctx.SelectedBuild = build
				return nil
			},
		},
		{
			Submodel: onboardlocation.NewModel(),
			Resolve: func(ctx *cmdruntui.Context, value any) error {
				location, ok := value.(*locations.Location)
				if !ok {

					return logs.NewError("error.unexpected:value", map[string]any{
						"Subject":  "location",
						"Expected": "*locations.Location",
						"Received": value,
					})
				}
				ctx.SelectedLocation = location
				return nil
			},
		},
	}
}
