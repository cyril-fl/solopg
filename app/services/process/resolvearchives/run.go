package resolvearchives

import (
	"solopg/app/domain/campaign"
	"solopg/app/domain/gameplay/codex"
	"solopg/app/services/mongo"
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
			Journal:    campaign.NewJournal([]campaign.Entry{}),
			Log:        campaign.NewJournal([]campaign.Entry{}),
		})
	}

	archives.Codex.EnsureInitialized()
	archives.Journal.EnsureInitialized()
	archives.Log.EnsureInitialized()

	return archives, nil
}
