package mongo

import (
	"context"
	"fmt"
	"time"

	"solopg/internal/domain/campaign"
	"solopg/types/id"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var campaignCollectionName = "campaigns"
var archivesCollectionName = "archives"

func loadFromCollection[T any](db *Mongo, collectionName string) ([]T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := db.collection(collectionName).Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to load %s: %w", collectionName, err)
	}

	defer cursor.Close(ctx)

	var data []T

	if err := cursor.All(ctx, &data); err != nil {
		return nil, fmt.Errorf("failed to decode %s: %w", collectionName, err)
	}

	return data, nil
}

func loadFromCollectionByFilter[T any](db *Mongo, collectionName string, filter bson.M) ([]T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := db.collection(collectionName).Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to load %s: %w", collectionName, err)
	}

	defer cursor.Close(ctx)

	var data []T

	if err := cursor.All(ctx, &data); err != nil {
		return nil, fmt.Errorf("failed to decode %s: %w", collectionName, err)
	}

	return data, nil
}

func saveToCollection[T Document](db *Mongo, collectionName string, data T) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data.SetUpdatedAt(time.Now().UTC())

	_, err := db.collection(collectionName).ReplaceOne(
		ctx,
		data.Filter(),
		data,
		options.Replace().SetUpsert(true),
	)

	return err
}

func (db *Mongo) collection(name string) *mongo.Collection {
	return db.instance.Collection(name)
}

func (db *Mongo) LoadCampaign() ([]campaign.Campaign, error) {
	return loadFromCollection[campaign.Campaign](db, campaignCollectionName)
}

func (db *Mongo) LoadCampaignByID(campaignID id.ID) (*campaign.Campaign, error) {
	campaigns, err := loadFromCollectionByFilter[campaign.Campaign](db, campaignCollectionName, bson.M{"id": campaignID})
	if err != nil {
		return nil, err
	}

	if len(campaigns) == 0 {
		return nil, fmt.Errorf("campaign with ID %s not found", campaignID)
	}

	return &campaigns[0], nil
}

func (db *Mongo) RegisterCampaign(save *campaign.Campaign) error {
	err := saveToCollection(db, campaignCollectionName, save)
	if err != nil {
		return fmt.Errorf("failed to save campaign: %w", err)
	}

	fmt.Println("Campaign saved successfully:", save.ID)
	return nil
}

func (db *Mongo) LoadAllArchives() ([]campaign.Archives, error) {
	return loadFromCollection[campaign.Archives](db, archivesCollectionName)
}

func (db *Mongo) LoadArchivesByCampaignID(campaignID id.ID) (*campaign.Archives, error) {
	archives, err := loadFromCollectionByFilter[campaign.Archives](db, archivesCollectionName, bson.M{"campaignId": campaignID})
	if err != nil {
		return nil, err
	}

	if len(archives) == 0 {
		return nil, nil
	}

	return &archives[0], nil
}

func (db *Mongo) RegisterArchives(archives *campaign.Archives) error {
	err := saveToCollection(db, archivesCollectionName, archives)
	if err != nil {
		return fmt.Errorf("failed to save archives: %w", err)
	}

	fmt.Println("Archives saved successfully:", archives.CampaignID)
	return nil
}
