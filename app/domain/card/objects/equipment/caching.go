package equipment

import (
	"errors"
	"fmt"
	"solopg/app/domain/card/attributes/objectcategory"
	"solopg/app/domain/card/attributes/potency"
	"solopg/app/domain/card/attributes/rarity"
	"solopg/app/domain/card/attributes/slot"
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/domain/card/attributes/variety"
	"solopg/app/services/yaml"
)

// - Configuration & caching - //
type yamlConfig struct {
	Name        string                  `yaml:"name"`
	Description string                  `yaml:"description"`
	Set         string                  `yaml:"set"`
	Rarity      rarity.Rarity           `yaml:"rarity"`
	Variety     variety.Variety         `yaml:"variety"`
	Category    objectcategory.Category `yaml:"category"`
	Potency     potency.Value           `yaml:"potency"`
	Effects     []stats.Effect          `yaml:"effects"`
	Pod         int                     `yaml:"pod"`
	Slot        slot.Slot               `yaml:"destinedslot"`
}

var folderConfigPath = "data/template/armor_sets"

var cachedConfig []yamlConfig

func loadFromSource() error {
	files, err := yaml.GetFilesFromSource(folderConfigPath, true)
	if err != nil {
		return fmt.Errorf("Error reading armor set folder: %v\n", err)
	}

	if errs := handleLoadFromFiles(files); len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func handleLoadFromFiles(files []string) []error {
	errs := []error{}
	for _, file := range files {
		if err := loadFromFile(file); err != nil {
			errs = append(errs, fmt.Errorf("Error loading armor set from file %s: %v", file, err))
			continue
		}
	}

	return errs
}

func loadFromFile(fileAddress string) error {
	params, err := yaml.LoadFromFile[yamlConfig](fileAddress)

	if err != nil {
		return err
	}

	cachedConfig = append(cachedConfig, *params)

	return nil
}

// - Collection - //
func List() []Gear {
	if cachedConfig == nil {
		if err := loadFromSource(); err != nil {
			fmt.Printf("Error loading dice: %v\n", err)
			return nil
		}
	}

	return makeSet(cachedConfig)
}

func makeSet(configs []yamlConfig) []Gear {
	result := make([]Gear, 0, len(configs))

	for _, config := range configs {
		if location := newFromYaml(config); location != nil {
			result = append(result, *location)
		}
	}

	return result
}

func newFromYaml(config yamlConfig) *Gear {
	newGear, err := NewGear(Template{
		Name:        config.Name,
		Description: config.Description,
		Rarity:      config.Rarity,
		Variety:     config.Variety,
		Category:    config.Category,
		Potency:     config.Potency,
		Effects:     config.Effects,
		Pod:         config.Pod,
		Slot:        config.Slot,
	})

	if err != nil {
		fmt.Printf("Error creating gear from YAML: %v\n", err)
		return nil
	}

	return newGear
}

func FindEquipementByName(name string) []Gear {
	if cachedConfig == nil {
		if err := loadFromSource(); err != nil {
			fmt.Printf("Error loading armor sets: %v\n", err)
			return nil
		}
	}

	var serchResults []yamlConfig
	for _, gear := range cachedConfig {
		if gear.Set == name {
			serchResults = append(serchResults, gear)
		}
	}

	return makeSet(serchResults)
}
