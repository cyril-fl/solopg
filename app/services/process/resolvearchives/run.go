package resolvearchives

import (
	"solopg/app/domain/campaign"
	"solopg/app/domain/gameplay/codex"
	"solopg/app/services/i19n"
	"solopg/app/services/mango"
	"solopg/app/services/mango/repository"
	"solopg/app/services/process/processkit"
	"solopg/app/types/id"
)

type process struct {
	processkit.Process
	cache

	db         *mango.Mongo
	campaignID id.ID
}

type cache struct {
	loadedArchives *campaign.Archives
}

func Process(db *mango.Mongo, campaignID id.ID) *process {
	return &process{
		db:         db,
		campaignID: campaignID,
	}
}

func (p *process) Run() {
	p.loadArchives()
	p.ensureArchives()
	p.ensureArchivesInitialized()
}

func (p *process) GetResult() *campaign.Archives {
	return p.cache.loadedArchives
}

// Helpers
func (p *process) loadArchives() {
	reppo := repository.NewArchivesRepo()
	reppo.SetDb(p.db)

	result, err := reppo.LoadByCampaignID(p.campaignID)
	if err != nil {
		p.SetErr(err)
		return
	}

	p.cache.loadedArchives = result
}

func (p *process) ensureArchives() {
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

func (p *process) ensureArchivesInitialized() {
	if p.HasErr() {
		return
	}

	if p.cache.loadedArchives == nil {
		err := i19n.NewError("error:unexpected:action", map[string]any{
			"Action": i19n.Localize("unexpected:action.load_archives"),
			"Error": i19n.Localize("error.not_found", map[string]any{
				"Subject": i19n.Localize("archives"),
			}),
		})

		p.SetErr(err)
		return
	}

	p.cache.loadedArchives.EnsureInitialized()
}
