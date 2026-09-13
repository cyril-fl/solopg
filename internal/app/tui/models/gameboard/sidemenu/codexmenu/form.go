package codexmenu

import (
	"solopg/internal/platform/form"
	"solopg/internal/platform/form/field"
)

func getNpcForm() *form.Model {
	return form.NewForm(
		field.InputField(field.InputTemplate[string]{Label: "Name", Validator: nil}),
		field.InputField(field.InputTemplate[string]{Label: "Description", Validator: nil}),
		field.SelectField(field.SelectTemplate[string]{Label: "Class", Options: []string{}, Validator: nil}),
		field.SelectField(field.SelectTemplate[string]{Label: "Race", Options: []string{}, Validator: nil}),
	)
}

func getMonsterForm() *form.Model {
	return form.NewForm(
		field.InputField(field.InputTemplate[string]{Label: "Name",  Validator: nil}),
		field.InputField(field.InputTemplate[string]{Label: "Description",  Validator: nil}),
		field.SelectField(field.SelectTemplate[string]{Label: "Class",  Options: []string{}, Validator: nil}),
		field.SelectField(field.SelectTemplate[string]{Label: "Race",  Options: []string{}, Validator: nil}),
	)
}
func getLocationForm() *form.Model {
	return form.NewForm(
		field.InputField(field.InputTemplate[string]{Label: "Name",  Validator: nil}),
		field.InputField(field.InputTemplate[string]{Label: "Description",  Validator: nil}),
	)
}
func getObjectForm() *form.Model {
	return form.NewForm(
		field.InputField(field.InputTemplate[string]{Label: "Name",  Validator: nil}),
		field.InputField(field.InputTemplate[string]{Label: "Description",  Validator: nil}),
		field.SelectField(field.SelectTemplate[string]{Label: "Category",  Options: []string{}, Validator: nil}),
	)
}
func getObjectifForm() *form.Model {
	return form.NewForm(
		field.InputField(field.InputTemplate[string]{Label: "Name",  Validator: nil}),
		field.InputField(field.InputTemplate[string]{Label: "Description",  Validator: nil}),
	)
}
