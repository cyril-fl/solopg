package config

import (
	"solopg/internal/infrastructure/yaml"
)

type Config struct {
	Name     string   `yaml:"name"`
	Verbose  bool     `yaml:"verbose"`
	Commands commands `yaml:"commands"`
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

var Current Config

func init() {
	Current = *Load()
}

func Load() *Config {
	config, err := yaml.LoadFromFile[Config]("data/systems/config/default.yaml")
	if err != nil {
		return nil
	}
	return config
}
