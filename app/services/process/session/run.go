package runsession

import (
	"solopg/app/domain/campaign"
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/domain/card/characters/classes"
	"solopg/app/domain/card/characters/races"
	"solopg/app/domain/card/locations"
	"solopg/app/services/i19n"
	"solopg/app/services/mango"
	"solopg/app/services/mango/repository"
	"solopg/app/services/process/generatengine"
	"solopg/app/services/process/processkit"

	"solopg/app/services/process/savestate"
	"solopg/app/tui"
	"solopg/app/tui/models"
	"solopg/app/tui/view/gameboard"
	"solopg/app/tui/view/loadsave"
	"solopg/app/tui/view/onboarding/onboardarchetype"
	"solopg/app/tui/view/onboarding/onboardename"
	"solopg/app/tui/view/onboarding/onboardforgecharacter"
	"solopg/app/tui/view/onboarding/onboardlocation"
)

type process struct {
	processkit.Process
	cache

	db  *mango.Mongo
	ctx *tui.Context
	err []error
}

type cache struct {
	view *tui.Ui
}

func Process(db *mango.Mongo) *process {
	return &process{
		db: db,
	}
}

func (p *process) Run() {
	p.createTuiView()
	p.makeSelectsaveStep()
	p.makeResolveStep()
	p.runTuiView()
}

func (p *process) GetResult() {
	i19n.NewError("error:not_implemented", map[string]any{
		"Function": "GetResult",
		"Subject":  "session",
	})
}

// Methods
func (p *process) createTuiView() {
	p.cache.view = tui.New()
}

func (p *process) runTuiView() {
	p.cache.view.Run()
}

func (p *process) makeSelectsaveStep() {
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
		tui.Step{
			Submodel: loadsave.NewModel(campaigns),
			Resolve: func(ctx *tui.Context, value any) error {

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

func (p *process) makeResolveStep() {
	if p.HasErr() {
		return
	}

	p.cache.view.Add(
		tui.Step{
			Submodel: models.Resolve(),
			Resolve: func(ctx *tui.Context, value any) error {

				process := generatengine.Process(p.db, ctx)
				process.Run()

				if process.HasErr() {
					return i19n.NewError("error.unexpected:action", map[string]any{
						"Action": i19n.Localize("unexpected:action.build_engine"),
						"Error":  process.GetErr().Error(),
					})
				}

				p.cache.view.Insert(tui.Step{
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
func makeOnboardingSteps() []tui.Step {
	return []tui.Step{
		{
			Submodel: onboardename.NewModel(),
			Resolve: func(ctx *tui.Context, value any) error {
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
			Resolve: func(ctx *tui.Context, value any) error {
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
			Resolve: func(ctx *tui.Context, value any) error {
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
			Resolve: func(ctx *tui.Context, value any) error {
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
			Resolve: func(ctx *tui.Context, value any) error {
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
