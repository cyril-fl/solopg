package characters

import (
	"errors"
	"fmt"
	"solopg/internal/domain/card/attributes/rarity"
	"solopg/internal/domain/card/attributes/variety"
	"solopg/internal/domain/card/characters/wallet"
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/t"
	"solopg/internal/infrastructure/yaml"
)

// - Configuration & caching - //
type yamlConfig struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Rarity      string `yaml:"rarity"`
	Variety     string `yaml:"variety"`
	Class       string `yaml:"class"`
	Race        string `yaml:"race"`
	// Stats       []stats.YamlEffectConfig `yaml:"stats"`
	// Equipment   []yaml.YamlGearConfig    `yaml:"equipment"`
	// Inventory   []yaml.YamlObjectConfig  `yaml:"inventory"`
	Wallet wallet.Wallet `yaml:"wallet"`
}

var folderConfigPath = config.Current.Documents.Folders.Characters

var cachedConfig []yamlConfig

func loadFromSource() error {
	files, err := yaml.GetFilesFromSource(folderConfigPath, true)
	if err != nil {
		return t.NewError("error.locations.load_folder", map[string]any{"Folder": folderConfigPath, "Error": err})
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
		return t.NewError("error.locations.load", map[string]any{"Error": err})
	}
	/*
		TODO LOW exemple a suivre pour les autres variete et autres pourquoi ? si ca vien d'une carte loader, c'est valider !
	*/
	variety.MakeDefault(params.Variety)

	cachedConfig = append(cachedConfig, *params)

	return nil
}

// - Collection - //
func List() []Character {
	if cachedConfig == nil {
		if err := loadFromSource(); err != nil {
			fmt.Printf("Error loading dice: %v\n", err)
			return nil
		}
	}

	return makeSet(cachedConfig)
}

func makeSet(configs []yamlConfig) []Character {
	result := make([]Character, 0, len(configs))

	for _, config := range configs {
		if location := newFromYaml(config); location != nil {
			result = append(result, *location)
		}
	}

	return result
}

func newFromYaml(config yamlConfig) *Character {
	newLocation, err := New(Template{
		Name:        config.Name,
		Description: config.Description,
		Rarity:      rarity.MakeDefault(config.Rarity),
		Variety:     variety.MakeDefault(config.Variety),
		// Stats:       stats.MakeStatsFromYamlConfig(config.Stats),
		Wallet: config.Wallet,
		// Equipment:   equipment.MakeArmorSetFromYamlConfig(config.Equipment),
		// Inventory:   objects.MakeObjectArrayFromYamlConfig(config.Inventory),

		/*
			TODO HIGH implementer class et race
		*/
	})

	if err != nil {
		fmt.Printf("Error creating new character: %v\n", err)
		return nil
	}

	return newLocation
}
