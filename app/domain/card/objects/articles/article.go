package articles

import (
	"fmt"

	"solopg/app/domain/card/attributes/objectcategory"
	"solopg/app/domain/card/attributes/rarity"
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/domain/card/attributes/variety"
	"solopg/app/domain/card/objects"
	"solopg/app/services/yaml"
)

type Article struct {
	objects.Object

	IsConsumable bool
}

type Template struct {
	Name        string
	Description string
	Rarity      rarity.Rarity
	Variety     variety.Variety
	Category    objectcategory.Category
	Effects     []stats.Effect
	Pod         int
	Consumable  bool
}

func New(params Template) (*Article, error) {
	// Check Card
	newItem, err := objects.New(objects.Template{
		Name:        params.Name,
		Description: params.Description,
		Rarity:      params.Rarity,
		Variety:     params.Variety,
		Category:    params.Category,
		Effects:     params.Effects,
		Pod:         params.Pod,
	})

	if err != nil {
		return nil, err
	}

	if newItem == nil {
		return nil, fmt.Errorf("failed to create new item for article")
	}

	return &Article{
		Object:       *newItem,
		IsConsumable: params.Consumable,
	}, nil
}

func FromFile(fileAddress string) (*Article, error) {
	params, err := yaml.LoadFromFile[Template](fileAddress)
	if err != nil {
		return nil, err
	}

	return New(*params)
}
