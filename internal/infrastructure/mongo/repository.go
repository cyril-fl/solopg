package mongo

import (
	"context"
	"fmt"
	"time"

	"solopg/internal/domain/campaign"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *Mongo) collection(name string) *mongo.Collection {
	return db.instance.Collection(name)
}

func (db *Mongo) LoadSaves() ([]campaign.Campaign, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := db.collection("saves").Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to load saves: %w", err)
	}

	defer cursor.Close(ctx)

	var saves []campaign.Campaign

	if err := cursor.All(ctx, &saves); err != nil {
		return nil, fmt.Errorf("failed to decode saves: %w", err)
	}

	return saves, nil
}

func (db *Mongo) SaveCampaign(save *campaign.Campaign) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	save.UpdatedAt = time.Now().UTC()

	_, err := db.collection("saves").ReplaceOne(
		ctx,
		bson.M{"campaignId": save.ID},
		save,
		options.Replace().SetUpsert(true),
	)

	if err != nil {
		return fmt.Errorf("failed to upsert save: %w", err)
	}

	fmt.Println("Save upserted successfully:", save.ID)

	return nil
}

/* func (db *Mongo) LoadGameState() *Repository {
	return &Repository {
		saves: db.instance.Collection("saves"),
		logs: db.instance.Collection("logs"),
		codex: db.instance.Collection("codex"),
	}
}
type Repository struct {
	saves  *mongo.Collection
	logs   *mongo.Collection
	codex  *mongo.Collection
} */

/* func (db *Mongo) LoadCodex() *Repository {
	return &Repository {
		codex: db.instance.Collection("codex"),
	}
}

func (db *Mongo) LoadLogs() *Repository {
	return &Repository {
		logs: db.instance.Collection("logs"),
	}
}

func (db *Mongo) LoadSaves() *Repository {
	fmt.Println("Loading saves repository...")

	return &Repository {
		saves: db.instance.Collection("saves"),
	}
} */

/* func (repository *Repository) FindByCampaignID(ctx context.Context, campaignID uuid.UUID) (*Save, error) {
	if repository == nil || repository.collection == nil {
		return nil, fmt.Errorf("save store is not initialized")
	}

	var save Save
	err := repository.collection.FindOne(ctx, map[string]uuid.UUID{"campaignId": campaignID}).Decode(&save)
	if err != nil {
		return nil, err
	}

	return &save, nil
}
*/
