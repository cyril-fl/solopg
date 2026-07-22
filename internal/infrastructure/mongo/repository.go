package mongo

import "go.mongodb.org/mongo-driver/mongo"

func (db *Mongo) Collection(name string) *mongo.Collection {
	return db.instance.Collection(name)
}

func (db *Mongo) Saves() *mongo.Collection {
	return db.instance.Collection("saves")
}

func (db *Mongo) Logs() *mongo.Collection {
	return db.instance.Collection("logs")
}

func (db *Mongo) Codex() *mongo.Collection {
	return db.instance.Collection("codex")
}

// func (repository *Repository) Upsert(ctx context.Context, save *Save) error {
// 	if repository == nil || repository.collection == nil {
// 		return fmt.Errorf("save store is not initialized")
// 	}
// 	if save == nil {
// 		return fmt.Errorf("save is nil")
// 	}

// 	if save.CreatedAt.IsZero() {
// 		now := time.Now().UTC()
// 		save.CreatedAt = now
// 		save.UpdatedAt = now
// 	} else {
// 		save.UpdatedAt = time.Now().UTC()
// 	}

// 	_, err := repository.collection.ReplaceOne(
// 		ctx,
// 		map[string]uuid.UUID{"campaignId": save.CampaignID},
// 		save,
// 		options.Replace().SetUpsert(true),
// 	)

// 	return err
// }

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
