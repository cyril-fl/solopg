package journal

import (
	"time"

	"github.com/google/uuid"
)

type Journal struct {
	Entries []Entry
}

type Entry struct {
	ID        uuid.UUID
	Message   string
	Timestamp time.Time
}

func New(entries []Entry) *Journal {
	return &Journal{
		Entries: entries,
	}
}

func (j *Journal) AddEntry(message string) *Entry {
	entry := &Entry{
		ID:        uuid.New(),
		Message:   message,
		Timestamp: time.Now().UTC(),
	}

	j.Entries = append(j.Entries, *entry)
	return entry
}
