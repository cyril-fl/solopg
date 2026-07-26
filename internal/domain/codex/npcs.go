package codex

import (
	"solopg/internal/domain/card/characters"
	"time"
)

type NpcsTable struct {
	*Table[NpcsEntry]
}

type NpcsEntry struct {
	timestamp time.Time
	character *characters.Character
}

type NpcsEntryTemplate struct {
	character *characters.Character
}

func NewNpcsTable(entries []NpcsEntry) *NpcsTable {
	return &NpcsTable{
		Table: NewTable(entries),
	}
}

func (n *NpcsTable) AddEntry(entry NpcsEntryTemplate) {
	n.Entries = append(n.Entries, NpcsEntry{
		timestamp: time.Now().UTC(),
		character: entry.character,
	})
}
