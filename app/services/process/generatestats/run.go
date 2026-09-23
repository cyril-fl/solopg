package generatestats

import (
	"errors"
	"fmt"
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/domain/card/characters/classes"
	"solopg/app/domain/card/characters/races"
	"solopg/app/domain/gameplay/oracle"
	"strconv"
)

type StatsGenerator struct {
	values map[string]string
	modifiers []stats.Modifier
	entropy string
	error  []error
}

func NewStatsGenerator(values map[string]string) *StatsGenerator {
	return &StatsGenerator{
		values: values,
		entropy: "stat_generation",
	}
}
// Getter & Setter
func (sg *StatsGenerator) GetStats() stats.Stats {
	s := stats.GetBasic()
	s.ApplyModifiers(sg.modifiers)
	return s
}

func (sg *StatsGenerator) GetModifiers() []stats.Modifier {
	return sg.modifiers
}

func (sg *StatsGenerator) GetError() error {
	return errors.Join(sg.error...)
}

func (sg *StatsGenerator) SetEntropy(entropy string) *StatsGenerator {
	sg.entropy = entropy
	return sg
}

// Handlers
// GenerateFromValues generates full build for character, based on provided values (class, race, stats and fallback oracle). 
func (sg *StatsGenerator) GenerateFromValues(params GenerateParams) {
	sg.reset()
	sg.makeModifiersFromValuesWithFallback()
	sg.makeModifiersFromAttributesValues()
	
	if params.Randomness {
		sg.makeRandomModifiersFromOracleValue()
	}
}

func (sg *StatsGenerator) makeModifiersFromValuesWithFallback() {
	rules, err := oracle.GetByID(sg.values["encounter"])
	if err != nil {
		sg.error = append(sg.error, fmt.Errorf("failed to get oracle for encounter: %w", err))
		return
	}

	for _, stat := range stats.List() {
		sg.modifiers = append(sg.modifiers, makeModifierWithRandomFallback(withFallbackTemplate{
			stat:     stat,
			value:    sg.values[stat.String()],
			fallback: *rules,
		}))
	}
}

func (sg *StatsGenerator) makeModifiersFromAttributesValues() {	
	if class := classes.FindByName(sg.values["class"]); class != nil {
		sg.modifiers = append(sg.modifiers, class.GetBonus()...)
	}
	if race := races.FindByName(sg.values["race"]); race != nil {
		sg.modifiers = append(sg.modifiers, race.GetBonus()...)
	}
}

func (sg *StatsGenerator) makeRandomModifiersFromOracleValue() {
	modifiers, err := makeModifiersFromOracle(sg.entropy)
	if err != nil {
		sg.error = append(sg.error, fmt.Errorf("failed to generate random modifiers: %w", err))
		return
	}

	sg.modifiers = append(sg.modifiers, modifiers...)
}

func (sg *StatsGenerator) reset() {
	sg.modifiers = nil
	sg.error = nil
}

// Helper
type GenerateParams struct {
	Randomness bool
}

type withFallbackTemplate struct {
	stat     stats.Stat
	value    string
	fallback oracle.Oracle
}

func makeModifierWithRandomFallback(params withFallbackTemplate) stats.Modifier {
	if value, err := strconv.Atoi(params.value); err == nil && params.value != "" {
		return stats.Modifier{
			Stat:  params.stat,
			Value: value,
		}
	}

	roll, err := oracle.Roll[int](params.fallback)
	if err != nil {
		fmt.Printf("Error rolling for stat %s: %v\n", params.stat.String(), err)
		return stats.Modifier{
			Stat:  params.stat,
			Value: 0,
		}
	}

	return stats.Modifier{
		Stat:  params.stat,
		Value: roll.Result,
	}
}

func makeModifiersFromOracle(id string) ([]stats.Modifier, error) {
	var modifiers []stats.Modifier
	rules, err := oracle.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get stat generation oracle: %w", err)
	}

	for _, stat := range stats.List() {
		roll, err := oracle.Roll[int](*rules)
		if err != nil {
			return nil, fmt.Errorf("failed to roll for stat %s: %w", stat.String(), err)
		}

		modifiers = append(modifiers, stats.Modifier{
			Stat:  stat,
			Value: roll.Result,
		})
	}
	
	return modifiers, nil
}