package archetypes

import (
	"solopg/internal/domain/card/effects"
	"solopg/internal/infrastructure/yaml"
)

type Class struct {
	Name     string `yaml:"name"`
	Playable bool `yaml:"playable"`
	Bonus    []effects.Modifier `yaml:"bonus"`
	ArmorSet string `yaml:"armor_set"`
}

func ListClasses() []Class {
	class, err := yaml.LoadListFromFile[Class]("data/systems/archetypes/class.yaml")
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

func FindClassByName(name string) *Class {
	for _, class := range ListClasses() {
		if class.Name == name {
			return &class
		}
	}
	return nil
}
