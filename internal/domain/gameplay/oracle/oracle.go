package oracle

import (
	"errors"
	"fmt"

	"solopg/internal/domain/gameplay/dice"
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/t"
	"solopg/internal/infrastructure/yaml"
)

// - Configuration & caching - //
type yamlConfig struct {
	ID        string     `yaml:"id"`
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

func load() error {
	folderContents, err := loadFromFolder()
	if err != nil {
		return err
	}

	errs := []error{}
	for _, file := range folderContents {
		filepath := fmt.Sprintf("%s/%s", folderConfigPath, file)

		if err := loadFromFile(filepath); err != nil {
			errs = append(errs, t.NewError("error.oracle.load_file", map[string]any{"File": file, "Error": err}))
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
		return nil, t.NewError("error.oracle.load_folder", map[string]any{"Folder": folderConfigPath, "Error": err})
	}

	return list, nil
}

func loadFromFile(fileAddress string) error {
	params, err := yaml.LoadFromFile[yamlConfig](fileAddress)

	if err != nil {
		return t.NewError("error.oracle.load", map[string]any{"Error": err})
	}

	if params.Dice <= 0 {
		return t.NewError("error.oracle.invalid_dice", map[string]any{"ID": params.ID, "Dice": params.Dice})
	}

	if len(params.Intervals) == 0 {
		return t.NewError("error.oracle.no_intervals", map[string]any{"ID": params.ID})
	}

	cachedConfig = append(cachedConfig, *params)

	return nil
}

// - Oracle - //
type Oracle struct {
	id        string
	dice      int
	typ       string
	visible   bool
	intervals []interval
}

func new(config yamlConfig) Oracle {
	return Oracle{
		id:        config.ID,
		dice:      config.Dice,
		typ:       config.Type,
		visible:   config.Visible,
		intervals: config.Intervals,
	}
}

func GetByID(id string) (*Oracle, error) {
	oracles := List()

	for _, oracle := range oracles {
		if oracle.id == id {
			return &oracle, nil
		}
	}

	return nil, t.NewError("error.oracle.not_found", map[string]any{"ID": id})
}

func (o Oracle) ID() string {
	return o.id
}

func (o Oracle) IsVisible() bool {
	return o.visible
}

// - Collection - //
func List() []Oracle {

	if cachedConfig == nil {
		err := load()
		if err != nil {
			fmt.Printf("Error loading oracles: %v\n", err)
			return nil
		}
	}

	return makeSet(cachedConfig)
}

func makeSet(configs []yamlConfig) []Oracle {
	result := make([]Oracle, 0, len(configs))

	for _, config := range configs {
		result = append(result, new(config))
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
					"ID": o.id, "Actual": fmt.Sprintf("%T", interval.Result), "Expected": fmt.Sprintf("%T", res),
				})
			}

			return &Result[T]{
				Roll:     value,
				Critical: interval.Critical,
				Result:   res,
			}, nil
		}
	}

	return nil, t.NewError("error.oracle.no_matching_interval", map[string]any{"ID": o.id, "Roll": value})
}
