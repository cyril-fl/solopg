package races

import (
	"solopg/internal/domain/card/attributes/stats"
	"solopg/internal/infrastructure/yaml"
)

const filePath = "data/systems/archetypes/races.yaml"

type Race struct {
	Name     string             `yaml:"name"`
	Playable bool               `yaml:"playable"`
	Beast  bool               `yaml:"beast"`
	Bonus    []stats.Modifier `yaml:"bonus"`
}

func (r Race) GetName() string {
	return r.Name
}

func (r Race) IsPlayable() bool {
	return r.Playable
}

func (r Race) IsBeast() bool {
	return r.Beast
}

func (r Race) GetBonus() []stats.Modifier {
	return r.Bonus
}

func List() []Race {
	list, err := yaml.LoadListFromFile[Race](filePath)
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

func FindByName(name string) *Race {
	for _, i := range List() {
		if i.GetName() == name {
			return &i
		}
	}
	return nil
}
