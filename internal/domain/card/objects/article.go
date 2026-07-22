package objects

import (
	"fmt"

	"solopg/internal/domain/card/attributes"
	"solopg/internal/domain/card/effects"
)

type Article struct {
	Object

	IsConsumable bool
}

type ArticleTemplate struct {
	Name        string
	Description string
	Rarity      attributes.Rarity
	Variety     attributes.Variety
	Category    Category
	Effects     []effects.Effect
	Pod         int
	Consumable  bool
}

func NewArticle(params ArticleTemplate) (*Article, error) {
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
