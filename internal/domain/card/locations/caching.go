package locations

import (
	"errors"
	"fmt"
	"solopg/internal/domain/card/attributes/rarity"
	"solopg/internal/domain/card/attributes/stats"
	"solopg/internal/domain/card/attributes/variety"
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/t"
	"solopg/internal/infrastructure/yaml"
)

// - Configuration & caching - //
type yamlConfig struct {
	Name        string                   `yaml:"name"`
	Description string                   `yaml:"description"`
	Rarity      string                   `yaml:"rarity"`
	Variety     string                   `yaml:"variety"`
	Effects     []stats.YamlEffectConfig `yaml:"effects"`
}

var folderConfigPath = config.Current.Documents.Folders.Locations

var cachedConfig []yamlConfig

func load() error {
	folderContents, err := loadFromFolder()
	if err != nil {
		return err
	}

	errs := []error{}
	for _, file := range folderContents {
		filepath := fmt.Sprintf("%s/%s", folderConfigPath, file)

		if err := loadFromFile(filepath); err != nil {
			errs = append(errs, t.NewError("error.locations.load_file", map[string]any{"File": file, "Error": err}))
			continue
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func loadFromFolder() ([]string, error) {
	list, err := yaml.GetFolderFiles(folderConfigPath)

	if err != nil {
		return nil, t.NewError("error.locations.load_folder", map[string]any{"Folder": folderConfigPath, "Error": err})
	}

	return list, nil
}

func loadFromFile(fileAddress string) error {
	params, err := yaml.LoadFromFile[yamlConfig](fileAddress)

	if err != nil {
		return t.NewError("error.locations.load", map[string]any{"Error": err})
	}
	/*
		TODO
		LOW exemple a suivre pour les autres variete et autres pourquoi ? si ca vien d'une carte loader, c'est valider !
	*/
	variety.MakeDefault(params.Variety)

	cachedConfig = append(cachedConfig, *params)

	return nil
}

// - Collection - //
func List() []Location {
	if cachedConfig == nil {
		err := load()
		if err != nil {
			fmt.Printf("Error loading locations: %v\n", err)
			return nil
		}
	}

	return makeSet(cachedConfig)
}

func makeSet(configs []yamlConfig) []Location {
	result := make([]Location, 0, len(configs))

	for _, config := range configs {
		if location := newFromYaml(config); location != nil {
			result = append(result, *location)
		}
	}

	return result
}

func newFromYaml(config yamlConfig) *Location {
	newLocation, err := New(Template{
		Name:        config.Name,
		Description: config.Description,
		Rarity:      rarity.MakeDefault(config.Rarity),
		Variety:     variety.MakeDefault(config.Variety),
		Effects:     stats.MakeEffectFromYamlConfigArray(config.Effects),
	})

	if err != nil {
		fmt.Printf("Error creating new location: %v\n", err)
		return nil
	}

	return newLocation
}
