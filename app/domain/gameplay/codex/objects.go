package codex

import (
	"fmt"
	"solopg/app/domain/card/attributes/objectcategory"
	"solopg/app/domain/card/attributes/rarity"
	"solopg/app/domain/card/attributes/variety"
	"solopg/app/domain/card/objects"
	"solopg/app/services/i18n"
	"strings"
	"time"
)

// - Table --//
/*
TODO LOW Modifier Object table pour quelle prenne des card weapon / article en fonction de et creer la bonne factory pour aller avec
*/
type ObjectsTable struct {
	*TableData[ObjectEntry]
}

type ObjectEntry struct {
	Timestamp time.Time       `bson:"timestamp"`
	Object    *objects.Object `bson:"object"`
}

var tableobject = "codex.objects"

func NewObjectsTable(entries []ObjectEntry) *ObjectsTable {
	return &ObjectsTable{
		TableData: newTable(i18n.Localize(tableobject), entries),
	}
}

// - Methods --//
func (tb *ObjectsTable) Add(object *objects.Object) {
	tb.Entries = append(tb.Entries, ObjectEntry{
		Timestamp: time.Now().UTC(),
		Object:    object,
	})
}

func (tb *ObjectsTable) AddFromMappedValues(values map[string]string) error {
	if err := tb.assertEntry(values); err != nil {
		return err
	}

	object, err := objects.New(objects.Template{
		Name:        values["name"],
		Description: values["description"],

		Rarity:   rarity.Default(),
		Variety:  variety.AssertWithDefault(values["variety"]),
		Category: objectcategory.AssertWithDefault(values["category"]),
		// ___
	})

	if err != nil {
		return fmt.Errorf("failed to create object from mapped values: %w", err)
	}

	tb.Add(object)

	return nil
}

func (tb *ObjectsTable) Summaries() []string {
	if len(tb.Entries) == 0 {
		return []string{i18n.Localize("codex.entry:empty")}
	}

	content := strings.Builder{}
	for _, entry := range tb.Entries {
		if entry.Object == nil {
			content.WriteString(i18n.Localize("codex.codex.unknown:entry"))
		} else {
			fmt.Fprintf(&content, "%s — %s", i18n.Localize(entry.Object.GetName()), i18n.Localize(entry.Object.GetDescription()))
		}
		content.WriteString("\n\n")
	}

	return []string{content.String()}
}

// - Helper --//
func (tb *ObjectsTable) assertEntry(entry map[string]string) error {
	var err []error

	if entry["name"] == "" {
		err = append(err, fmt.Errorf("name is required"))
	}
	if entry["description"] == "" {
		err = append(err, fmt.Errorf("description is required"))
	}
	/*
		TODO HIGH quand tout sera fix ajouter ces champs dans le form
		if entry["rarity"] == "" {
			err = append(err, fmt.Errorf("rarity is required"))
		}
		if entry["variety"] == "" {
			err = append(err, fmt.Errorf("variety is required"))
		}
		if entry["category"] == "" {
			err = append(err, fmt.Errorf("category is required"))
		}
	*/

	if len(err) > 0 {
		return tb.formatAssertErrors(err)
	}

	return nil
}

func (tb *ObjectsTable) Ensure() {
	ensureEmbeddedTable(i18n.Localize(tableobject), &tb.TableData)
}
