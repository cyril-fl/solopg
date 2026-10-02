package runsession

import (
	"solopg/app/cmdrun/domain/campaign"
	"solopg/app/cmdrun/domain/card/attributes/stats"
	"solopg/app/cmdrun/domain/card/characters/classes"
	"solopg/app/cmdrun/domain/card/characters/races"
	"solopg/app/cmdrun/domain/card/locations"
	"solopg/app/cmdrun/services/process/generatengine"
	cmdruntui "solopg/app/cmdrun/tui"
	"solopg/app/shared/services/i19n"
	"solopg/app/shared/services/mango"
	"solopg/app/shared/services/mango/repository"
	"solopg/app/shared/services/process"

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
	err []error
}

type cache struct {
	view *cmdruntui.Ui
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
	i19n.NewError("error:not_implemented", map[string]any{
		"Function": "GetResult",
		"Subject":  "session",
	})
}

// Methods
func (p *runner) createTuiView() {
	p.cache.view = cmdruntui.New()
}

func (p *runner) runTuiView() {
	p.cache.view.Run()
}

func (p *runner) makeSelectsaveStep() {
	if p.HasErr() {
		return
	}

	repository := repository.NewCampaignRepo().SetDb(p.db)
	campaigns, err := repository.Load()
	if err != nil { // i18N -- register
		err := i19n.NewError("error.unexpected:action", map[string]any{
			"Action":   i19n.Localize("unexpected:action.load_campaigns"),
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
					return i19n.NewError("error.unexpected:value", map[string]any{
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
					return i19n.NewError("error.unexpected:action", map[string]any{
						"Action": i19n.Localize("unexpected:action.build_engine"),
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

// Helpers
func makeOnboardingSteps() []cmdruntui.Step {
	return []cmdruntui.Step{
		{
			Submodel: onboardename.NewModel(),
			Resolve: func(ctx *cmdruntui.Context, value any) error {
				name, ok := value.(string)
				if !ok {
					// i18N -- register
					return i19n.NewError("error.unexpected:value", map[string]any{
						"Subject":  i19n.Localize("name"),
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
					// i18N -- register
					return i19n.NewError("error.unexpected:value", map[string]any{
						"Subject":  i19n.Localize("race"),
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
					// i18N -- register
					return i19n.NewError("error.unexpected:value", map[string]any{
						"Subject":  i19n.Localize("class"),
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
					// i18N -- register
					return i19n.NewError("error.unexpected:value", map[string]any{
						"Subject":  i19n.Localize("build"),
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
					// i18N -- register
					return i19n.NewError("error.unexpected:value", map[string]any{
						"Subject":  i19n.Localize("location"),
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
