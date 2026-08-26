package process

import (
	"solopg/internal/domain/campaign"
	"solopg/internal/domain/codex"
	"solopg/internal/domain/journal"
	"solopg/internal/infrastructure/mongo"
	"solopg/types/id"
)

func LoadArchivesFromDbByCampaignID(db *mongo.Mongo, campaignID id.ID) (*campaign.Archives, error) {
	maybeArchives, err := db.LoadArchivesByCampaignID(campaignID)
	if err != nil {
		return nil, err
	}

	var archives *campaign.Archives
	if maybeArchives != nil {
		archives = maybeArchives
	} else {
		archives = campaign.NewArchives(campaign.ArchivesTemplate{
			CampaignID: campaignID,
			Codex:      codex.New(),
			Journal:    journal.New([]journal.Entry{}),
			Log:        journal.New([]journal.Entry{}),
		})
	}

	archives.Codex.EnsureInitialized()
	archives.Journal.EnsureInitialized()
	archives.Log.EnsureInitialized()

	return archives, nil
}
