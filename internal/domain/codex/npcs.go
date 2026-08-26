package codex

import (
	"fmt"
	"solopg/internal/domain/card/characters"
	"time"
)

type NpcsTable struct {
	*Table[NpcsEntry]
}

func (n *NpcsTable) Summaries() []string {
	n.ensureTable()
	summaries := make([]string, 0, len(n.Entries))
	for _, entry := range n.Entries {
		if entry.character == nil {
			summaries = append(summaries, "PNJ inconnu")
			continue
		}
		summaries = append(summaries, fmt.Sprintf("%s — %s", entry.character.Name, entry.character.Description))
	}
	return summaries
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

func (n *NpcsTable) ensureTable() {
	ensureEmbeddedTable(&n.Table)
}

func (n *NpcsTable) AddEntry(entry NpcsEntryTemplate) {
	n.ensureTable()
	n.Entries = append(n.Entries, NpcsEntry{
		timestamp: time.Now().UTC(),
		character: entry.character,
	})
}
