package rarity

import (
	"slices"
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/yaml"
)

// - Configuration & caching - //
type yamlConfig struct {
	Values  []Rarity `yaml:"values"`
	Default Rarity   `yaml:"default"`
}

var fileConfigPath = config.Current.StructureFiles.Rarity

var cachedConfig yamlConfig

func loadFromFile() error {
	params, err := yaml.LoadFromFile[yamlConfig](fileConfigPath)
	if err != nil {
		return err
	}

	cachedConfig = *params

	return nil
}

// - Rarity - //
type Rarity string

func List() []Rarity {
	if len(cachedConfig.Values) == 0 {
		err := loadFromFile()
		if err != nil {
			return nil
		}
	}

	return cachedConfig.Values
}

func Default() Rarity {
	if cachedConfig.Default == "" {
		if err := loadFromFile(); err != nil || !cachedConfig.Default.Validate()  {
			return ""
		}
	}

	return cachedConfig.Default
}

func (r Rarity) Validate() bool {
	return slices.Contains(List(), r)
}

func (r Rarity) String() string {
	return string(r)
}
