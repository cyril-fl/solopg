package logs

import (
	"fmt"
	"solopg/app/cmdrun/types/id"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Log struct {
	ID        id.ID
	Type      Kind
	Author    string
	Message   string
	CreatedAt time.Time
}

type Kind string

const (
	ERR  Kind = "error"
	INFO Kind = "info"
)

var Kinds = []Kind{ERR, INFO}

type Template struct {
	Type    Kind
	Author  *string // Player Id O system
	Message string
}

func New(params Template) *Log {
	author := "System"
	if params.Author != nil {
		author = *params.Author
	}

	return &Log{
		ID:        id.New(),
		Type:      params.Type,
		Author:    author,
		Message:   params.Message,
		CreatedAt: time.Now().UTC(),
	}
}

// -- Methods -- //
func (l *Log) String() string {
	return fmt.Sprintf("%s - [%s] - %s : %s", l.CreatedAt.Format(time.RFC3339), l.Type, l.Author, l.Message)
}

// Repository implementation
func (l *Log) SetUpdatedAt(t time.Time) {
	// NOTE Log are imutable
}

func (l *Log) Filter() bson.M {
	return bson.M{
		"id": l.ID,
	}
}
