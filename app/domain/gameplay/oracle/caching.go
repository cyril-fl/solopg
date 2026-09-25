package oracle

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
	Name      string     `yaml:"name"`
	Dice      int        `yaml:"dice"`
	Type      string     `yaml:"type"`
	Visible   bool       `yaml:"visible"`
	Intervals []interval `yaml:"intervals"`
}

type interval struct {
	Min      int  `yaml:"min"`
	Max      int  `yaml:"max"`
	Critical bool `yaml:"critical"`
	Result   any  `yaml:"result"`
}

type Result[T any] struct {
	Roll     int  `yaml:"roll"`
	Critical bool `yaml:"critical"`
	Result   T    `yaml:"result"`
}

var folderConfigPath = config.Current.Documents.Folders.Oracle

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
			errs = append(errs, fmt.Errorf("Error loading oracle set from file %s: %v", file, err))
			continue
		}
	}

	return errs
}

func loadFromFile(fileAddress string) error {
	params, err := yaml.LoadFromFile[yamlConfig](fileAddress)

	if err != nil {
		return t.NewError("error.oracle.load", map[string]any{"Error": err})
	}

	if params.Dice <= 0 {
		return t.NewError("error.oracle.invalid_dice", map[string]any{"ID": params.Name, "Dice": params.Dice})
	}

	if len(params.Intervals) == 0 {
		return t.NewError("error.oracle.no_intervals", map[string]any{"ID": params.Name})
	}

	cachedConfig = append(cachedConfig, *params)

	return nil
}

// - Oracle - //
type Oracle struct {
	name      string
	dice      int
	typ       string
	visible   bool
	intervals []interval
}

func newFromYamlConfig(config yamlConfig) Oracle {
	return Oracle{
		name:      config.Name,
		dice:      config.Dice,
		typ:       config.Type,
		visible:   config.Visible,
		intervals: config.Intervals,
	}
}

func GetByName(name string) (*Oracle, error) {
	oracles := List()

	for _, oracle := range oracles {
		if oracle.name == name {
			return &oracle, nil
		}
	}

	return nil, t.NewError("error.oracle.not_found", map[string]any{"Name": name})
}

func (o Oracle) Name() string {
	return o.name
}

func (o Oracle) IsVisible() bool {
	return o.visible
}

// - Collection - //
func List() []Oracle {
	if cachedConfig == nil {
		if err := loadFromSource(); err != nil {
			fmt.Printf("Error loading oracles: %v\n", err)
			return nil
		}
	}

	return makeSet(cachedConfig)
}

func ListFromFolder(path string) ([]string, error) {
	list, err := yaml.GetFolderFiles(path)
	if err != nil {
		return nil, t.NewError("error.oracle.load_folder", map[string]any{"Folder": path, "Error": err})
	}

	errs := []error{}
	oracles := []string{}
	for _, file := range list {
		if params, err := yaml.LoadFromFile[yamlConfig](filepath.Join(path, file)); err != nil {
			errs = append(errs, fmt.Errorf("Error loading oracle from file %s: %v", file, err))
		} else {
			oracles = append(oracles, params.Name)
		}
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return oracles, nil
}

func makeSet(configs []yamlConfig) []Oracle {
	result := make([]Oracle, 0, len(configs))

	for _, config := range configs {
		result = append(result, newFromYamlConfig(config))
	}

	return result
}

// - Helpers - //
func Roll[T any](o Oracle) (*Result[T], error) {
	value := dice.Roll(o.dice)
	for _, interval := range o.intervals {
		if value >= interval.Min && value <= interval.Max {
			res, ok := interval.Result.(T)
			if !ok {
				return nil, t.NewError("error.oracle.invalid_result_type", map[string]any{
					"ID": o.name, "Actual": fmt.Sprintf("%T", interval.Result), "Expected": fmt.Sprintf("%T", res),
				})
			}

			return &Result[T]{
				Roll:     value,
				Critical: interval.Critical,
				Result:   res,
			}, nil
		}
	}

	return nil, t.NewError("error.oracle.no_matching_interval", map[string]any{"ID": o.name, "Roll": value})
}
