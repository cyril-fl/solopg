package objects

import (
	"fmt"
	// "os"

	// "solopg/models/cards/objects"
	"solopg/internal/domain/card/attributes"
	"solopg/internal/domain/card/effects"
	// "solopg/utils"
	// "gopkg.in/yaml.v3"
)

type Article struct {
	Object

	IsConsumable bool
}

type NewArticleParams struct {
	Name        string
	Description string
	Rarity      attributes.Rarity
	Variety     attributes.Variety
	Category    Category
	Effects     []effects.Effect
	Pod         int
	Consumable  bool
}

func NewArticle(params NewArticleParams) (*Article, error) {
	// Check Card
	newItem, err := NewObject(NewObjectParams{
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

// func LoadFromFile(fileAddress string) (*Article, error) {
// 	data, err := os.ReadFile(fileAddress)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		return nil, err
// 	}

// 	var articleParams NewArticleParams
// 	if err := yaml.Unmarshal(data, &articleParams); err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		return nil, err
// 	}

// 	utils.JsonifiedLog(articleParams)

// 	article, err := NewArticle(articleParams)
// 	if err != nil {
// 		return nil, err
// 	}

// 	utils.JsonifiedLog(article)

// 	return article, nil
// }
