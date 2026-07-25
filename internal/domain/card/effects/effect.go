package effects

import (
	"fmt"
	"slices"
	"solopg/internal/domain/card/attributes"
	"solopg/internal/infrastructure/yaml"
	"solopg/internal/platform/jsonlog"
)

const fileAddress = "data/systems/stats.yaml"

var config *statsConfig

type Stat string

type Stats map[Stat]int

type statsConfig struct {
	Names     []Stat `yaml:"names"`
	BaseStats Stats  `yaml:"baseStats"`
}

type Modifier struct {
	Stat  Stat
	Value int
}

type Effect struct {
	attributes.Description
	Modifier Modifier
}

func loadStatsFromFile() error {
	params, err := yaml.LoadFromFile[statsConfig](fileAddress)
	if err != nil {
		return fmt.Errorf("failed to load stats from file: %w", err)
	}

	config = params

	return nil
}

func ListStats() []Stat {
	if config == nil {
		err := loadStatsFromFile()
		if err != nil {
			fmt.Printf("Error loading stats: %v\n", err)
			return nil
		}
	}

	return config.Names
}

func BaseStats() Stats {
	if config == nil {
		err := loadStatsFromFile()
		if err != nil {
			fmt.Printf("Error loading stats: %v\n", err)
			return nil
		}
	}

	copy := make(Stats)

	for stat, value := range config.BaseStats {
		copy[stat] = value
	}

	return copy
}

func (s Stat) Validate() bool {

	jsonlog.JsonifiedLog(s)

	jsonlog.JsonifiedLog(config)
	return slices.Contains(ListStats(), s)
}

func (s *Stats) ApplyModifier(mod Modifier) {
	if s == nil {
		return
	}

	isValidStat := mod.Stat.Validate()
	if !isValidStat {
		fmt.Printf("Invalid stat: %s\n", mod.Stat)
		return
	}

	(*s)[mod.Stat] += mod.Value
}

func (s *Stats) ApplyModifiers(mods []Modifier) {
	for _, mod := range mods {
		s.ApplyModifier(mod)
	}
}
