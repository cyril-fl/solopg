package logger

import (
	"fmt"
	"solopg/app/types/id"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Log struct {
	ID        id.ID
	Type      Kind
	Author    string // Player Id O system
	Message   string
	Timestamp time.Time
}

type Kind string

// TODO mettre les autre enume en MAJ aussi
const (
	ERR  Kind = "ERROR"
	INFO Kind = "INFO"
)

type Template struct {
	Type    Kind
	Author  *string // Player Id O system
	Message string
}

func New(params Template) Log {
	author := "System"
	if params.Author != nil {
		author = *params.Author
	}

	return Log{
		ID:        id.New(),
		Type:      params.Type,
		Author:    author,
		Message:   params.Message,
		Timestamp: time.Now().UTC(),
	}
}

func (l *Log) String() string {
	return fmt.Sprintf("%s - [%s] - %s : %s", l.Timestamp.Format(time.RFC3339), l.Type, l.Author, l.Message)
}

// NOTE Log are imutable
func (l *Log) SetUpdatedAt(t time.Time) {
}

func (l *Log) Filter() bson.M {
	return bson.M{
		"id": l.ID,
	}
}
