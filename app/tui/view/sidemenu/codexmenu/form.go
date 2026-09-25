package codexmenu

import (
	"solopg/app/components/form"
	"solopg/app/components/form/field"
	"solopg/app/domain/card/attributes/objectcategory"
	"solopg/app/services/t"
	"solopg/app/tui/models"
)

// - NPC Form - //
func getNpcForm() *form.Form {
	statsFields := getFormStatsFields()
	return form.
		NewForm(
			field.TextField(field.TextTemplate[string]{ID: "name", Label: "field.name", Validator: nil, Required: true}),
			field.TextField(field.TextTemplate[string]{ID: "description", Label: "field.description", Validator: nil, Required: true}),
			field.SelectField(field.SelectTemplate[string]{ID: "class", Label: "field.class", Options: getFormClassOptions(), Validator: nil, Required: true}),
			field.SelectField(field.SelectTemplate[string]{ID: "race", Label: "field.race", Options: getFormRaceOptions(), Validator: nil, Required: true}),
			field.SelectField(field.SelectTemplate[string]{ID: "encounter", Label: "field.encounter", Options: getFormOracle(), Validator: nil, Required: true}),
		).
		Add(statsFields...).
		Focus()
}

// - Beast Form - //
func getBeastiaryForm() *form.Form {
	statsFields := getFormStatsFields()
	return form.
		NewForm(
			field.TextField(field.TextTemplate[string]{ID: "name", Label: "field.name", Validator: nil, Required: true}),
			field.TextField(field.TextTemplate[string]{ID: "description", Label: "field.description", Validator: nil, Required: true}),
			field.SelectField(field.SelectTemplate[string]{ID: "encounter", Label: "field.encounter", Options: getFormOracle(), Validator: nil, Required: true}),
		).
		Add(statsFields...).
		Focus()
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
	var categoryItems []models.Item[string]

	if len(objectcategory.List()) == 0 {
		objectcategory.MakeDefault("object")
	}

	for _, i := range objectcategory.List() {
		categoryItems = append(categoryItems, models.NewItem(t.Localize(i.String()), "", i.String()))
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
