package characters

import (
	"errors"
	"fmt"
	"solopg/app/domain/card/attributes/rarity"
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/domain/card/attributes/variety"
	"solopg/app/domain/card/characters/classes"
	"solopg/app/domain/card/characters/wallet"
	"solopg/app/domain/card/objects"
	"solopg/app/domain/card/objects/equipment"
	"solopg/app/services/t"
	"solopg/app/services/yaml"
	"solopg/config"
)

// - Configuration & caching - //
type yamlConfig struct {
	Name        string              `yaml:"name"`
	Description string              `yaml:"description"`
	Rarity      string              `yaml:"rarity"`
	Variety     string              `yaml:"variety"`
	Class       string              `yaml:"class"`
	Race        string              `yaml:"race"`
	Stats       stats.Stats         `yaml:"stats"`
	Equipment   equipment.Equipment `yaml:"equipment"`
	Inventory   []objects.Object    `yaml:"inventory"`
	Wallet      wallet.Wallet       `yaml:"wallet"`
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
	TODO LOW Verrifier ou "MakeDefault" pourrait être utile
	ex: Si la data vien d'une carte loader, c'est valider 
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
	class, err := classes.InsertIfIsnt(config.Class)
	if err != nil {
		fmt.Printf("Error inserting class: %v\n", err)
		return nil
	}

	newLocation, err := New(Template{
		Name:        config.Name,
		Description: config.Description,
		Class:       class.GetName(),
		Race:        config.Race,
		Rarity:      rarity.MakeDefault(config.Rarity),
		Variety:     variety.MakeDefault(config.Variety),
		Stats:       config.Stats,
		Wallet:      config.Wallet,
		Equipment:   config.Equipment,
		Inventory:   config.Inventory,
	})

	if err != nil {
		fmt.Printf("Error creating new character: %v\n", err)
		return nil
	}

	return newLocation
}
