package codexmenu

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/attributes/objectcategory"
	"solopg/internal/domain/card/attributes/stats"
	"solopg/internal/domain/card/characters/classes"
	"solopg/internal/domain/card/characters/races"
	"solopg/internal/domain/gameplay/oracle"
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/t"
	"solopg/internal/platform/form"
	"solopg/internal/platform/form/field"
)

// - NPC Form - //
func getNpcForm() *form.Form {
	return form.NewForm(
		field.TextField(field.TextTemplate[string]{ID: "name", Label: "field.name", Validator: nil, Required: true}),
		field.TextField(field.TextTemplate[string]{ID: "description", Label: "field.description", Validator: nil, Required: true}),
		field.SelectField(field.SelectTemplate[string]{ID: "class", Label: "field.class", Options: getFormClassOptions(), Validator: nil, Required: true}),
		field.SelectField(field.SelectTemplate[string]{ID: "race", Label: "field.race", Options: getFormRaceOptions(), Validator: nil, Required: true}),
	).Focus()
}

func getFormRaceOptions() []tui.Item[string] {
	var raceItems []tui.Item[string]
	for _, i := range races.List() {
		raceItems = append(raceItems, tui.NewItem(i.GetName(), "", i.GetName()))
	}
	return raceItems
}

func getFormClassOptions() []tui.Item[string] {
	var classItems []tui.Item[string]
	for _, i := range classes.List() {
		classItems = append(classItems, tui.NewItem(i.GetName(), "", i.GetName()))
	}
	classItems = append(classItems, tui.NewItem("None", "", ""))
	return classItems
}

// - Beast Form - //
func getBeastiaryForm() *form.Form {
	statsFields := getFormStatsFields()
	return form.
		NewForm(
			field.TextField(field.TextTemplate[string]{ID: "name", Label: "field.name", Validator: nil, Required: true}),
			field.TextField(field.TextTemplate[string]{ID: "description", Label: "field.description", Validator: nil, Required: true}),
			field.SelectField(field.SelectTemplate[string]{ID: "encounter", Label: "field.encounter", Options: getBestiaryFormOracle(), Validator: nil, Required: true}),
		).
		Add(statsFields...).
		Focus()
}

func getBestiaryFormOracle() []tui.Item[string] {
	list, err := oracle.ListFromFolder(config.Current.Documents.Folders.Encounters)
	if err != nil {
		return nil
	}

	var oracleItems []tui.Item[string]
	for _, i := range list {
		oracleItems = append(oracleItems, tui.NewItem(t.Localize("oracles."+i), "", i))
	}

	return oracleItems
}

func getFormStatsFields() []form.Field {
	fields := []form.Field{}

	for _, i := range stats.List() {
		f := field.TextField(field.TextTemplate[string]{
			ID:           i.String(),
			Label:        t.Localize(i.String()),
			Validator:    nil,
			Required:     false,
			Defaultvalue: "",
		})
		fields = append(fields, f)
	}

	return fields
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
