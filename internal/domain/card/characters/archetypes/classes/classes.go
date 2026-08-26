package classes

import (
	"solopg/internal/domain/card/effects"
	"solopg/internal/infrastructure/yaml"
)

type Class struct {
	Name     string             `yaml:"name"`
	Playable bool               `yaml:"playable"`
	Bonus    []effects.Modifier `yaml:"bonus"`
	ArmorSet string             `yaml:"armor_set"`
}

// TODO optimiser class / races / ect avec des caches.

func (c Class) GetName() string {
	return c.Name
}

func (c Class) IsPlayable() bool {
	return c.Playable
}

func (c Class) GetBonus() []effects.Modifier {
	return c.Bonus
}

func (c Class) GetArmorSet() string {
	return c.ArmorSet
}

func List() []Class {
	list, err := yaml.LoadListFromFile[Class]("data/systems/archetypes/class.yaml")
	if err != nil {
		return nil
	}
	return list
}

func ListNames() []string {
	var list []string
	for _, i := range List() {
		list = append(list, i.GetName())
	}
	return list
}

func FindByName(name string) *Class {
	for _, i := range List() {
		if i.GetName() == name {
			return &i
		}
	}
	return nil
}
