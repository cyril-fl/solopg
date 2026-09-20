package races

import (
	"fmt"
	"solopg/internal/domain/card/attributes/description"
	"solopg/internal/domain/card/attributes/stats"
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/yaml"
)

// - Configuration & caching - //
type yamlConfig struct {
	Name        string           `yaml:"name"`
	Description string           `yaml:"description"`
	Playable    bool             `yaml:"playable"`
	Bonus       []stats.Modifier `yaml:"bonus"`
}

var fileConfigPath = config.Current.Documents.Files.Races

var cachedConfig []yamlConfig

func loadFromFile() error {
	params, err := yaml.LoadListFromFile[yamlConfig](fileConfigPath)
	if err != nil {
		return fmt.Errorf("failed to load races from file: %w", err)
	}

	cachedConfig = params

	return nil
}

func Has(value Template) bool {
	new := New(value)

	for _, race := range List() {
		if race.GetHash() == new.GetHash() {
			return true
		}
	}
	return false
}

func AddInConfig(value Template) error {
	if Has(value) {
		return fmt.Errorf("race already exists in config")
	}

	new := yamlConfig{
		Name:        value.Name,
		Description: value.Description,
		Playable:    value.Playable,
		Bonus:       value.Bonus,
	}

	cachedConfig = append(cachedConfig, new)

	return nil
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

func makeSet(configs []yamlConfig) []Race {
	result := make([]Race, 0, len(configs))

	for _, config := range configs {
		result = append(result, newFromYaml(config))
	}

	return result
}

func newFromYaml(config yamlConfig) Race {
	return Race{
		Description: description.New(config.Name, config.Description),
		Playable:    config.Playable,
		Bonus:       config.Bonus,
	}
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
