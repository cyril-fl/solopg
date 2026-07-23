package game

import "solopg/internal/infrastructure/mongo"

func Start(db *mongo.Mongo) error {

	// load.ArticleFromFile("data/template/articles/potion/heal_lvl1.yaml")
	// load.CharacterFromFile("data/template/characters/monsters/slime.yaml")
	// load.CharacterFromFile("data/template/characters/npcs/blacksmith.yaml")
	// load.LocationFromFile("data/template/locations/tavern.yaml")

	// adventurerSword, err := load.GearFromFile("data/template/equipments/adventurers/sword.yaml")
	// if err != nil {
	// 	return err
	// }

	// hero, err := load.CharacterFromFile("data/template/characters/hero.yaml")
	// if err != nil {
	// 	return err
	// }

	// if hero == nil {
	// 	fmt.Println("hero is nil")
	// 	return nil
	// }

	// if adventurerSword == nil {
	// 	fmt.Println("adventurerSword is nil")
	// 	return nil
	// }

	// hero.SetEquipmentSlot(adventurerSword)
	// jsonlog.JsonifiedLog(hero)

	return nil
}

// func persistGameState(repo *mongo.Repository, hero *characters.Character, location *locations.Location) error {
// 	if repo == nil {
// 		return fmt.Errorf("database repository is nil")
// 	}
// 	if hero == nil {
// 		return fmt.Errorf("hero is nil")
// 	}
// 	if location == nil {
// 		return fmt.Errorf("location is nil")
// 	}

// 	save := mongo.NewSave(*hero, *location)
// 	return repo.Upsert(context.Background(), save)
// }

/*
Je la mettrais **pas dans `main.go`**, mais dans un **package dédié au déroulé du jeu**.

### Recommandation
Crée un package comme :

```txt
internal/game/
  engine.go
  loop.go
  state.go
```

ou plus simple :

```txt
internal/application/game/
  game.go
  loop.go
```

### Pourquoi
La boucle de gameplay n’est pas du domaine pur comme `Character` ou `Gear`.
Ce n’est pas non plus de l’infrastructure comme YAML ou Mongo.

Elle orchestre :
- le tour de jeu
- les actions du joueur
- les effets
- les transitions d’état
- les entrées/sorties

Donc elle mérite une couche dédiée.

---

## Ce que je ferais concrètement

### `cmd/solopg/main.go`
- démarre l’application
- charge les données
- instancie le moteur
- lance la boucle

### `internal/game/engine.go`
- contient `Run()`
- contient la boucle principale
- gère `Turn`, `Phase`, `Action`

### `internal/game/state.go`
- contient l’état courant :
  - personnage
  - location
  - inventaire
  - combat éventuel
  - tour courant

---

## Si ton jeu est simple
Tu peux même faire :

```txt
internal/game/game.go
```

avec une seule struct :

```go
type Game struct {
    State State
}
```

et une méthode :

```go
func (g *Game) Run() error
```

---

## Mon avis
Pour ton projet, le meilleur compromis est :

```txt
internal/game/
```

C’est plus clair que `load`, plus métier que `infrastructure`, et ça évite de mettre la logique de gameplay dans `main`.

Si tu veux, je peux te proposer **une arborescence complète de la couche game** adaptée à ton projet actuel. */
