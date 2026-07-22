package main

import (
	"context"
	"fmt"
	"os"
	"time"

	// "solopg/cards/characters"
	"solopg/internal/infrastructure/mongo"
	// "solopg/cards/items/articles"
	// "solopg/cards/items/equipments"
	// "solopg/cards/utils"
	// "solopg/cards/locations"
)

func main() {
	fmt.Fprintln(os.Stdout, "Create character cards")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	repo, err := mongo.Connect(ctx, "mongodb://localhost:27017", "solopg")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer repo.Close(ctx)

	_ = repo

	//

	// articles.LoadFromFile("cards/template/articles/potion/heal_lvl1.yaml")
	// characters.LoadFromFile("cards/template/characters/monsters/slime.yaml")
	// characters.LoadFromFile("cards/template/characters/npcs/blacksmith.yaml")
	// locations.LoadFromFile("cards/template/locations/tavern.yaml")

	// adventurerSword, _:= equipments.LoadFromFile("cards/template/equipments/adventurers/sword.yaml")
	// hero, _ := characters.LoadFromFile("cards/template/characters/hero.yaml")

	// hero.SetEquipmentSlot(adventurerSword)
	// utils.JsonifiedLog(hero)

}
