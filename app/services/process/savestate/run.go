package savestate

import (
	"solopg/app/services/game"
	"solopg/app/services/i19n"
	"solopg/app/services/mango"
	"solopg/app/services/mango/repository"
	"solopg/app/services/process/processkit"
)

type process struct {
	processkit.Process

	db     *mango.Mongo
	engine *game.Engine
}

func Process(db *mango.Mongo, engine *game.Engine) *process {
	return &process{
		db:     db,
		engine: engine,
	}
}

func (p *process) Run() {
	p.saveCampaign()
	p.saveArchives()
}

func (p *process) GetResult() {
	// i18N register
	i19n.NewError("error:not_implemented", map[string]any{
		"Function": "GetResult",
		"Subject":  "savestate",
	})
}

// Helpers
func (p *process) saveCampaign() {
	repo := repository.NewCampaignRepo().SetDb(p.db)
	data := p.engine.ExportCampaign()

	p.SetErr(repo.Register(data))
}

func (p *process) saveArchives() {
	repo := repository.NewArchivesRepo().SetDb(p.db)
	data := p.engine.ExportArchives()

	p.SetErr(repo.Register(data))
}

// Save returns a function that saves the game state to the database when called.
func Save(db *mango.Mongo, engine *game.Engine) func() error {
	return func() error {
		process := Process(db, engine)
		process.Run()

		return process.GetErr()
	}
}
