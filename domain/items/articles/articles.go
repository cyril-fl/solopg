package articles

import (
	"fmt"
	"os"
	"solopg/domain/cards"
	"solopg/domain/items"
	"solopg/domain/stats"
	"solopg/domain/utils"

	"gopkg.in/yaml.v3"
)

type Article struct {
	items.Item

	IsConsumable bool
}

type NewArticleParams struct {
	Name string
	Description string
	Rarity cards.Rarity
	Variety cards.Variety
	Category items.Category
	Effects []stats.Effect
	Pod int
	Consumable bool 
}

func NewArticle(params NewArticleParams) (*Article, error) {
	// Check Card
	newItem, err := items.NewItem(items.NewItemParams{
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
		IsConsumable: params.Consumable,
	}, nil
}

func LoadFromFile(fileAddress string) (*Article, error) {
	data, err := os.ReadFile(fileAddress)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}
	
	var articleParams NewArticleParams
	if err := yaml.Unmarshal(data, &articleParams); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	utils.JsonifiedLog(articleParams)

	article, err := NewArticle(articleParams)	
	if err != nil {
		return nil, err
	}

	utils.JsonifiedLog(article)

	return article, nil
}