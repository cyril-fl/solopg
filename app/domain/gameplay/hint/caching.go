package hint

import (
	"errors"
	"fmt"
	"path/filepath"

	"solopg/app/domain/gameplay/dice"
	"solopg/app/services/t"
	"solopg/app/services/yaml"
	"solopg/config"
)

// - Configuration & caching - //
type yamlConfig struct {
	Name   string   `yaml:"name"`
	Values []string `yaml:"values"`
}

var folderConfigPath = config.Current.Documents.Folders.Hint

var cachedConfig []yamlConfig

func loadFromSource() error {
	// TODO ameliorer. Si c'est un fichier on load si c'est un folder on load mais autrement
	// Appliquer cette logique la au oracles dice ect...
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
			errs = append(errs, fmt.Errorf("Error loading hint set from file %s: %v", file, err))
			continue
		}
	}

	return errs
}

func loadFromFile(fileAddress string) error {
	params, err := yaml.LoadFromFile[yamlConfig](fileAddress)

	if err != nil {
		return t.NewError("error.hint.load", map[string]any{"Error": err})
	}

	cachedConfig = append(cachedConfig, *params)

	return nil
}

// - Hint - //
type Hint struct {
	name   string
	values []string
}

func newFromYamlConfig(config yamlConfig) Hint {
	return Hint{
		name:   config.Name,
		values: config.Values,
	}
}

func GetByName(name string) (*Hint, error) {
	hints := List()

	for _, hint := range hints {
		if hint.name == name {
			return &hint, nil
		}
	}

	return nil, t.NewError("error.hint.not_found", map[string]any{"Name": name})
}

func (o Hint) Name() string {
	return o.name
}
func (o Hint) Values() []string {
	return o.values
}

// - Collection - //
func List() []Hint {
	if len(cachedConfig) == 0 {
		err := loadFromSource()
		if err != nil {
			return nil
		}
	}

	return makeSet(cachedConfig)
}

func ListFromFolder(path string) ([]string, error) {
	list, err := yaml.GetFolderFiles(path)
	if err != nil {
		return nil, t.NewError("error.hint.load_folder", map[string]any{"Folder": path, "Error": err})
	}

	errs := []error{}
	hints := []string{}
	for _, file := range list {
		if params, err := yaml.LoadFromFile[yamlConfig](filepath.Join(path, file)); err != nil {
			errs = append(errs, fmt.Errorf("Error loading hint from file %s: %v", file, err))
		} else {
			hints = append(hints, params.Name)
		}
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return hints, nil
}

func makeSet(configs []yamlConfig) []Hint {
	result := make([]Hint, 0, len(configs))

	for _, config := range configs {
		result = append(result, newFromYamlConfig(config))
	}

	return result
}

// - Helpers - //
func Roll(hint Hint) []string {
	res := []string{}
	values := hint.Values()

	for i := 0; i < 3; i++ {
		value := dice.Roll(len(values))
		res = append(res, values[value-1])
	}

	return res
}
