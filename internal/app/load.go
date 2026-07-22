package app

import (
	"context"
	"fmt"
	"solopg/internal/infrastructure/mongo"
	"time"
)

func Load() (*mongo.Mongo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	database, err := mongo.Connect(ctx, "mongodb://localhost:27017", "solopg")
	if err != nil {
		return nil, err
	}

	fmt.Println("🥭 Database connected successfully!")

	return database, nil
}
