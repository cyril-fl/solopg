package codexmenu

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/characters/archetypes/classes"
	"solopg/internal/domain/card/characters/archetypes/races"
	"solopg/internal/infrastructure/t"
	"solopg/internal/platform/form"
	"solopg/internal/platform/form/field"
)

/*
REFACTOR
MEDIUM getNpcForm && getMonsterForm
*/
func getNpcForm() *form.Model {
	var raceItems []tui.Item[string]
	for _, i := range races.List() {
		raceItems = append(raceItems, tui.NewItem(i.GetName(), "", i.GetName()))
	}

	var classItems []tui.Item[string]
	for _, i := range classes.List() {
		classItems = append(classItems, tui.NewItem(i.GetName(), "", i.GetName()))
	}

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
MEDIUM getMonsterForm && getNpcForm

faire en sorte de n'ajouter que les classe de monstre ! pas les autres...

*/
func getMonsterForm() *form.Model {
	var raceItems []tui.Item[string]
	for _, i := range races.List() {
		if !i.IsMonster() { continue }
		raceItems = append(raceItems, tui.NewItem(t.Localize(i.GetName()), "", i.GetName()))
	}

	var classItems []tui.Item[string]
	// TODO creer un iteme avec "none" et value ""
	for _, i := range classes.List() {
		/*
		REFACTOR
		MEDIUM faire en sorte d'adpter le isMonster de Race a cette sauce pour pour le sortir et l'appler comme attribues.A et ne plus laisser le choix, ou alors de faire en sorte qu'il n'y ai plusieur classe avec isMosnster true comme
		- Humanoid
		- Wyvern
		- Beast
		- Undead
		- Elemental...

		bien que s'a s'appenrente pas a une Class mais a des genre ./ race

		regarder la conv avec GPT, et resumer, class pas obigatoire en somme. is monster devrais peu etre changer
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
func getLocationForm() *form.Model {
	return form.NewForm(
		field.TextField(field.TextTemplate[string]{ID: "name", Label: "field.name", Validator: nil, Required: true}),
		field.TextField(field.TextTemplate[string]{ID: "description", Label: "field.description", Validator: nil, Required: true}),
	).Focus()
}

/* 
FIXME
HIGH Toi tu fais carrement crasher le bouzin...
*/
func getObjectForm() *form.Model {



	return form.NewForm(
		field.TextField(field.TextTemplate[string]{ID: "name", Label: "field.name", Validator: nil, Required: true}),
		field.TextField(field.TextTemplate[string]{ID: "description", Label: "field.description", Validator: nil, Required: true}),
		/*
		TODO 
		HIGH fix en meme temps que le data driven
		*/
		// field.SelectField(field.SelectTemplate[string]{ID: "category", Label: "field.category", Options: []tui.Item[string]{}, Validator: nil, Required: true}),
	).Focus()
}

func getObjectifForm() *form.Model {
	return form.NewForm(
		field.TextField(field.TextTemplate[string]{ID: "title", Label: "field.title", Validator: nil, Required: true}),
		field.TextField(field.TextTemplate[string]{ID: "description", Label: "field.description", Validator: nil, Required: true}),
	).Focus()
}
