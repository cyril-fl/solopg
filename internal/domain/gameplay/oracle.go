package gameplay

import (
	"fmt"

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

	return nil, fmt.Errorf("oracle with ID %q not found", id)
}

func loadOracles() []*Oracle {
	list := loadOraclesList()

	for _, file := range list {
		oracle, err := loadOracleFromFile(fmt.Sprintf("%s/%s", fileAddress, file))
		if err != nil {
			fmt.Printf("Error loading oracle from file %s: %v\n", file, err)
			continue
		}
		_oracles = append(_oracles, oracle)
	}

	return _oracles
}

func loadOraclesList() []string {
	list, err := yaml.GetFolderFiles(fileAddress)
	if err != nil {
		fmt.Printf("Error loading oracles list from folder %s: %v\n", fileAddress, err)
		return []string{}
	}
	return list
}

func loadOracleFromFile(fileAddress string) (*Oracle, error) {
	oracle, err := yaml.LoadFromFile[Oracle](fileAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to load oracle: %w", err)
	}

	if oracle.Dice <= 0 {
		return nil, fmt.Errorf("oracle %q has an invalid dice value: %d", oracle.ID, oracle.Dice)
	}
	if len(oracle.Intervals) == 0 {
		return nil, fmt.Errorf("oracle %q has no intervals", oracle.ID)
	}

	return oracle, nil
}

func  RollOracle[T any](o *Oracle) (*OracleResult[T], error) {
	value := roll(o.Dice)
	for _, interval := range o.Intervals {
		if value >= interval.Min && value <= interval.Max {
			res, ok := interval.Result.(T)
			if !ok {
				return nil, fmt.Errorf("oracle %q has an interval with a result of type %T, expected %T", o.ID, interval.Result, res)
			}

			return &OracleResult[T]{
				Roll:     value,
				Critical: interval.Critical,
				Result:   res,
			}, nil
		}
	}

	return nil, fmt.Errorf("oracle %q has no interval matching roll %d", o.ID, value)
}