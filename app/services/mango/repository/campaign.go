package repository

import (
	"solopg/app/domain/campaign"
	"solopg/app/services/i19n"
	"solopg/app/services/mango"
	"solopg/app/types/id"

	"go.mongodb.org/mongo-driver/bson"
)

const CampaignCollection mango.Collection = "campaigns"

type CampaignRepo struct {
	*MongoRepository
}

func NewCampaignRepo() *CampaignRepo {
	return &CampaignRepo{
		NewRepository(CampaignCollection),
	}
}

// Getters & Setters
func (r *CampaignRepo) GetDb() *mango.Mongo {
	return r.getDb()
}

func (r *CampaignRepo) SetDb(db *mango.Mongo) *CampaignRepo {
	r.setDb(db)
	return r
}

// Methods
func (r *CampaignRepo) Load() ([]campaign.Campaign, error) {
	return r.fromCollection[campaign.Campaign](nil)
}

func (r *CampaignRepo) FromCollectionByID(campaignID id.ID) (*campaign.Campaign, error) {
	campaigns, err := r.fromCollection[campaign.Campaign](bson.M{"id": campaignID})
	if err != nil {
		return nil, err
	}

	if len(campaigns) == 0 {
		return nil, i19n.NewError("error.not_found.id", map[string]any{
			"Subject": "Campaign",
			"ID":      campaignID,
		})
	}

	if len(campaigns) > 1 {
		return nil, i19n.NewError("error.unexpected:value", map[string]any{
			"Subject":  "Campaign",
			"Expected": 1,
			"Received": len(campaigns),
		})
	}

	return &campaigns[0], nil
}

func (r *CampaignRepo) Register(save *campaign.Campaign) error {
	return r.toCollection(save)
}
