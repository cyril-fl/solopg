package t

import (
	"fmt"
	"solopg/types"
)

type Config struct {
	Default string   `yaml:"default"`
	Dir     string   `yaml:"dir"`
	Format  Format   `yaml:"format"`
	Locales []Locale `yaml:"locales"`
}

type Format string

const (
	FormatJSON Format = "json"
	FormatYAML Format = "yaml"
)

type Locale struct {
	Code string `yaml:"code"`
	ISO  string `yaml:"iso"`
	Name string `yaml:"name"`
	File string `yaml:"file"`
}

func (cfg Config) Validate() error {
	if cfg.Dir == "" {
		return fmt.Errorf("i18n directory is not specified")
	}

	if !cfg.Format.IsValid() {
		return fmt.Errorf("invalid i18n format: %s", cfg.Format)
	}

	if len(cfg.Locales) == 0 {
		return fmt.Errorf("no locales specified in i18n configuration")
	}

	localeCodes := make(types.Set[string])
	for _, locale := range cfg.Locales {
		if locale.Code == "" {
			return fmt.Errorf("locale code is not specified for one of the locales")
		}
		if locale.ISO == "" {
			return fmt.Errorf("locale ISO code is not specified for locale '%s'", locale.Code)
		}
		if locale.Name == "" {
			return fmt.Errorf("locale name is not specified for locale '%s'", locale.Code)
		}
		if locale.File == "" {
			return fmt.Errorf("locale file is not specified for locale '%s'", locale.Code)
		}

		if localeCodes.Has(locale.Code) {
			return fmt.Errorf("duplicate locale code '%s'", locale.Code)
		}
		localeCodes.Add(locale.Code)
	}

	if !localeCodes.Has(cfg.Default) {
		return fmt.Errorf("default locale '%s' is not in the list of available locales", cfg.Default)
	}

	return nil
}

func (f Format) IsValid() bool {
	switch f {
	case FormatJSON, FormatYAML:
	default:
		return false
	}
	return true
}

func (cfg Config) LocaleByCode(code string) (*Locale, error) {
	for index := range cfg.Locales {
		if cfg.Locales[index].Code == code {
			return &cfg.Locales[index], nil
		}
	}

	return nil, fmt.Errorf("locale '%s' not found", code)
}

func (cfg Config) DefaultLocale() (*Locale, error) {
	return cfg.LocaleByCode(cfg.Default)
}
