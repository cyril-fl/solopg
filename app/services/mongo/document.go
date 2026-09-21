package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Document interface {
	SetUpdatedAt(time.Time)
	Filter() bson.M
}
