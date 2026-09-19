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
// - NPC Form - //
func getNpcForm() *form.Form {
	return form.NewForm(
		field.TextField(field.TextTemplate[string]{ID: "name", Label: "field.name", Validator: nil, Required: true}),
		field.TextField(field.TextTemplate[string]{ID: "description", Label: "field.description", Validator: nil, Required: true}),
		field.SelectField(field.SelectTemplate[string]{ID: "class", Label: "field.class", Options: getNpcFormClass(), Validator: nil, Required: true}),
		field.SelectField(field.SelectTemplate[string]{ID: "race", Label: "field.race", Options: getNpcFormRace(), Validator: nil, Required: true}),
	).Focus()
}

func getNpcFormRace() []tui.Item[string] {
	var raceItems []tui.Item[string]
	for _, i := range races.List() {
		raceItems = append(raceItems, tui.NewItem(i.GetName(), "", i.GetName()))
	}
	return raceItems
}

func getNpcFormClass() []tui.Item[string] {
	var classItems []tui.Item[string]
	for _, i := range classes.List() {
		classItems = append(classItems, tui.NewItem(i.GetName(), "", i.GetName()))
	}
	classItems = append(classItems, tui.NewItem("None", "", ""))
	return classItems
}

// - Beast Form - //
func getBeastForm() *form.Form {
	var classItems []tui.Item[string]
	classItems = append(classItems, tui.NewItem("None", "", ""))

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

			est ce qu'un bete oourrais avoir un classe ? exemple un chien ne peu pas etre chasseur (quoi) mais un gobelin ?
			reponse non car tout les chien ne sont pas chasseur , la partie bestiray est plus un pokedex, et npcs l'app contacte
		*/
		classItems = append(classItems, tui.NewItem(t.Localize(i.GetName()), "", i.GetName()))
	}

	return form.NewForm(
		field.TextField(field.TextTemplate[string]{ID: "name", Label: "field.name", Validator: nil, Required: true}),
		field.TextField(field.TextTemplate[string]{ID: "description", Label: "field.description", Validator: nil, Required: true}),
		field.SelectField(field.SelectTemplate[string]{ID: "race", Label: "field.race", Options: getBeastFormRace(), Validator: nil, Required: true}),
	).Focus()
}

func getBeastFormRace() []tui.Item[string] {
	var raceItems []tui.Item[string]

	for _, i := range races.List() {
		if i.IsBeast() {
			raceItems = append(raceItems, tui.NewItem(t.Localize(i.GetName()), "", i.GetName()))
		}
	}

	return raceItems
}

// - Location Form - //
func getLocationForm() *form.Form {
	return form.NewForm(
		field.TextField(field.TextTemplate[string]{ID: "name", Label: "field.name", Validator: nil, Required: true}),
		field.TextField(field.TextTemplate[string]{ID: "description", Label: "field.description", Validator: nil, Required: true}),
	).Focus()
}

// - Object Form - //
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

// - Objectif Form - //
func getObjectifForm() *form.Form {
	return form.NewForm(
		field.TextField(field.TextTemplate[string]{ID: "title", Label: "field.title", Validator: nil, Required: true}),
		field.TextField(field.TextTemplate[string]{ID: "description", Label: "field.description", Validator: nil, Required: true}),
	).Focus()
}
