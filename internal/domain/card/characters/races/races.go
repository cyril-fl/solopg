package races

import (
	"fmt"
	"slices"
	"solopg/internal/domain/card/attributes/stats"
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/yaml"
)

// - Configuration & caching - //
type yamlConfig struct {
	Name     string `yaml:"name"`
	Playable bool   `yaml:"playable"`
	Beast    bool   `yaml:"beast"`
	// TODO verrifeir pour location si ce serait pas mieux quun yaml bricoler
	Bonus []stats.Modifier `yaml:"bonus"`
}

var filePath = config.Current.Documents.Files.Races

var cachedConfig []yamlConfig

func loadFromFile() error {
	params, err := yaml.LoadListFromFile[yamlConfig](filePath)
	if err != nil {
		return fmt.Errorf("failed to load races from file: %w", err)
	}

	cachedConfig = params

	return nil
}

// - Race - //
type Race struct {
	name     string
	playable bool
	beast    bool
	bonus    []stats.Modifier
}

func new(config yamlConfig) Race {
	return Race{
		name:     config.Name,
		playable: config.Playable,
		beast:    config.Beast,
		bonus:    config.Bonus,
	}
}

func (r Race) GetName() string {
	return r.name
}

func (r Race) IsPlayable() bool {
	return r.playable
}

func (r Race) IsBeast() bool {
	return r.beast
}

func (r Race) GetBonus() []stats.Modifier {
	return r.bonus
}

func Assert(provided string) bool {
	return provided != "" || slices.Contains(ListNames(), provided)
}

// - Collection - //
func List() []Race {
	if cachedConfig == nil {
		if err := loadFromFile(); err != nil {
			fmt.Printf("Error loading races: %v\n", err)
			return nil
		}
	}
	return makeSet(cachedConfig)
}

func ListNames() []string {
	var list []string
	for _, i := range List() {
		list = append(list, i.GetName())
	}
	return list
}

func FindByName(name string) *Race {
	for _, i := range List() {
		if i.GetName() == name {
			return &i
		}
	}
	return nil
}

func makeSet(configs []yamlConfig) []Race {
	result := make([]Race, 0, len(configs))

	for _, config := range configs {
		result = append(result, new(config))
	}

	return result
}
