package yaml

import "solopg/internal/domain/card/objects"

func ArticleFromFile(fileAddress string) (*objects.Article, error) {
	var articleParams objects.ArticleTemplate
	if err := loadYAMLFromFile(fileAddress, &articleParams); err != nil {
		return nil, err
	}

	return objects.NewArticle(articleParams)
}
