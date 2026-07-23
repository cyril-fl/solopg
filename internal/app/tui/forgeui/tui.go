package forgeui

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/card/characters"

	tea "charm.land/bubbletea/v2"
)

type Ui struct {
	program *tea.Program
}

func New() *Ui {
	return &Ui{
		program: tea.NewProgram(newModel()),
	}
}

func (ui *Ui) CreateCharacter() (*characters.Character, error) {
	res, err := ui.program.Run()
	if err != nil {
		return nil, err
	}

	finalModel, ok := res.(model)
	if !ok {
		return nil, nil
	}

	if finalModel.cancelled {
		return nil, tui.ErrCreationCancelled
	}

	return finalModel.buildCharacter()
}

/*
	load.ArticleFromFile("data/template/articles/potion/heal_lvl1.yaml")
	load.CharacterFromFile("data/template/characters/monsters/slime.yaml")
	load.CharacterFromFile("data/template/characters/npcs/blacksmith.yaml")
	load.LocationFromFile("data/template/locations/tavern.yaml")

	adventurerSword, err := load.GearFromFile("data/template/equipments/adventurers/sword.yaml")
	if err != nil {
		return err
	}

	hero, err := load.CharacterFromFile("data/template/characters/hero.yaml")
	if err != nil {
		return err
	}

	if hero == nil {
		fmt.Println("hero is nil")
		return nil
	}

	if adventurerSword == nil {
		fmt.Println("adventurerSword is nil")
		return nil
	}

	hero.SetEquipmentSlot(adventurerSword)
	jsonlog.JsonifiedLog(hero)
*/
