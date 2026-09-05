package effects

import (
	"fmt"
	"slices"
	"solopg/internal/domain/card/attributes"
	"solopg/internal/infrastructure/yaml"
)

// TODO sinspirer de domain/gameplay/dice.go pour le cache et le load du fichier YAML
type yamlStatsConfig struct {
	Names     []Stat `yaml:"names"`
	BaseStats Stats  `yaml:"baseStats"`
}

const effectConfigFilePath = "data/systems/stats.yaml"

var cacheStatsConfig *yamlStatsConfig

func loadStatsFromFile() error {
	params, err := yaml.LoadFromFile[yamlStatsConfig](effectConfigFilePath)
	if err != nil {
		return fmt.Errorf("failed to load stats from file: %w", err)
	}

	cacheStatsConfig = params

	return nil
}

type Stat string

type Stats map[Stat]int

type Modifier struct {
	Stat  Stat
	Value int
}

type Effect struct {
	attributes.Description
	Modifier Modifier
}

func ListStats() []Stat {
	if cacheStatsConfig == nil {
		err := loadStatsFromFile()
		if err != nil {
			fmt.Printf("Error loading stats: %v\n", err)
			return nil
		}
	}

	return cacheStatsConfig.Names
}

func BaseStats() Stats {
	if cacheStatsConfig == nil {
		err := loadStatsFromFile()
		if err != nil {
			fmt.Printf("Error loading stats: %v\n", err)
			return nil
		}
	}

	copy := make(Stats)

	for stat, value := range cacheStatsConfig.BaseStats {
		copy[stat] = value
	}

	return copy
}

func (s Stat) Validate() bool {
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

func (m Modifier) String() string {
	return fmt.Sprintf("%s: %d", m.Stat, m.Value)
}
