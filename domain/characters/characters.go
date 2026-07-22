package characters

import (
	"fmt"
	"os"
	"solopg/domain/cards"
	"solopg/domain/items"
	"solopg/domain/items/equipments"

	"solopg/domain/stats"
	// "solopg/domain/utils"

	"gopkg.in/yaml.v3"
)

type Character struct {
	cards.Card

	Class Class
	Race  Race
	Stats stats.Stats

	Equipment equipments.Equipment
	Inventory []items.Item
	Wallet items.Wallet
}

type NewCharacterParams struct {
	Name string
	Description string
	Rarity cards.Rarity
	Class Class
	Race Race
	Stats stats.Stats
	Wallet items.Wallet
	Equipment equipments.Equipment
	Inventory []items.Item
}

func NewCharacter(params NewCharacterParams) (*Character, error) {
	// Check Card
	newCard, err := cards.NewCard(cards.NewCardParams{
		Name: params.Name,
		Description: params.Description,
		Rarity: params.Rarity,
		Variety: cards.CharacterCard,
	})
	
	if err != nil {
		return nil, err
	}

	if newCard == nil  {
		return nil, fmt.Errorf("failed to create new card for character")
	}	

	if newCard.Variety != cards.CharacterCard {
		return nil, fmt.Errorf("invalid card variety for character: %s", newCard.Variety)
	}	

	// Check Class
	if err := params.Class.Validate(); err != nil {
		return nil, err
	}
	
	// Check Race
	if err := params.Race.Validate(); err != nil {
		return nil, err
	}

	return &Character{
		Card: *newCard,
		Class: params.Class,
		Race: params.Race,
		Stats: params.Stats,
		Equipment: params.Equipment,
		Inventory: params.Inventory,
		Wallet: params.Wallet,
	}, nil
}	

func LoadFromFile(fileAddress string) (*Character, error) {
	data, err := os.ReadFile(fileAddress)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	var characterParams NewCharacterParams
	if err := yaml.Unmarshal(data, &characterParams); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	// TODO: Ajouter une config "verbose"
	// utils.JsonifiedLog(characterParams)

	character, err := NewCharacter(characterParams)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	// TODO: Ajouter une config "verbose"
	// utils.JsonifiedLog(character)

	return character, nil
}

func (character *Character) SetEquipment(gear []*equipments.Gear) {
	for _, g := range gear {
		character.SetEquipmentSlot(g)
	}
}

func (character *Character) SetEquipmentSlot(gear *equipments.Gear) {
	switch gear.DestinedSlot {
	case equipments.Helmet:
		character.Equipment.Helmet = gear
	case equipments.Chestplate:
		character.Equipment.Chestplate = gear
	case equipments.Gauntlets:
		character.Equipment.Gauntlets = gear
	case equipments.Greaves:
		character.Equipment.Greaves = gear
	case equipments.Boots:
		character.Equipment.Boots = gear
	case equipments.RightHand:
		character.Equipment.RightHand = gear
	case equipments.LeftHand:
		character.Equipment.LeftHand = gear
	default:
		fmt.Printf("Invalid equipment slot: %s\n", gear.DestinedSlot)
	}
}


/* -
Point 3: stats de Character “en dur” sans redondance.
Oui je comprend qui'il puisse y avoir une incoherence. mais j'ai peur que lorsque que je chercherais a ppliquer un effet en parccouranr un tableau de Effects je n'arrive pas a ppliquer l'effet ou il faut si je ne fais pas comme c'est atuellement


- Oui, ton inquiétude est légitime, et c’est précisément pour ça que je ne te conseille pas de supprimer toute information de ciblage. La bonne séparation, c’est:

- Character garde les stats en dur, comme tu le veux.
- Effect indique quelle stat il modifie.
- Une méthode de Character fait le routage vers le bon champ.

Autrement dit, tu n’as pas besoin de stocker Stat dans chaque champ de Character pour pouvoir appliquer un effet. Le champ d’effet suffit, à condition d’avoir une fonction de conversion claire côté Character. Par exemple, un effet dit “Health -2” peut être appliqué par une méthode qui fait un switch sur l’attribut ciblé et modifie le bon champ.

Le vrai point à éviter, ce n’est pas l’absence de Stat dans Character, c’est l’absence de point unique d’application. Si tu laisses la logique de modification dispersée dans plusieurs endroits, tu vas te compliquer la vie. Si tu centralises tout dans une méthode comme ApplyEffect ou AddToStat, tu gardes un schéma rigide et simple.

Donc mon retour pragmatique est celui-ci: garde bien les stats en dur dans characters.go, mais fais porter le ciblage par stats.go via Effect.Modifier.Stat. Ensuite, applique l’effet avec une seule méthode de Character qui route vers le bon champ. C’est minimal, lisible, et ça évite la redondance sans perdre la capacité d’appliquer des effets correctement.

Si tu veux, je peux te proposer juste la forme exacte de cette méthode d’application, en restant très simple et sans réarchitecture. */

