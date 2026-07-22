package locations

import (
	"fmt"
	"os"
	"solopg/domain/cards"
	"solopg/domain/stats"
	// "solopg/domain/utils"

	"gopkg.in/yaml.v3"
)		

type Location struct {
	cards.Card

	Effects []stats.Effect
}

type NewLocationParams struct {
	Name string
	Description string
	Rarity cards.Rarity
	Variety cards.Variety
	Effects []stats.Effect
}

func NewLocation(params NewLocationParams) (*Location, error) {
	// Check Card
	newCard, err := cards.NewCard(cards.NewCardParams{
		Name: params.Name,
		Description: params.Description,
		Rarity: params.Rarity,
		Variety: params.Variety,
	})
	
	if err != nil {
		return nil, err
	}

	if newCard == nil  {
		return nil, fmt.Errorf("failed to create new card for location")
	}
	
	return &Location{
		Card: *newCard,
		Effects: params.Effects,
	}, nil
}

func LoadFromFile(fileAddress string) (*Location, error) {
	data, err := os.ReadFile(fileAddress)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}
	
	var locationParams NewLocationParams
	if err := yaml.Unmarshal(data, &locationParams); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	// TODO: Ajouter une config "verbose"
	// utils.JsonifiedLog(locationParams)
	
	location, err := NewLocation(locationParams)	
	if err != nil {
		fmt.Fprintln(os.Stderr, err, "NewLocation:", fileAddress)
		return nil, err
	}

	// TODO: Ajouter une config "verbose"
	// utils.JsonifiedLog(location)

	return location, nil
}