package articles

import (
	"solopg/app/cmdrun/domain/card/attributes/objectcategory"
	"solopg/app/cmdrun/domain/card/attributes/rarity"
	"solopg/app/cmdrun/domain/card/attributes/stats"
	"solopg/app/cmdrun/domain/card/attributes/variety"
	"solopg/app/cmdrun/domain/card/objects"
	"solopg/app/shared/services/logs"
	"solopg/app/shared/services/yaml"
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

		return nil, logs.Error("error.invalid:new", map[string]any{
			"Subject": "article",
			"Error":   err,
		})
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
