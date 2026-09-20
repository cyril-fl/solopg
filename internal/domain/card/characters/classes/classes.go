package classes

import (
	"fmt"
	"slices"
	"solopg/internal/domain/card/attributes/stats"
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/yaml"
)

// - Configuration & caching - //
type yamlConfig struct {
	Name     string           `yaml:"name"`
	Playable bool             `yaml:"playable"`
	Bonus    []stats.Modifier `yaml:"bonus"`
	ArmorSet string           `yaml:"armor_set"`
}

var filePath = config.Current.Documents.Files.Classes

var cachedConfig []yamlConfig

func loadFromFile() error {
	params, err := yaml.LoadListFromFile[yamlConfig](filePath)
	if err != nil {
		return fmt.Errorf("failed to load races from file: %w", err)
	}

	cachedConfig = params

	return nil
}

// - Class - //
type Class struct {
	name           string           `yaml:"name"`
	playable       bool             `yaml:"playable"`
	bonus          []stats.Modifier `yaml:"bonus"`
	equipementName string           `yaml:"armor_set"`
}

func new(config yamlConfig) Class {
	return Class{
		name:           config.Name,
		playable:       config.Playable,
		bonus:          config.Bonus,
		equipementName: config.ArmorSet,
	}
}

func (c Class) GetName() string {
	return c.name
}

func (c Class) IsPlayable() bool {
	return c.playable
}

func (c Class) GetBonus() []stats.Modifier {
	return c.bonus
}

func (c Class) GetEquipementName() string {
	return c.equipementName
}

func Assert(provided string) bool {
	return provided == "" || slices.Contains(ListNames(), provided)
}

// - Collection - //
func List() []Class {
	if cachedConfig == nil {
		if err := loadFromFile(); err != nil {
			fmt.Printf("Error loading classes: %v\n", err)
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

func FindByName(name string) *Class {
	for _, i := range List() {
		if i.GetName() == name {
			return &i
		}
	}
	return nil
}

func makeSet(configs []yamlConfig) []Class {
	result := make([]Class, 0, len(configs))

	for _, config := range configs {
		result = append(result, new(config))
	}

	return result
}
