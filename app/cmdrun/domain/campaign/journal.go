package campaign

import (
	"solopg/app/cmdrun/types/id"
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

func NewJournal(entries []Entry) *Journal {
	return &Journal{
		Entries: entries,
	}
}

func (j *Journal) EnsureInitialized() {
	if j == nil {
		j = NewJournal([]Entry{})
	}

	if j.Entries == nil {
		j.Entries = []Entry{}
	}
}

// STEP 2
/*
L'entrée est enregistré dans l'élément journal "log" | "adventurer"
mais pas encore persiter en DB
*/
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
	// TODO modifier avec log
	//  Attention ca c'est les log montrée en jeux au debut de partie !
	// STEP ..
	return e.Timestamp.Format(time.RFC3339) + " - " + e.Author + ": " + e.Message
}
