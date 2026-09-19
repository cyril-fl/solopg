package codex

import (
	"fmt"
	"solopg/internal/domain/card/attributes/rarity"
	"solopg/internal/domain/card/attributes/variety"
	"solopg/internal/domain/card/objects"
	"solopg/internal/infrastructure/t"
	"time"
)

// - Table --//
/*
TODO
LOW Modifier Object table pour quelle prenne des card weapon / article en fonction de et creer la bonne factory pour aller avec
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
		TableData: newTable(t.Localize(tableobject), entries),
	}
}

// - Methods --//
func (o *ObjectsTable) Add(object *objects.Object) {
	o.Entries = append(o.Entries, ObjectEntry{
		Timestamp: time.Now().UTC(),
		Object:    object,
	})
}

func (o *ObjectsTable) AddFromMappedValues(values map[string]string) error {
	if err := o.assertEntry(values); err != nil {
		return err
	}

	object, err := objects.New(objects.Template{
		Name:        values["name"],
		Description: values["description"],

		Rarity:   rarity.Default(),
		Variety:  variety.AssertWithDefault(values["variety"]),
		Category: objects.AssertWithDefault(values["category"]),
		// ___
	})

	if err != nil {
		return fmt.Errorf("failed to create object from mapped values: %w", err)
	}

	o.Add(object)

	return nil
}

func (o *ObjectsTable) Summaries() []string {
	summaries := make([]string, 0, len(o.Entries))
	for _, entry := range o.Entries {
		if entry.Object == nil {
			summaries = append(summaries, t.Localize("codex.unknown_object"))
			continue
		}
		summaries = append(summaries, fmt.Sprintf("%s — %s", entry.Object.Name, entry.Object.Description))
	}

	if len(summaries) == 0 {
		summaries = append(summaries, t.Localize("codex.no_objects"))
	}

	return summaries
}

// - Helper --//
func (o *ObjectsTable) assertEntry(entry map[string]string) error {
	var err []error

	if entry["name"] == "" {
		err = append(err, fmt.Errorf("name is required"))
	}
	if entry["description"] == "" {
		err = append(err, fmt.Errorf("description is required"))
	}
	/*
		TODO
		MEDIUM quand tout sera fix ajouter ces champs dans le form
	*/
	// if entry["rarity"] == "" {
	// 	err = append(err, fmt.Errorf("rarity is required"))
	// }
	// if entry["variety"] == "" {
	// 	err = append(err, fmt.Errorf("variety is required"))
	// }
	// if entry["category"] == "" {
	// 	err = append(err, fmt.Errorf("category is required"))
	// }

	if len(err) > 0 {
		return o.formatAssertErrors(err)
	}

	return nil
}

func (o *ObjectsTable) Ensure() {
	ensureEmbeddedTable(t.Localize(tableobject), &o.TableData)
}
