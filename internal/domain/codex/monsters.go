package codex

import (
	"solopg/internal/domain/card/characters"
	"time"
)

type MonstersTable struct {
	*Table[MonstersEntry]
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
