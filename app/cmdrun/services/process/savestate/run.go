package savestate

import (
	"solopg/app/cmdrun/services/game"
	"solopg/app/shared/services/i19n"
	"solopg/app/shared/services/mango"
	"solopg/app/shared/services/mango/repository"
	"solopg/app/shared/services/process"
)

type saver struct {
	process.Process

	db     *mango.Mongo
	engine *game.Engine
}

func Process(db *mango.Mongo, engine *game.Engine) *saver {
	return &saver{
		db:     db,
		engine: engine,
	}
}

func (p *saver) Run() {
	p.saveCampaign()
	p.saveArchives()
}

func (p *saver) GetResult() {
	// i18N register
	i19n.NewError("error:not_implemented", map[string]any{
		"Function": "GetResult",
		"Subject":  "savestate",
	})
}

// Helpers
func (p *saver) saveCampaign() {
	repo := repository.NewCampaignRepo().SetDb(p.db)
	data := p.engine.ExportCampaign()

	p.SetErr(repo.Register(data))
}

func (p *saver) saveArchives() {
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
