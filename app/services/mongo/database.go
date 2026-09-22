package mongo

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var env = NewEnvironment()

// -- Mongo & Database -- //
type Mongo struct {
	client   *mongo.Client
	instance *mongo.Database
	err	  []error
}

func NewMongo() *Mongo {
	return &Mongo{}
}

// Dis.connection
func (db *Mongo) Connect()  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	env.ReadEnv()
	env.SetEnvURI()

	db.open(ctx)
}

func (db *Mongo) Disconnect() {
	db.close(context.Background())
}

func (db *Mongo) HasErrors() bool {
	return len(db.err) > 0
}

func (db *Mongo) GetErrors() error {
	return errors.Join(db.err...)
}

// Helper
func (db *Mongo) open(ctx context.Context) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.GetEnvURI()))
	if err != nil {
		db.setErrors(err)
		return 
	}

	if err := client.Ping(ctx, nil); err != nil {
		client.Disconnect(ctx)
		db.setErrors(err)
		return
	}

	db.client = client
	db.instance = client.Database(env.GetEnvDBName())
}

func (db *Mongo) close(ctx context.Context) error {
	if db.client == nil {
		return nil
	}

	return db.client.Disconnect(ctx)
}

func (db *Mongo) setErrors(err error) {
	db.err = append(db.err, err)
}
