package codexmenu

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/characters/archetypes/classes"
	"solopg/internal/domain/card/characters/archetypes/races"
	"solopg/internal/platform/form"
	"solopg/internal/platform/form/field"
)

/*
	REFACTOR
	MEDIUM
	getNpcForm && getMonsterForm
*/
func getNpcForm() *form.Model {
	var raceItems []tui.Item[races.Race]
	for _, i := range races.List() {
		raceItems = append(raceItems, tui.NewItem(i.GetName(), "", i))
	}

	var classItems []tui.Item[classes.Class]
	for _, i := range classes.List() {
		classItems = append(classItems, tui.NewItem(i.GetName(), "", i))
	}

	return form.NewForm(
		field.TextField(field.TextTemplate[string]{Label: "Name", Validator: nil}),
		field.TextField(field.TextTemplate[string]{Label: "Description", Validator: nil}),
		field.SelectField(field.SelectTemplate[classes.Class]{Label: "Class", Options: classItems, Validator: nil}),
		field.SelectField(field.SelectTemplate[races.Race]{Label: "Race", Options: raceItems, Validator: nil}),
	).Focus()
}

/*
	REFACTOR
	MEDIUM
	getMonsterForm && getNpcForm
*/
func getMonsterForm() *form.Model {
	var raceItems []tui.Item[races.Race]
	for _, i := range races.List() {
		raceItems = append(raceItems, tui.NewItem(i.GetName(), "", i))
	}

	var classItems []tui.Item[classes.Class]
	for _, i := range classes.List() {
		classItems = append(classItems, tui.NewItem(i.GetName(), "", i))
	}

	return form.NewForm(
		field.TextField(field.TextTemplate[string]{Label: "Name", Validator: nil}),
		field.TextField(field.TextTemplate[string]{Label: "Description", Validator: nil}),
		field.SelectField(field.SelectTemplate[classes.Class]{Label: "Class", Options: classItems, Validator: nil}),
		field.SelectField(field.SelectTemplate[races.Race]{Label: "Race", Options: raceItems, Validator: nil}),
	).Focus()
}
func getLocationForm() *form.Model {
	return form.NewForm(
		field.TextField(field.TextTemplate[string]{Label: "Name", Validator: nil}),
		field.TextField(field.TextTemplate[string]{Label: "Description", Validator: nil}),
	).Focus()
}
func getObjectForm() *form.Model {
	return form.NewForm(
		field.TextField(field.TextTemplate[string]{Label: "Name", Validator: nil}),
		field.TextField(field.TextTemplate[string]{Label: "Description", Validator: nil}),
		field.SelectField(field.SelectTemplate[string]{Label: "Category", Options: []tui.Item[string]{}, Validator: nil}),
	).Focus()
}
func getObjectifForm() *form.Model {
	return form.NewForm(
		field.TextField(field.TextTemplate[string]{Label: "Name", Validator: nil}),
		field.TextField(field.TextTemplate[string]{Label: "Description", Validator: nil}),
	).Focus()
}
