package variety

import (
	"slices"
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/yaml"
)

// - Configuration & caching - //
type yamlConfig struct {
	Values  []Variety `yaml:"values"`
}

var fileConfigPath = config.Current.StructureFiles.Variety

var cachedConfig yamlConfig

func loadFromFile() error {
	params, err := yaml.LoadFromFile[yamlConfig](fileConfigPath)
	if err != nil {
		return err
	}

	cachedConfig = *params

	return nil
}

func (c *yamlConfig) addValue(value Variety) {
	if !slices.Contains(c.Values, value) {
		c.Values = append(c.Values, value)
	}
}


// - Variety - //
type Variety string

func List() []Variety {
	if len(cachedConfig.Values) == 0 {
		err := loadFromFile()
		if err != nil {
			return nil
		}
	}

	return cachedConfig.Values
}

func MakeDefault(defaultValue string) Variety {
	cachedConfig.addValue(Variety(defaultValue))
	return Variety(defaultValue)
}

func (r Variety) Validate() bool {
	return slices.Contains(List(), r)
}

func (r Variety) String() string {
	return string(r)
}
