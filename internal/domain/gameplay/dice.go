package gameplay

import (
	"fmt"
	"math/rand"
	"solopg/internal/infrastructure/yaml"
	"strconv"
)

// - Configuration & caching - //
type yamlDiceConfig struct {
	Name  string `yaml:"name"`
	Sides int `yaml:"sides"`
}

const diceConfigFilePath = "data/systems/dice.yaml"

var cacheDiceConfig []yamlDiceConfig

func loadDiceFromFile() error {
	paramslist, err := yaml.LoadListFromFile[yamlDiceConfig](diceConfigFilePath)
	if err != nil {
		return fmt.Errorf("failed to load dice from file: %w", err)
	}

	cacheDiceConfig = paramslist

	return nil
}

// - Dice - //
type Dice struct {
	name string
	sides int
}

func newDice(config yamlDiceConfig) Dice {
	return Dice{
		name: config.Name,
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
	return roll(d.sides)
}

// - Dice collection - //
func ListDices() []Dice {
	if cacheDiceConfig == nil {
		err := loadDiceFromFile()
		if err != nil {
			fmt.Printf("Error loading dice: %v\n", err)
			return nil
		}
	}

	return makeDiceSet(cacheDiceConfig)
}

func makeDiceSet(dices []yamlDiceConfig) []Dice {
	result := make([]Dice, 0, len(dices))
	for _, dice := range dices {
		result = append(result, newDice(dice))
	}
	return result
}

// - Helpers - //
func roll(sides int) int {
	if sides <= 0 {
		return 0
	}
	return 1 + (rand.Intn(sides))
}