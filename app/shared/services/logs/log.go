package logs

import (
	"fmt"
	"solopg/app/cmdrun/types/id"
	"solopg/app/shared/services/i19n"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Entry struct {
	ID        id.ID
	Type      Kind
	Author    string
	Message   string
	CreatedAt time.Time
}

type Kind string

const (
	SUCC Kind = "success"
	ERR  Kind = "error"
	INFO Kind = "info"
	WARN Kind = "warn"
)

var Kinds = []Kind{ERR, INFO}

type Template struct {
	Type    Kind
	Author  string // Player Id O system
	Message string
}

func New(params Template) *Entry {
	author := "System"
	if params.Author != "" {
		author = params.Author
	}

	return &Entry{
		ID:        id.New(),
		Type:      params.Type,
		Author:    author,
		Message:   params.Message,
		CreatedAt: time.Now().UTC(),
	}
}

// -- Methods -- //
func (l *Entry) String() string {
	return fmt.Sprintf("%s - [%s] - (%s) : \"%s\"", l.CreatedAt.Format(time.RFC3339), strings.ToUpper(string(l.Type)), l.Author, l.Message)
}

// Repository implementation
func (l *Entry) SetUpdatedAt(t time.Time) {
	// NOTE Log are imutable
}

func (l *Entry) Filter() bson.M {
	return bson.M{
		"id": l.ID,
	}
}

// -- Helpers -- //
// SystemLog creates a new log entry, registers the unlocalized message, and returns the localized one.
func SystemLog(id string, data ...map[string]any) string {
	registerFromTemplate(Template{
		Type:    ERR,
		Message: i19n.Unlocalize(id, data...),
	})

	return i19n.Localize(id, data...)
}

func registerFromTemplate(params Template) {
	log := New(params)
	cache.Register(log)
}
