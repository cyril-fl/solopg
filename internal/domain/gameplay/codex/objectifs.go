package codex

import (
	"fmt"
	"solopg/internal/infrastructure/t"
	"time"
)

// - Table --//
type ObjectivesTable struct {
	*TableData[ObjectifEntry]
}
type ObjectifEntry struct {
	Timestamp   time.Time `bson:"timestamp"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
}

var tableobjectives = "codex.objectives"

func NewObjectivesTable(entries []ObjectifEntry) *ObjectivesTable {
	return &ObjectivesTable{
		TableData: newTable(t.Localize(tableobjectives), entries),
	}
}

// - Methods --//
func (o *ObjectivesTable) Add(title, description string) {
	o.Entries = append(o.Entries, ObjectifEntry{
		Timestamp:   time.Now().UTC(),
		Title:       title,
		Description: description,
	})
}

func (o *ObjectivesTable) AddFromMappedValues(values map[string]string) error {
	if err := o.assertEntry(values); err != nil {
		return err
	}

	o.Add(values["title"], values["description"])

	return nil
}

func (o *ObjectivesTable) Summaries() []string {
	summaries := make([]string, 0, len(o.Entries))
	for i := range o.Entries {
		if o.Entries[i].Title == "" {
			summaries = append(summaries, t.Localize("objective.number", map[string]any{"Number": i + 1}))
			continue
		}
		summaries = append(summaries, fmt.Sprintf("%s — %s", o.Entries[i].Title, o.Entries[i].Description))
	}

	if len(summaries) == 0 {
		summaries = append(summaries, t.Localize("codex.no_objectives"))
	}

	return summaries
}

// - Helper --//
func (o *ObjectivesTable) assertEntry(values map[string]string) error {
	var err []error

	if values["title"] == "" {
		err = append(err, fmt.Errorf("title is required"))
	}
	if values["description"] == "" {
		err = append(err, fmt.Errorf("description is required"))
	}

	if len(err) > 0 {
		return o.formatAssertErrors(err)
	}

	return nil
}

func (o *ObjectivesTable) Ensure() {
	ensureEmbeddedTable(t.Localize(tableobjectives), &o.TableData)
}
