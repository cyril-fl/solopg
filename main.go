package main

import (
	"fmt"
	"os"
	// "solopg/domain/characters"
	"solopg/domain/items/articles"
	// "solopg/domain/items/equipments"
	// "solopg/domain/utils"
	// "solopg/domain/locations"
)

func main() {
	fmt.Fprintln(os.Stdout, "Create character cards")

	articles.LoadFromFile("template/articles/potion/heal_lvl1.yaml")
	// characters.LoadFromFile("template/characters/monsters/slime.yaml")
	// characters.LoadFromFile("template/characters/npcs/blacksmith.yaml")
	// locations.LoadFromFile("template/locations/tavern.yaml")

	// adventurerSword, _:= equipments.LoadFromFile("template/equipments/adventurers/sword.yaml")
	// hero, _ := characters.LoadFromFile("template/characters/hero.yaml")

	// hero.SetEquipmentSlot(adventurerSword)
	// utils.JsonifiedLog(hero)
	
}