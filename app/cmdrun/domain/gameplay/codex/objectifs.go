package codex

import (
	"fmt"
	"solopg/app/shared/services/i19n"
	"solopg/app/shared/services/logs"
	"solopg/app/shared/utils/transform"
	"strings"
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
		TableData: newTable(i19n.Localize(tableobjectives), entries),
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
		// i18N
		return err
	}

	o.Add(values["title"], values["description"])

	logs.SilentSuccess("codex.entry:added", map[string]any{
		"Table": tableobjectives,
		"Entry": transform.ParseJson(values),
	})

	return nil
}

// TODO refavtor tout ca en se basant sur object ect
func (o *ObjectivesTable) Summaries() []string {
	if len(o.Entries) == 0 {
		return []string{i19n.Localize("codex.entry:empty")}
	}

	content := strings.Builder{}
	for i, entry := range o.Entries {
		if entry.Title == "" {
			content.WriteString(i19n.Localize("codex.objectives.number", map[string]any{"Number": i + 1}))
		} else {
			fmt.Fprintf(&content, "%s — %s", i19n.Localize(entry.Title), i19n.Localize(entry.Description))
		}
		content.WriteString("\n\n")
	}

	return []string{content.String()}
	// summaries := make([]string, 0, len(o.Entries))
	// for i := range o.Entries {
	// 	if o.Entries[i].Title == "" {
	// 		summaries = append(summaries, i18n.Localize("codex.objectives.number", map[string]any{"Number": i + 1}))
	// 		continue
	// 	}
	// 	summaries = append(summaries, fmt.Sprintf("%s — %s", i18n.Localize(o.Entries[i].Title), i18n.Localize(o.Entries[i].Description)))
	// }

	// if len(summaries) == 0 {
	// 	summaries = append(summaries, i18n.Localize("codex.no_objectives"))
	// }

	// return summaries
}

// - Helper --//
func (o *ObjectivesTable) assertEntry(values map[string]string) error {
	var err []error

	if values["title"] == "" {

		err = append(err, logs.Error("error.required", map[string]any{
			"Subject":  "objectives",
			"Property": "property.title",
		}))
	}
	if values["description"] == "" {

		err = append(err, logs.Error("error.required", map[string]any{
			"Subject":  "objectives",
			"Property": "property.description",
		}))
	}

	if len(err) > 0 {
		return o.formatAssertErrors(err)
	}

	return nil
}

func (o *ObjectivesTable) Ensure() {
	ensureEmbeddedTable(i19n.Localize(tableobjectives), &o.TableData)
}
