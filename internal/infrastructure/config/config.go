package config

import (
	"fmt"
	"solopg/internal/infrastructure/i18n"
	"solopg/internal/infrastructure/yaml"
)

type config struct {
	Name     string      `yaml:"name"`
	Verbose  bool        `yaml:"verbose"`
	I18n     i18n.Config `yaml:"i18n"`
	Commands commands    `yaml:"commands"`
}

type commands struct {
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
	current, err := yaml.LoadFromFile[config]("data/systems/config/default.yaml")
	if err != nil {
		return nil, err
	}

	if err := current.I18n.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return current, nil
}
