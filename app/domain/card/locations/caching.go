package locations

import (
	"errors"
	"fmt"
	"solopg/app/domain/card/attributes/rarity"
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/domain/card/attributes/variety"
	"solopg/app/services/i18n"
	"solopg/app/services/yaml"
	"solopg/config"
)

// - Configuration & caching - //
type yamlConfig struct {
	Name        string         `yaml:"name"`
	Description string         `yaml:"description"`
	Rarity      string         `yaml:"rarity"`
	Variety     string         `yaml:"variety"`
	Effects     []stats.Effect `yaml:"effects"`
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
			// i18N -- register
			errs = append(errs, i18n.NewError("error.loading:file", map[string]any{
				"Subject": i18n.Localize("location"),
				"File":    filepath,
				"Error":   err,
			}))
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
		// i18N -- register
		return nil, i18n.NewError("error.loading:folder", map[string]any{
			"Subject": i18n.Localize("location"),
			"Folder":  folderConfigPath,
			"Error":   err,
		})
	}

	return list, nil
}

func loadFromFile(fileAddress string) error {
	params, err := yaml.LoadFromFile[yamlConfig](fileAddress)

	if err != nil {
		// i18N -- register
		return i18n.NewError("error.loading:file", map[string]any{
			"Subject": i18n.Localize("location"),
			"File":    fileAddress,
			"Error":   err,
		})
	}

	/*
		TODO LOW Verrifier ou "MakeDefault" pourrait être utile
		ex: Si la data vien d'une carte loader, c'est valider

		Exemple a suivre pour les autres variete et autres pourquoi ? si ca vien d'une carte loader, c'est valider !
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
			// i18N -- register
			i18n.NewError("error.loading", map[string]any{
				"Subject": i18n.Localize("location"),
				"Error":   err,
			})
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
		Effects:     config.Effects,
	})

	if err != nil {
		// i18N -- register
		i18n.NewError("error.invalid:new", map[string]any{
			"Subject": i18n.Localize("location"),
			"Error":   err,
		})
		return nil
	}

	return newLocation
}
