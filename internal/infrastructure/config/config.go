package config

import (
	"fmt"
	"solopg/internal/infrastructure/t"
	"solopg/internal/infrastructure/yaml"
)

type config struct {
	Name           string         `yaml:"name"`
	Verbose        bool           `yaml:"verbose"`
	I18n           t.Config       `yaml:"i18n"`
	Commands       commands       `yaml:"commands"`
	Documents 	structureDocument `yaml:"documents"`
}

type commands struct {
	Config  commandParams `yaml:"config"`
	Root    commandParams `yaml:"root"`
	Run     commandParams `yaml:"run"`
	Try     commandParams `yaml:"try"`
	Version commandParams `yaml:"version"`
}

type commandParams struct {
	Use     string `yaml:"use"`
	Short   string `yaml:"short"`
	Long    string `yaml:"long"`
	Example string `yaml:"example"`
	Args    []arg  `yaml:"args"`
}

type structureDocument struct {
	Folders structureFolders `yaml:"folders"`
	Files  structureFiles  `yaml:"files"`
}

type structureFolders struct {
	Characters string `yaml:"characters"`
	Locations string `yaml:"locations"`
	Oracle string `yaml:"oracle"`
}

type structureFiles struct {
	Classes string `yaml:"classes"`
	Dice   string `yaml:"dice"`
	ObjectsCategory string `yaml:"objects_category"`
	Races string `yaml:"races"`
	Rarity string `yaml:"rarity"`
	Stats  string `yaml:"stats"`
	Variety string `yaml:"variety"`
}

type arg struct {
	Name     string `yaml:"name"`
	Required bool   `yaml:"required"`
}

var Current config

func init() {
	loaded, err := Load()
	if err != nil {
		panic(err)
	}
	Current = *loaded
}

func Load() (*config, error) {
	current, err := yaml.LoadFromFile[config]("data/systems/default/config.yaml")
	if err != nil {
		return nil, err
	}

	if err := current.I18n.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return current, nil
}
