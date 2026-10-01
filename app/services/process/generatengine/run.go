package generatengine

import (
	"solopg/app/domain/campaign"
	"solopg/app/services/game"
	"solopg/app/services/mango"
	"solopg/app/services/process/processkit"
	"solopg/app/services/process/resolvearchives"
	"solopg/app/services/process/resolvecampaign"
	"solopg/app/services/process/savestate"
	"solopg/app/tui"
)

// Todo mettre en factory

type process struct {
	processkit.Process
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

func Process(db *mango.Mongo, ctx *tui.Context) *process {
	return &process{
		db:  db,
		ctx: ctx,
	}
}

func (m *process) Run() {
	m.setCampaign()
	m.setArchives()
	m.setEngine()
	m.ensureSaveState()
}

func (m *process) GetResult() *game.Engine {
	return m.cache.engine
}

// Methods
func (m *process) setCampaign() {
	process := resolvecampaign.Process(m.ctx)
	process.Run()

	m.cache.resolvedCampaign = process.GetResult()

	m.SetErr(process.GetErr())
}

func (m *process) setArchives() {
	if m.HasErr() {
		return
	}

	process := resolvearchives.Process(m.db, m.cache.resolvedCampaign.ID)
	process.Run()

	m.cache.loadedArchives = process.GetResult()

	m.SetErr(process.GetErr())
}

func (m *process) setEngine() {
	if m.HasErr() {
		return
	}

	m.cache.engine = game.NewEngine(m.cache.resolvedCampaign.ID, game.NewState(game.CampaignData{
		Campaign: m.cache.resolvedCampaign,
		Archives: m.cache.loadedArchives,
	}))
}

func (m *process) ensureSaveState() {
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
