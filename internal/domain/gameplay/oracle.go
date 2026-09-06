package gameplay

import (
	"fmt"

	"solopg/internal/infrastructure/t"
	"solopg/internal/infrastructure/yaml"
)

var _oracles []*Oracle

type Oracle struct {
	ID        string     `yaml:"id"`
	Dice      int        `yaml:"dice"`
	Type      string     `yaml:"type"`
	Visible   bool       `yaml:"visible"`
	Intervals []Interval `yaml:"intervals"`
}

type Interval struct {
	Min      int  `yaml:"min"`
	Max      int  `yaml:"max"`
	Critical bool `yaml:"critical"`
	Result   any  `yaml:"result"`
}

type OracleResult[T any] struct {
	Roll     int  `yaml:"roll"`
	Critical bool `yaml:"critical"`
	Result   T    `yaml:"result"`
}

const fileAddress = "data/template/systems/rules/oracles"

func GetOracle() []*Oracle {
	if _oracles == nil {
		_oracles = loadOracles()
	}

	return _oracles
}

func GetOracleByID(id string) (*Oracle, error) {
	oracles := GetOracle()

	for _, oracle := range oracles {
		if oracle.ID == id {
			return oracle, nil
		}
	}

	return nil, t.NewError("error.oracle.not_found", map[string]any{"ID": id})
}

func loadOracles() []*Oracle {
	list := loadOraclesList()

	for _, file := range list {
		oracle, err := loadOracleFromFile(fmt.Sprintf("%s/%s", fileAddress, file))
		if err != nil {
			message := t.NewError("error.oracle.load_file", map[string]any{"File": file, "Error": err})
			fmt.Println(message)
			continue
		}
		_oracles = append(_oracles, oracle)
	}

	return _oracles
}

func loadOraclesList() []string {
	list, err := yaml.GetFolderFiles(fileAddress)
	if err != nil {
		message := t.NewError("error.oracle.load_folder", map[string]any{"Folder": fileAddress, "Error": err})
		fmt.Println(message)
		return []string{}
	}
	return list
}

func loadOracleFromFile(fileAddress string) (*Oracle, error) {
	oracle, err := yaml.LoadFromFile[Oracle](fileAddress)
	if err != nil {
		return nil, t.NewError("error.oracle.load", map[string]any{"Error": err})
	}

	if oracle.Dice <= 0 {
		return nil, t.NewError("error.oracle.invalid_dice", map[string]any{"ID": oracle.ID, "Dice": oracle.Dice})
	}
	if len(oracle.Intervals) == 0 {
		return nil, t.NewError("error.oracle.no_intervals", map[string]any{"ID": oracle.ID})
	}

	return oracle, nil
}

func RollOracle[T any](o *Oracle) (*OracleResult[T], error) {
	value := roll(o.Dice)
	for _, interval := range o.Intervals {
		if value >= interval.Min && value <= interval.Max {
			res, ok := interval.Result.(T)
			if !ok {
				return nil, t.NewError("error.oracle.invalid_result_type", map[string]any{
					"ID": o.ID, "Actual": fmt.Sprintf("%T", interval.Result), "Expected": fmt.Sprintf("%T", res),
				})
			}

			return &OracleResult[T]{
				Roll:     value,
				Critical: interval.Critical,
				Result:   res,
			}, nil
		}
	}

	return nil, t.NewError("error.oracle.no_matching_interval", map[string]any{"ID": o.ID, "Roll": value})
}
