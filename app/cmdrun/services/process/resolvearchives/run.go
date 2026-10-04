package resolvearchives

import (
	"solopg/app/cmdrun/domain/campaign"
	"solopg/app/cmdrun/domain/gameplay/codex"
	"solopg/app/cmdrun/types/id"
	"solopg/app/shared/services/i19n"
	"solopg/app/shared/services/logs"
	"solopg/app/shared/services/mango"
	"solopg/app/shared/services/mango/repository"
	"solopg/app/shared/services/process"
)

type resolver struct {
	process.Process
	cache

	db         *mango.Mongo
	campaignID id.ID
}

type cache struct {
	loadedArchives *campaign.Archives
}

func Process(db *mango.Mongo, campaignID id.ID) *resolver {
	return &resolver{
		db:         db,
		campaignID: campaignID,
	}
}

func (p *resolver) Run() {
	p.loadArchives()
	p.ensureArchives()
	p.ensureArchivesInitialized()
}

func (p *resolver) GetResult() *campaign.Archives {
	return p.cache.loadedArchives
}

// Helpers
func (p *resolver) loadArchives() {
	reppo := repository.Archives()
	reppo.SetDb(p.db)

	result, err := reppo.LoadByCampaignID(p.campaignID)
	if err != nil {
		p.SetErr(err)
		return
	}

	p.cache.loadedArchives = result
}

func (p *resolver) ensureArchives() {
	if p.HasErr() {
		return
	}

	if p.cache.loadedArchives != nil {
		return
	}

	p.cache.loadedArchives = campaign.NewArchives(campaign.ArchivesTemplate{
		CampaignID: p.campaignID,
		Codex:      codex.New(),
		Journal:    campaign.NewJournal([]campaign.Entry{}),
		Log:        campaign.NewJournal([]campaign.Entry{}),
	})
}

func (p *resolver) ensureArchivesInitialized() {
	if p.HasErr() {
		return
	}

	if p.cache.loadedArchives == nil {
		err := logs.CeaseError("error:unexpected:action", map[string]any{
			"Action": "unexpected:action.load_archives",
			"Error": i19n.Localize("error.not_found", map[string]any{
				"Subject": "archives",
			}),
		})

		p.SetErr(err)
		return
	}

	p.cache.loadedArchives.EnsureInitialized()
}
