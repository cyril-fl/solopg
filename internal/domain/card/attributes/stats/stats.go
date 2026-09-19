package stats

import (
	"fmt"
	"slices"

	"solopg/internal/domain/card/attributes"
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/yaml"
)

// - Configuration & caching - //
type yamlStatsConfig struct {
	Names     []Stat `yaml:"names"`
	BaseStats Stats  `yaml:"baseStats"`
}

var fileConfigPath = config.Current.StructureFiles.Stats

var cachedConfig *yamlStatsConfig

func loadFromFile() error {
	params, err := yaml.LoadFromFile[yamlStatsConfig](fileConfigPath)
	if err != nil {
		return err
	}

	cachedConfig = params

	return nil
}

// - Stats & Effects - //
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

func List() []Stat {
	if cachedConfig == nil {
		err := loadFromFile()
		if err != nil {
			fmt.Printf("Error loading stats: %v\n", err)
			return nil
		}
	}

	return cachedConfig.Names
}

func GetBasic() Stats {
	if cachedConfig == nil {
		err := loadFromFile()
		if err != nil {
			fmt.Printf("Error loading stats: %v\n", err)
			return nil
		}
	}

	copy := make(Stats)

	for stat, value := range cachedConfig.BaseStats {
		copy[stat] = value
	}

	return copy
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

func (s Stat) Validate() bool {
	return slices.Contains(List(), s)
}


func (m Modifier) String() string {
	return fmt.Sprintf("%s: %d", m.Stat, m.Value)
}

// - Helpers - //
type YamlEffectConfig struct {
	Description YamlDescriptionConfig   `yaml:"description"`
	Modifier    YamlModifierConfig      `yaml:"modifier"`
}
 
type YamlDescriptionConfig struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type YamlModifierConfig struct {
	Stat  Stat `yaml:"stat"`
	Value int  `yaml:"value"`
}

func MakeEffectFromYamlConfigArray(yamlConfigs []YamlEffectConfig) []Effect {
	effects := make([]Effect, 0, len(yamlConfigs))

	for _, yamlConfig := range yamlConfigs {
		effect := MakeEffectFromYamlConfig(yamlConfig)
		effects = append(effects, effect)
	}

	return effects
}


func MakeEffectFromYamlConfig(yamlConfig YamlEffectConfig) Effect {
	return Effect{
		Description: attributes.Description{
			Name:        yamlConfig.Description.Name,
			Description: yamlConfig.Description.Description,
		},
		Modifier: Modifier{
			Stat:  yamlConfig.Modifier.Stat,
			Value: yamlConfig.Modifier.Value,
		},
	}
}