package app

import (
	"context"
	"solopg/internal/infrastructure/mongo"
)

func CloseDatabase(repo *mongo.Mongo) {
	if repo == nil {
		return
	}
	_ = repo.Close(context.Background())
}