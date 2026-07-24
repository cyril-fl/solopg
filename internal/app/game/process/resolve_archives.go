package process

import (
	"solopg/internal/domain/campaign"
	"solopg/internal/domain/codex"
	"solopg/internal/domain/journal"

	"github.com/google/uuid"
)

func ResolveArchives(id uuid.UUID, archives *campaign.Archives) *campaign.Archives {
	if archives == nil {
		archives = campaign.NewArchives(campaign.ArchivesTemplate{
			CampaignID: id,
			Codex:      codex.New(codex.Template{}),
			Journal:    journal.New([]journal.Entry{}),
		})
	}

	return archives
}
