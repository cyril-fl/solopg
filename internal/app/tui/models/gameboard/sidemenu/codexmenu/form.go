package codexmenu

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/attributes/objectcategory"
	"solopg/internal/domain/card/characters/classes"
	"solopg/internal/domain/card/characters/races"
	"solopg/internal/infrastructure/t"
	"solopg/internal/platform/form"
	"solopg/internal/platform/form/field"
)

/*
REFACTOR
MEDIUM getNpcForm && getBeastForm
*/
func getNpcForm() *form.Form {
	var raceItems []tui.Item[string]
	for _, i := range races.List() {
		raceItems = append(raceItems, tui.NewItem(i.GetName(), "", i.GetName()))
	}

	var classItems []tui.Item[string]
	for _, i := range classes.List() {
		classItems = append(classItems, tui.NewItem(i.GetName(), "", i.GetName()))
	}
	// classItems = append(classItems, tui.NewItem("None", "", ""))

	//
	return form.NewForm(
		field.TextField(field.TextTemplate[string]{ID: "name", Label: "field.name", Validator: nil, Required: true}),
		field.TextField(field.TextTemplate[string]{ID: "description", Label: "field.description", Validator: nil, Required: true}),
		field.SelectField(field.SelectTemplate[string]{ID: "class", Label: "field.class", Options: classItems, Validator: nil, Required: true}),
		field.SelectField(field.SelectTemplate[string]{ID: "race", Label: "field.race", Options: raceItems, Validator: nil, Required: true}),
	).Focus()
}

/*
REFACTOR
MEDIUM getBeastForm && getNpcForm

faire en sorte de n'ajouter que les classe de betes ! pas les autres...
*/
func getBeastForm() *form.Form {
	var raceItems []tui.Item[string]
	for _, i := range races.List() {
		if !i.IsBeast() {
			continue
		}
		raceItems = append(raceItems, tui.NewItem(t.Localize(i.GetName()), "", i.GetName()))
	}

	var classItems []tui.Item[string]
	// TODO
	// HIGH implementer
	// classItems = append(classItems, tui.NewItem("None", "", ""))

	for _, i := range classes.List() {
		/*
			REFACTOR
			MEDIUM faire en sorte d'adpter le isBeast de Race a cette sauce pour pour le sortir et l'appler comme attribues.A et ne plus laisser le choix, ou alors de faire en sorte qu'il n'y ai plusieur classe avec isMosnster true comme
			- Humanoid
			- Wyvern
			- Beast
			- Undead
			- Elemental...

			bien que s'a s'appenrente pas a une Class mais a des genre ./ race

			regarder la conv avec GPT, et resumer, class pas obigatoire en somme. is beast devrais peu etre changer
		*/
		classItems = append(classItems, tui.NewItem(t.Localize(i.GetName()), "", i.GetName()))
	}

	return form.NewForm(
		field.TextField(field.TextTemplate[string]{ID: "name", Label: "field.name", Validator: nil, Required: true}),
		field.TextField(field.TextTemplate[string]{ID: "description", Label: "field.description", Validator: nil, Required: true}),
		// field.SelectField(field.SelectTemplate[string]{ID: "class", Label: "field.class", Options: classItems, Validator: nil, Required: true}),
		field.SelectField(field.SelectTemplate[string]{ID: "race", Label: "field.race", Options: raceItems, Validator: nil, Required: true}),
	).Focus()
}
func getLocationForm() *form.Form {
	return form.NewForm(
		field.TextField(field.TextTemplate[string]{ID: "name", Label: "field.name", Validator: nil, Required: true}),
		field.TextField(field.TextTemplate[string]{ID: "description", Label: "field.description", Validator: nil, Required: true}),
	).Focus()
}

func getObjectForm() *form.Form {
	var categoryItems []tui.Item[string]

	if len(objectcategory.ListCategory()) == 0 {
		objectcategory.MakeDefault("object")
	}

	for _, i := range objectcategory.ListCategory() {
		categoryItems = append(categoryItems, tui.NewItem(t.Localize(i.String()), "", i.String()))
	}

	return form.NewForm(
		field.TextField(field.TextTemplate[string]{ID: "name", Label: "field.name", Validator: nil, Required: true}),
		field.TextField(field.TextTemplate[string]{ID: "description", Label: "field.description", Validator: nil, Required: true}),
		field.SelectField(field.SelectTemplate[string]{ID: "category", Label: "field.category", Options: categoryItems, Validator: nil, Required: true}),
	).Focus()
}

func getObjectifForm() *form.Form {
	return form.NewForm(
		field.TextField(field.TextTemplate[string]{ID: "title", Label: "field.title", Validator: nil, Required: true}),
		field.TextField(field.TextTemplate[string]{ID: "description", Label: "field.description", Validator: nil, Required: true}),
	).Focus()
}
