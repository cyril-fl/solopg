package codexform

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/characters/archetypes/classes"
	"solopg/internal/domain/card/characters/archetypes/races"
	"solopg/internal/domain/card/objects"
	"solopg/internal/infrastructure/t"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

type Kind string

const (
	NPCs      Kind = "NPCs"
	Monsters  Kind = "Monsters"
	Locations Kind = "Locations"
	Objects   Kind = "Objects"
	Objectifs Kind = "Objectifs"
)

type Result struct {
	Kind   Kind
	Values map[string]string
}

type field struct {
	name        string
	label       string
	input       textinput.Model
	choices     []string
	choiceIndex int
}

type Model struct {
	kind   Kind
	fields []field
	index  int
	err    string
}

func New(kind Kind, width int) Model {
	names := fieldsFor(kind)
	fields := make([]field, 0, len(names))
	for _, item := range names {
		input := textinput.New()
		input.SetWidth(max(10, width-18))
		input.Prompt = "> "
		current := field{name: item[0], label: item[1], input: input, choices: choicesFor(kind, item[0])}
		if len(current.choices) > 0 {
			current.input.SetValue(current.choices[0])
		}
		fields = append(fields, current)
	}

	model := Model{kind: kind, fields: fields}
	if len(model.fields) > 0 {
		model.fields[0].input.Focus()
	}
	return model
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
	return t.Localizer.MustLocalize(&goi18n.LocalizeConfig{MessageID: id})
}

func (m *Model) SetWidth(width int) {
	for i := range m.fields {
		m.fields[i].input.SetWidth(max(10, width-18))
	}
}

func (m *Model) Update(msg tea.Msg) tea.Cmd {
	key, ok := msg.(tea.KeyPressMsg)
	if ok {
		switch key.String() {
		case tui.KeyUp:
			if m.currentIsChoice() {
				m.previousChoice()
				return nil
			}
			m.previousField()
			return nil
		case tui.KeyDown:
			if m.currentIsChoice() {
				m.nextChoice()
				return nil
			}
			m.nextField()
			return nil
		case tui.KeyTab:
			m.nextField()
			return nil
		case tui.KeyShiftTab:
			m.previousField()
			return nil
		}
	}

	if m.index < 0 || m.index >= len(m.fields) {
		return nil
	}
	if m.currentIsChoice() {
		return nil
	}
	var cmd tea.Cmd
	m.fields[m.index].input, cmd = m.fields[m.index].input.Update(msg)
	return cmd
}

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

func (m Model) View() string {
	parts := []string{lipgloss.NewStyle().Bold(true).Render(t.Localizer.MustLocalize(&goi18n.LocalizeConfig{
		MessageID:    "codex.add",
		TemplateData: map[string]any{"Kind": string(m.kind)},
	})), ""}
	for i, item := range m.fields {
		label := item.label
		if i == m.index {
			label = "▸ " + label
		}
		parts = append(parts, label)
		if len(item.choices) > 0 {
			for choiceIndex, choice := range item.choices {
				marker := "  "
				if choiceIndex == item.choiceIndex {
					marker = "› "
				}
				parts = append(parts, marker+choice)
			}
		} else {
			parts = append(parts, item.input.View())
		}
	}
	if m.err != "" {
		parts = append(parts, "", lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(m.err))
	}
	return lipgloss.JoinVertical(lipgloss.Left, parts...)
}
