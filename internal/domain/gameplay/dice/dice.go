package dice

import (
	"fmt"
	"math/rand"
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/yaml"
	"strconv"
)

// - Configuration & caching - //
type yamlConfig struct {
	Name  string `yaml:"name"`
	Sides int    `yaml:"sides"`
}

var fileConfigPath = config.Current.StructureFiles.Dice

var cachedConfig []yamlConfig

func loadFromFile() error {
	paramslist, err := yaml.LoadListFromFile[yamlConfig](fileConfigPath)
	if err != nil {
		return fmt.Errorf("failed to load dice from file: %w", err)
	}

	cachedConfig = paramslist

	return nil
}

// - Dice - //
type Dice struct {
	name  string
	sides int
}

func new(config yamlConfig) Dice {
	return Dice{
		name:  config.Name,
		sides: config.Sides,
	}
}

func (d Dice) GetName() string {
	if d.name != "" {
		return d.name
	}
	return "D" + strconv.Itoa(d.sides)
}

func (d Dice) GetSides() int {
	return d.sides
}

func (d Dice) Roll() int {
	return Roll(d.sides)
}

// - Collection - //
func List() []Dice {
	if cachedConfig == nil {
		err := loadFromFile()
		if err != nil {
			fmt.Printf("Error loading dice: %v\n", err)
			return nil
		}
	}

	return makeSet(cachedConfig)
}

func makeSet(configs []yamlConfig) []Dice {
	result := make([]Dice, 0, len(configs))

	for _, config := range configs {
		result = append(result, new(config))
	}

	return result
}

// - Helpers - //
func Roll(sides int) int {
	if sides <= 0 {
		return 0
	}
	return 1 + (rand.Intn(sides))
}
