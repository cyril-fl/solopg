package codex

import (
	"fmt"
	"solopg/internal/domain/card/characters"
	"time"
)

type MonstersTable struct {
	*Table[MonstersEntry]
}

func (m *MonstersTable) Summaries() []string {
	m.ensureTable()
	summaries := make([]string, 0, len(m.Entries))
	for _, entry := range m.Entries {
		if entry.character == nil {
			summaries = append(summaries, "Monstre inconnu")
			continue
		}
		summaries = append(summaries, fmt.Sprintf("%s — %s", entry.character.Name, entry.character.Description))
	}
	return summaries
}

type MonstersEntry struct {
	timestamp time.Time
	character *characters.Character
}

type MonstersEntryTemplate struct {
	character *characters.Character
}

func NewMonstersTable(entries []MonstersEntry) *MonstersTable {
	return &MonstersTable{
		Table: NewTable(entries),
	}
}

func (m *MonstersTable) ensureTable() {
	ensureEmbeddedTable(&m.Table)
}

func (m *MonstersTable) AddEntry(entry MonstersEntryTemplate) {
	m.ensureTable()
	m.Entries = append(m.Entries, MonstersEntry{
		timestamp: time.Now().UTC(),
		character: entry.character,
	})
}
