package classes

import (
	"fmt"
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/services/yaml"
	"solopg/config"
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

func Has(value Template) bool {
	new := New(value)

	for _, class := range List() {
		if class.GetHash() == new.GetHash() {
			return true
		}
	}
	return false
}

func AddInConfig(value Template) error {
	if Has(value) {
		return fmt.Errorf("class already exists in config")
	}

	new := yamlConfig{
		Name:     value.Name,
		Playable: value.Playable,
		Bonus:    value.Bonus,
		ArmorSet: value.EquipementName,
	}

	cachedConfig = append(cachedConfig, new)

	return nil
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

func InsertIfIsnt(name string) (*Class, error) {
	if class := FindByName(name); class != nil {
		return class, nil
	}

	if err := AddInConfig(Template{
		Name:     name,
		Playable: true,
	}); err != nil {
		return nil, err
	}

	return FindByName(name), nil
}

func makeSet(configs []yamlConfig) []Class {
	result := make([]Class, 0, len(configs))

	for _, config := range configs {
		result = append(result, newFromYaml(config))
	}

	return result
}

func newFromYaml(config yamlConfig) Class {
	return Class{
		name:           config.Name,
		playable:       config.Playable,
		bonus:          config.Bonus,
		equipementName: config.ArmorSet,
	}
}
