package process

import (
	"solopg/internal/domain/campaign"
	"solopg/internal/domain/codex"
	"solopg/internal/domain/journal"
	"solopg/types/id"
)

func ResolveArchives(id id.ID, archives *campaign.Archives) *campaign.Archives {
	if archives == nil {
		archives = campaign.NewArchives(campaign.ArchivesTemplate{
			CampaignID: id,
			Codex:      codex.New(codex.Template{}),
			Journal:    journal.New([]journal.Entry{}),
		})
	}

	return archives
}
