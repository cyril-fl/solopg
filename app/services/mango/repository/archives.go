package repository

import (
	"solopg/app/domain/campaign"
	"solopg/app/services/i18n"
	"solopg/app/services/mango"
	"solopg/app/types/id"

	"go.mongodb.org/mongo-driver/bson"
)

const ArchivesCollection mango.Collection = "archives"

type ArchivesRepo struct {
	*MongoRepository
}

func NewArchivesRepo() *ArchivesRepo {
	return &ArchivesRepo{
		NewRepository(ArchivesCollection),
	}
}

// Getters & Setters
func (r *ArchivesRepo) GetDb() *mango.Mongo {
	return r.getDb()
}

func (r *ArchivesRepo) SetDb(db *mango.Mongo) *ArchivesRepo {
	r.setDb(db)
	return r
}

// Methods
func (r *ArchivesRepo) Load() ([]campaign.Archives, error) {
	return r.fromCollection[campaign.Archives](nil)
}

func (r *ArchivesRepo) LoadByCampaignID(campaignID id.ID) (*campaign.Archives, error) {
	archives, err := r.fromCollection[campaign.Archives](bson.M{"campaignId": campaignID})
	if err != nil {
		return nil, err
	}

	if len(archives) == 0 {
		return nil, i18n.NewError("error.not_found.id", map[string]any{
			"Subject": "Archives",
			"ID":      campaignID,
		})
	}

	if len(archives) > 1 {
		return nil, i18n.NewError("error.unexpected:value", map[string]any{
			"Subject":  "Archives",
			"Expected": 1,
			"Received": len(archives),
		})
	}

	return &archives[0], nil
}

func (r *ArchivesRepo) Register(archives *campaign.Archives) error {
	return r.toCollection(archives)
}
