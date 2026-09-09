package codexmenu

import (
	"solopg/internal/app/tui"
	"solopg/internal/app/tui/models/gameboard/codexform"
	"solopg/internal/domain/card/attributes"
	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/characters/archetypes/classes"
	"solopg/internal/domain/card/characters/archetypes/races"
	"solopg/internal/domain/card/effects"
	"solopg/internal/domain/codex"
	"solopg/internal/infrastructure/t"

	tea "charm.land/bubbletea/v2"
)

// -- Side panel menu -- //

// func (m Model) HandlesEscape() bool { return m.menu.isOpen() }

// func (m *Model) SetSize(width, height int) {
// 	m.viewport.SetWidth(width)
// 	m.viewport.SetHeight(height)
// 	// if m.FormOpen() {
// 	// 	m.form.SetWidth(width)
// 	// }
// }

func (m *Model) Update(msg tea.Msg) tea.Cmd {
	// if m.FormOpen() {
	// 	return m.updateForm(msg)
	// }
	if !m.menu.isOpen() {
		return nil
	}

	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return nil
	}
	switch key.String() {
	case tui.KeyEsc:
		// m.screen = ScreenClosed
	// case tui.KeyCtrlN:
	// m.form = codexform.New(m.formKind(), m.page.Width())
	// m.screen = ScreenForm
	case tui.KeyUp, tui.KeyDown:
		// previous := m.menu.Index()
		// m.UpdateMenu(msg)
		// if previous != m.menu.Index() {
		// m.selectActivePage()
		// }
	}
	return nil
}

type link struct {
	id   id
	name string
}

type id uint8

const (
	CodexNPCs id = iota
	CodexMonsters
	CodexLocations
	CodexObjects
	CodexObjectifs
)

// TODO refactor
// func (m *Model) addCodexEntry(result codexform.Result) error {
// 	values := result.Values
// 	var err error
// 	switch result.Kind {
// 	case codexform.NPCs, codexform.Monsters:
// 		err = addNPCToCodex(result, m.codex, values)
// 	case codexform.Locations:
// 		err = m.codex.AddLocation(values)
// 	case codexform.Objects:
// 		err = m.codex.AddObject(values)
// 	case codexform.Objectifs:
// 		err = m.codex.AddObjective(values)
// 	default:
// 		err = t.NewError("error.unknown_codex", map[string]any{"Kind": result.Kind})
// 	}

// 	if err != nil {
// 		return err
// 	}

// 	m.addJournalEntry(fmt.Sprintf("Nouvelle entrée ajoutée : %s", result.Kind))
// 	return nil
// }

// TODO bouger quand le systeme de perso aura été refondu
func addNPCToCodex(result codexform.Result, codexData *codex.Codex, values map[string]string) error {
	if result.Kind == codexform.NPCs && classes.FindByName(values["class"]) == nil {
		return t.NewError("error.unknown_class", map[string]any{"Class": values["class"]})
	}
	if races.FindByName(values["race"]) == nil {
		return t.NewError("error.unknown_race", map[string]any{"Race": values["race"]})
	}
	if result.Kind == codexform.Monsters && !races.FindByName(values["race"]).IsMonster() {
		return t.NewError("error.non_monster_race", map[string]any{"Race": values["race"]})
	}
	character, err := characters.New(characters.Template{
		Name:        values["name"],
		Description: values["description"],
		Rarity:      attributes.F,
		Class:       values["class"],
		Race:        values["race"],
		Stats:       effects.BaseStats(),
	})
	if err != nil {
		return err
	}
	if result.Kind == codexform.NPCs {
		codexData.NpcsTable.AddNPC(character)
	} else {
		codexData.MonstersTable.AddMonster(character)
	}
	return nil
}

// func (m *Model) SetMenuSize(width, height int) {
// 	m.menu.SetSize(width, height)
// }

// func (m *Model) SelectMenu(index int) { m.menu.Select(index) }

// func (m *Model) UpdateMenu(msg tea.Msg) tea.Cmd {
// 	var cmd tea.Cmd
// 	m.menu, cmd = m.menu.Update(msg)
// 	return cmd
// }

// func (m *Model) updateForm(msg tea.Msg) tea.Cmd {
// 	key, ok := msg.(tea.KeyPressMsg)
// 	if ok && key.String() == tui.KeyEsc {
// 		m.screen = ScreenPage
// 		return nil
// 	}
// 	if ok && key.String() == tui.KeyEnter {
// 		if m.form.AdvanceOnEnter() {
// 			return nil
// 		}
// 		result, err := m.form.Submit()
// 		if err != nil {
// 			m.form.SetError(err)
// 			return nil
// 		}
// 		if err := m.addCodexEntry(result); err != nil {
// 			m.form.SetError(err)
// 			return nil
// 		}
// 		m.screen = ScreenPage
// 		m.refreshPage()
// 		m.page.GotoTop()
// 		return nil
// 	}
// 	m.form.SetError(nil)
// 	return m.form.Update(msg)
// }

// func (m Model) formKind() codexform.Kind {
// 	switch m.active {
// 	case CodexNPCs:
// 		return codexform.NPCs
// 	case CodexMonsters:
// 		return codexform.Monsters
// 	case CodexLocations:
// 		return codexform.Locations
// 	case CodexObjects:
// 		return codexform.Objects
// 	default:
// 		return codexform.Objectifs
// 	}
// }
