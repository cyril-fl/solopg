package generatengine

import (
	"solopg/app/cmdrun/domain/campaign"
	"solopg/app/cmdrun/services/game"
	"solopg/app/cmdrun/services/process/resolvearchives"
	"solopg/app/cmdrun/services/process/resolvecampaign"
	"solopg/app/cmdrun/services/process/savestate"
	"solopg/app/cmdrun/tui"
	"solopg/app/shared/services/mango"
	"solopg/app/shared/services/process"
)

// Todo mettre en factory

type generator struct {
	process.Process
	cache

	db  *mango.Mongo
	ctx *tui.Context
	err []error
}

type cache struct {
	resolvedCampaign *campaign.Campaign
	loadedArchives   *campaign.Archives
	engine           *game.Engine

	err error
}

func Process(db *mango.Mongo, ctx *tui.Context) *generator {
	return &generator{
		db:  db,
		ctx: ctx,
	}
}

func (m *generator) Run() {
	m.setCampaign()
	m.setArchives()
	m.setEngine()
	m.ensureSaveState()
}

func (m *generator) GetResult() *game.Engine {
	return m.cache.engine
}

// Methods
func (m *generator) setCampaign() {
	process := resolvecampaign.Process(m.ctx)
	process.Run()

	m.cache.resolvedCampaign = process.GetResult()

	m.SetErr(process.GetErr())
}

func (m *generator) setArchives() {
	if m.HasErr() {
		return
	}

	process := resolvearchives.Process(m.db, m.cache.resolvedCampaign.ID)
	process.Run()

	m.cache.loadedArchives = process.GetResult()

	m.SetErr(process.GetErr())
}

func (m *generator) setEngine() {
	if m.HasErr() {
		return
	}

	m.cache.engine = game.NewEngine(m.cache.resolvedCampaign.ID, game.NewState(game.CampaignData{
		Campaign: m.cache.resolvedCampaign,
		Archives: m.cache.loadedArchives,
	}))
}

func (m *generator) ensureSaveState() {
	if m.HasErr() {
		return
	}

	if m.ctx.SelectedSave != nil {
		return
	}

	m.cache.engine.Initialize()

	process := savestate.Process(m.db, m.cache.engine)
	process.Run()

	m.SetErr(process.GetErr())
}
