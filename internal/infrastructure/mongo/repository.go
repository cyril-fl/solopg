package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func Connect(ctx context.Context, uri, databaseName string) (*Repository, error) {
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	if databaseName == "" {
		databaseName = "solopg"
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		_ = client.Disconnect(ctx)
		return nil, err
	}

	return &Repository{
		client:     client,
		collection: client.Database(databaseName).Collection("saves"),
	}, nil
}

func (repository *Repository) Close(ctx context.Context) error {
	if repository == nil || repository.client == nil {
		return nil
	}

	return repository.client.Disconnect(ctx)
}

func (repository *Repository) Upsert(ctx context.Context, save *Save) error {
	if repository == nil || repository.collection == nil {
		return fmt.Errorf("save store is not initialized")
	}
	if save == nil {
		return fmt.Errorf("save is nil")
	}

	if save.CreatedAt.IsZero() {
		now := time.Now().UTC()
		save.CreatedAt = now
		save.UpdatedAt = now
	} else {
		save.UpdatedAt = time.Now().UTC()
	}

	_, err := repository.collection.ReplaceOne(
		ctx,
		map[string]uuid.UUID{"campaignId": save.CampaignID},
		save,
		options.Replace().SetUpsert(true),
	)

	return err
}

// func (repository *Repository) FindByCampaignID(ctx context.Context, campaignID uuid.UUID) (*Save, error) {
// 	if repository == nil || repository.collection == nil {
// 		return nil, fmt.Errorf("save store is not initialized")
// 	}

// 	var save Save
// 	err := repository.collection.FindOne(ctx, map[string]uuid.UUID{"campaignId": campaignID}).Decode(&save)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &save, nil
// }
