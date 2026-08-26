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
	Author    string
	Message   string
	Timestamp time.Time
}

func New(entries []Entry) *Journal {
	return &Journal{
		Entries: entries,
	}
}

func (j *Journal) EnsureInitialized() *Journal {
	if j == nil {
		return New([]Entry{})
	}

	if j.Entries == nil {
		j.Entries = []Entry{}
	}

	return j
}

func (j *Journal) AddEntry(author, message string) *Entry {
	entry := &Entry{
		ID:        id.New(),
		Author:    author,
		Message:   message,
		Timestamp: time.Now().UTC(),
	}

	j.Entries = append(j.Entries, *entry)
	return entry
}

func (e *Entry) String() string {
	return e.Timestamp.Format(time.RFC3339) + " - " + e.Author + ": " + e.Message
}
