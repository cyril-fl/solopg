package archetypes

import (
	"solopg/internal/domain/card/effects"
	"solopg/internal/infrastructure/yaml"
)

type Race struct {
	Name     string             `yaml:"name"`
	Playable bool               `yaml:"playable"`
	Bonus    []effects.Modifier `yaml:"bonus"`
}

func ListRaces() []Race {
	race, err := yaml.LoadListFromFile[Race]("data/systems/archetypes/races.yaml")
	if err != nil {
		return nil
	}
	return race
}

func ListRaceNames() []string {
	var RaceList []string
	for _, race := range ListRaces() {
		RaceList = append(RaceList, race.Name)
	}
	return RaceList
}

func FindRaceByName(name string) *Race {
	for _, race := range ListRaces() {
		if race.Name == name {
			return &race
		}
	}
	return nil
}
