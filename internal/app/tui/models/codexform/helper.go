package codexform

import (
	"solopg/internal/domain/card/characters/archetypes/classes"
	"solopg/internal/domain/card/characters/archetypes/races"
	"solopg/internal/domain/card/objects"
	"solopg/internal/infrastructure/t"
	"strings"

	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

func (m *Model) nextField() {
	if len(m.fields) == 0 {
		return
	}
	m.fields[m.index].input.Blur()
	m.index = (m.index + 1) % len(m.fields)
	m.fields[m.index].input.Focus()
}

func (m *Model) previousField() {
	if len(m.fields) == 0 {
		return
	}
	m.fields[m.index].input.Blur()
	m.index--
	if m.index < 0 {
		m.index = len(m.fields) - 1
	}
	m.fields[m.index].input.Focus()
}

func (m Model) currentIsChoice() bool {
	return m.index >= 0 && m.index < len(m.fields) && len(m.fields[m.index].choices) > 0
}

func (m *Model) nextChoice() {
	if !m.currentIsChoice() {
		return
	}
	current := &m.fields[m.index]
	current.choiceIndex = (current.choiceIndex + 1) % len(current.choices)
	current.input.SetValue(current.choices[current.choiceIndex])
}

func (m *Model) previousChoice() {
	if !m.currentIsChoice() {
		return
	}
	current := &m.fields[m.index]
	current.choiceIndex--
	if current.choiceIndex < 0 {
		current.choiceIndex = len(current.choices) - 1
	}
	current.input.SetValue(current.choices[current.choiceIndex])
}

// AdvanceOnEnter advances to the next field. It returns false on the last
// field, where the caller should submit the form.
func (m *Model) AdvanceOnEnter() bool {
	if m.index >= 0 && m.index < len(m.fields)-1 {
		m.nextField()
		return true
	}
	return false
}

func (m Model) Submit() (Result, error) {
	values := make(map[string]string, len(m.fields))
	for _, item := range m.fields {
		values[item.name] = strings.TrimSpace(item.input.Value())
	}

	required := "name"
	if m.kind == Objectifs {
		required = "title"
	}
	if values[required] == "" {
		return Result{}, t.NewError(&goi18n.LocalizeConfig{
			MessageID: "error.required_field",
			TemplateData: map[string]any{
				"Field": fieldsFor(m.kind)[0][1],
			},
		})
	}
	if m.kind == NPCs {
		if values["class"] == "" || values["race"] == "" {
			return Result{}, t.NewError(&goi18n.LocalizeConfig{MessageID: "error.class_race_required"})
		}
	}
	if m.kind == Monsters && values["race"] == "" {
		return Result{}, t.NewError(&goi18n.LocalizeConfig{MessageID: "error.race_required"})
	}
	if m.kind == Objects && values["category"] == "" {
		return Result{}, t.NewError(&goi18n.LocalizeConfig{MessageID: "error.category_required"})
	}

	return Result{Kind: m.kind, Values: values}, nil
}

func (m *Model) SetError(err error) {
	if err == nil {
		m.err = ""
		return
	}
	m.err = err.Error()
}

func choicesFor(kind Kind, name string) []string {
	switch name {
	case "class":
		if kind == Monsters {
			return nil
		}
		classNames := make([]string, 0)
		for _, class := range classes.List() {
			if kind == NPCs && class.GetName() == "Monster" {
				continue
			}
			classNames = append(classNames, class.GetName())
		}
		return classNames
	case "race":
		raceNames := make([]string, 0)
		for _, race := range races.List() {
			if kind == Monsters && !race.IsMonster() {
				continue
			}
			raceNames = append(raceNames, race.GetName())
		}
		return raceNames
	case "category":
		return []string{string(objects.Weapon), string(objects.Armor), string(objects.Potion)}
	default:
		return nil
	}
}

func fieldsFor(kind Kind) [][2]string {
	switch kind {
	case NPCs:
		return [][2]string{{"name", label("field.name")}, {"description", label("field.description")}, {"class", label("field.class")}, {"race", label("field.race")}}
	case Monsters:
		return [][2]string{{"name", label("field.name")}, {"description", label("field.description")}, {"race", label("field.race")}}
	case Locations:
		return [][2]string{{"name", label("field.name")}, {"description", label("field.description")}}
	case Objects:
		return [][2]string{{"name", label("field.name")}, {"description", label("field.description")}, {"category", label("field.category")}}
	case Objectifs:
		return [][2]string{{"title", label("field.title")}, {"description", label("field.description")}}
	default:
		return nil
	}
}

func label(id string) string {
	return t.Local.MustLocalize(&goi18n.LocalizeConfig{MessageID: id})
}

func (m *Model) SetWidth(width int) {
	for i := range m.fields {
		m.fields[i].input.SetWidth(max(10, width-18))
	}
}
