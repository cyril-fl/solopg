package process

// DEPRECATED

import (
	"solopg/internal/domain/campaign"
	"solopg/internal/domain/codex"
	"solopg/internal/domain/journal"
	"solopg/types/id"
)

func ResolveArchives(id id.ID, archives *campaign.Archives) *campaign.Archives {
	if archives == nil {
		return campaign.NewArchives(campaign.ArchivesTemplate{
			CampaignID: id,
			Codex:      codex.New(),
			Journal:    journal.New([]journal.Entry{}),
		})
	}

	archives.Codex = codex.EnsureInitialized(archives.Codex)
	if archives.Journal == nil {
		archives.Journal = journal.New([]journal.Entry{})
	}

	return archives
}
