package items

import (
	"fmt"
	"solopg/domain/attributes"
	"solopg/domain/cards"
)

type Article struct {
	Item

	IsConsumable bool
}

type NewArticleParams struct {
	Name string
	Description string
	Rarity cards.Rarity
	Variety cards.Variety
	Category Category
	Effects []attributes.Effect
	Pod int
	IsConsumable bool
}

func NewArticle(params NewArticleParams) (*Article, error) {
	// Check Card
	newItem, err := NewItem(NewItemParams{
		Name: params.Name,
		Description: params.Description,
		Rarity: params.Rarity,
		Variety: params.Variety,
		Category: params.Category,
		Effects: params.Effects,
		Pod: params.Pod,
	})
	
	if err != nil {
		return nil, err
	}

	if newItem == nil  {
		return nil, fmt.Errorf("failed to create new item for article")
	}	
	
	return &Article{
		Item: *newItem,
		IsConsumable: params.IsConsumable,
	}, nil
}

