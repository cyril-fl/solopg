package journal

import (
	"solopg/types/id"
	"time"
)

type Journal struct {
	Entries []Entry
}

type Entry struct {
	ID        id.ID
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
		ID:        id.New(),
		Message:   message,
		Timestamp: time.Now().UTC(),
	}

	j.Entries = append(j.Entries, *entry)
	return entry
}
