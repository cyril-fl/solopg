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
const classConfigFilePath = "data/systems/archetypes/class.yaml"

// TODO optimiser class / races / ect avec des caches. Prendre exemple sur domain/gameplay/dice.go pour le cache et le load du fichier YAML
/* TODO retirer les getter inutiles car les class sont exporter pour YAML ou faire un wrapper avec 
type YAMLClassConfig struct {
	Name     string             `yaml:"name"`
	Playable bool               `yaml:"playable"`
	Bonus    []effects.Modifier `yaml:"bonus"`
	ArmorSet string             `yaml:"armor_set"`
}
	
type Class struct {
	name     string             `yaml:"name"`
	playable bool               `yaml:"playable"`
	bonus    []effects.Modifier `yaml:"bonus"`
	armorSet string             `yaml:"armor_set"`
}
*/
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
	list, err := yaml.LoadListFromFile[Class](classConfigFilePath)
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
