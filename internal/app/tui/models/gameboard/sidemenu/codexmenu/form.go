package codexmenu

import "solopg/internal/platform/form"

func getNpcForm() *form.Model {
	return form.NewForm(
	form.NewField("Name", "", nil),
	form.NewField("Description", "", nil),
	form.NewField("Class", "", nil),
	form.NewField("Race", "", nil),
)
}

func getMonsterForm() *form.Model {
	return form.NewForm(
		form.NewField("Name", "", nil),
		form.NewField("Description", "", nil),
		form.NewField("Class", "", nil),
		form.NewField("Race", "", nil),
	)
}
func getLocationForm() *form.Model {
	return form.NewForm(
		form.NewField("Name", "", nil),
		form.NewField("Description", "", nil),
	)
}
func getObjectForm() *form.Model {
	return form.NewForm(
		form.NewField("Name", "", nil),
		form.NewField("Description", "", nil),
		form.NewField("Category", "", nil),
	)
}
func getObjectifForm() *form.Model {
	return form.NewForm(
		form.NewField("Name", "", nil),
		form.NewField("Description", "", nil),
	)
}
