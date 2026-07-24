package archetypes

import (
	"solopg/internal/domain/card/effects"
	"solopg/internal/infrastructure/yaml"
)

type Archetype struct {
	Name     string
	Playable bool
	Bonus    []effects.Modifier
}

type Race Archetype

func (a Race) String() string {
	return a.Name
}

type Class Archetype

func (a Class) String() string {
	return a.Name
}

func ListRaces() []Race {
	race, err := yaml.LoadListFromFile[Race]("data/template/systems/archetypes/races.yaml")
	if err != nil {
		return nil
	}
	return race
}

func ListClasses() []Class {
	class, err := yaml.LoadListFromFile[Class]("data/template/systems/archetypes/class.yaml")
	if err != nil {
		return nil
	}
	return class
}

func ListClassNames() []string {
	var ClassList []string
	for _, class := range ListClasses() {
		ClassList = append(ClassList, class.Name)
	}
	return ClassList
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

func FindClassByName(name string) *Class {
	for _, class := range ListClasses() {
		if class.Name == name {
			return &class
		}
	}
	return nil
}