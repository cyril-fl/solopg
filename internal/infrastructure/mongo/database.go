package mongo

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	client   *mongo.Client
	instance *mongo.Database
}

func open(ctx context.Context, uri, databaseName string) (*Mongo, error) {
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

	return &Mongo{
		client:   client,
		instance: client.Database(databaseName),
	}, nil
}

func (db *Mongo) close(ctx context.Context) error {
	if db == nil || db.client == nil {
		return nil
	}

	return db.client.Disconnect(ctx)
}

func Connect() (*Mongo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	database, err := open(ctx, os.Getenv("MONGO_URI"), "solopg")
	if err != nil {
		return nil, err
	}

	return database, nil
}

func Disconnect(db *Mongo) {
	if db == nil {
		return
	}
	_ = db.close(context.Background())
}
