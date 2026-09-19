package objects

import (
	"slices"
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/yaml"
)

// - Configuration & caching - //
type yamlConfig struct {
	Values  []Category `yaml:"values"`
}

var fileConfigPath = config.Current.StructureFiles.ObjectsCategory

var cachedConfig yamlConfig

func loadFromFile() error {
	params, err := yaml.LoadFromFile[yamlConfig](fileConfigPath)
	if err != nil {
		return err
	}

	cachedConfig = *params

	return nil
}

func (c *yamlConfig) addValue(value Category) {
	if !slices.Contains(c.Values, value) {
		c.Values = append(c.Values, value)
	}
}

// - Category - //
type Category string

func ListCategory() []Category {
	if len(cachedConfig.Values) == 0 {
		err := loadFromFile()
		if err != nil {
			return nil
		}
	}

	return cachedConfig.Values
}

/*
TODO
LOW Verrifier quelle autre concepet pourrait utiliser le "MakeDefault"
*/
func MakeDefault(defaultValue string) Category {
	ListCategory()
	cachedConfig.addValue(Category(defaultValue))
	return Category(defaultValue)
}

func AssertWithDefault(provided string) Category {
	if Category(provided).Validate() {
		return Category(provided)
	}

	return MakeDefault("object")
}

func (r Category) Validate() bool {
	return slices.Contains(ListCategory(), r)
}

func (r Category) String() string {
	return string(r)
}
