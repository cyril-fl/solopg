package millstats

import (
	"errors"
	"os"
	"solopg/app/cmdrun/domain/card/attributes/stats"
	"solopg/app/cmdrun/domain/card/characters/classes"
	"solopg/app/cmdrun/domain/card/characters/races"
	"solopg/app/cmdrun/domain/gameplay/oracle"
	"solopg/app/shared/services/i19n"
	"strconv"
)

type mill struct {
	values    map[string]string
	modifiers []stats.Modifier
	oraclesId string
	error     []error
}

func New() *mill {
	return NewWithValues(nil)
}

func NewWithValues(values map[string]string) *mill {
	return &mill{
		values:    values,
		oraclesId: "stat_generation",
	}
}

// Getter & Setter
func (m *mill) GetStats() stats.Stats {
	s := stats.GetBasic()
	s.ApplyModifiers(m.modifiers)
	return s
}

func (m *mill) GetModifiers() []stats.Modifier {
	return m.modifiers
}

func (m *mill) GetError() error {
	return errors.Join(m.error...)
}

func (m *mill) HasError() bool {
	return len(m.error) > 0
}

func (m *mill) SetOracleId(entropy string) *mill {
	m.oraclesId = entropy
	return m
}

// Handlers
// GenerateFromValues generates full build for character, based on provided values (class, race, stats and fallback oracle).
func (m *mill) GenerateFromValues(params FromValuesParams) {
	m.reset()
	m.makeModifiersFromValuesWithFallback()
	m.makeModifiersFromAttributesValues()

	if params.Randomness {
		m.makeRandomModifiersFromOracleValue()
	}
}

func (m *mill) GenerateFromOracle() {
	m.reset()
	m.makeRandomModifiersFromOracleValue()
}

func (m *mill) makeModifiersFromValuesWithFallback() {
	rules, err := oracle.GetByName(m.values["encounter"])
	if err != nil {
		// i18N -- register
		m.error = append(m.error, i19n.NewError("error.not_found:id-error", map[string]any{
			"Subject": i19n.Localize("encounter"),
			"ID":      i19n.Localize(m.values["encounter"]),
			"Error":   err,
		}))
		return
	}

	for _, stat := range stats.List() {
		m.modifiers = append(m.modifiers, makeModifierWithRandomFallback(withFallbackTemplate{
			stat:     stat,
			value:    m.values[stat.String()],
			fallback: *rules,
		}))
	}
}

func (m *mill) makeModifiersFromAttributesValues() {
	if class := classes.FindByName(m.values["class"]); class != nil {
		m.modifiers = append(m.modifiers, class.GetBonus()...)
	}
	if race := races.FindByName(m.values["race"]); race != nil {
		m.modifiers = append(m.modifiers, race.GetBonus()...)
	}
}

func (m *mill) makeRandomModifiersFromOracleValue() {
	modifiers, err := makeModifiersFromOracle(m.oraclesId)
	if err != nil {
		// i18N -- register
		m.error = append(m.error, i19n.NewError("error.invalid:new", map[string]any{
			"Subject": i19n.Localize("stat"),
			"Error":   err,
		}))
		return
	}

	m.modifiers = append(m.modifiers, modifiers...)
}

func (m *mill) reset() {
	m.modifiers = nil
	m.error = nil
}

// Helper
type FromValuesParams struct {
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
		// i18N -- register
		cwd, cwderr := os.Getwd()
		i19n.NewError("error.unexpected", map[string]any{
			"Path":  cwd,
			"Error": errors.Join(cwderr, err),
		})
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
	rules, err := oracle.GetByName(id)
	if err != nil {
		// i18N -- register
		return nil, i19n.NewError("error.invalid:new: %w", map[string]any{
			"Subject": i19n.Localize("stat"),
			"Error":   err,
		})
	}

	for _, stat := range stats.List() {
		roll, err := oracle.Roll[int](*rules)
		if err != nil {
			// i18N -- register
			cwd, cwderr := os.Getwd()
			return nil,
				i19n.NewError("error.unexpected", map[string]any{
					"Path":  cwd,
					"Error": errors.Join(cwderr, err),
				})
		}

		modifiers = append(modifiers, stats.Modifier{
			Stat:  stat,
			Value: roll.Result,
		})
	}

	return modifiers, nil
}
