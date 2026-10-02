package mango

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var env = NewEnvironment()

// -- Mongo & Database -- //
type Mongo struct {
	client   *mongo.Client
	instance *mongo.Database
	err      []error
}

func New() *Mongo {
	return &Mongo{}
}

// - Methods -
// Dis.connection
func (db *Mongo) Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	env.ReadEnv()
	env.SetEnvURI()

	db.open(ctx)
}

func (db *Mongo) Disconnect() {
	db.close(context.Background())
}

// Error
func (db *Mongo) HasErrors() bool {
	return len(db.err) > 0
}

func (db *Mongo) GetErrors() error {
	return errors.Join(db.err...)
}

// Collection
func (db *Mongo) GetCollection(collection Collection) *mongo.Collection {
	return db.instance.Collection(collection.String())
}

// Helper
func (db *Mongo) open(ctx context.Context) {
	client, err := mongo.
		Connect(ctx, options.Client().
			ApplyURI(env.GetEnvURI()).
			SetDirect(true))

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

	fmt.Println("🥭 DB MongoDB connected successfully.")
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

// -- Collection -- //
type Collection string

func (c Collection) String() string {
	return string(c)
}

func Init() (db *Mongo, err error) {
	db = New()
	db.Connect()
	if db.HasErrors() {
		return nil, db.GetErrors()
	}
	return db, nil
}
