package t

import (
	"fmt"
	"os"
	"slices"
	"solopg/types"
)

// - Locale - //
type Locale struct {
	Code string `yaml:"code"`
	ISO  string `yaml:"iso"`
	Name string `yaml:"name"`
	File string `yaml:"file"`
}

/*
TODO voir pour faire de add on et etendre les ficher ect voir comment on pourrais faire, de maniere a ajouter par
exemple
-error
	-en
	-fr
- success
	-en
	-en
de maniere a ensuite faire un reduce et creée un fichier unique et temporaire par langue mais permettre une meilleur gestions.
*/
// - Format - //
type Format string

const (
	FormatJSON Format = "json"
	FormatYAML Format = "yaml"
)

// - Config - //
type Config struct {
	Default string   `yaml:"default"`
	Dir     string   `yaml:"dir"`
	Format  Format   `yaml:"format"`
	Locales []Locale `yaml:"locales"`
}

func (cfg Config) getDefaultLocale() (*Locale, error) {
	return cfg.getLocaleByCode(cfg.Default)
}

func (cfg Config) getLocaleByCode(code string) (*Locale, error) {
	for index := range cfg.Locales {
		if cfg.Locales[index].Code == code {
			return &cfg.Locales[index], nil
		}
	}

	return nil, fmt.Errorf("locale '%s' not found", code)
}

func (cfg Config) Validate() error {
	if err := isDirectoryValid(cfg.Dir); err != nil {
		return err
	}

	if err := isFileFormatValid(cfg.Format); err != nil {
		return err
	}

	if err := isLocalesConfigValid(cfg.Default, cfg.Locales); err != nil {
		return err
	}

	return nil
}

// - Validation - //
func isDirectoryValid(path string) error {
	if path == "" {
		return fmt.Errorf("i18n directory is not specified")
	}

	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("check i18n directory '%s': %w", path, err)
	}

	if !info.IsDir() {
		return fmt.Errorf("i18n directory '%s' does not exist", path)
	}

	return nil
}

func isFileFormatValid(format Format) error {
	if !slices.Contains([]Format{FormatJSON, FormatYAML}, format) {
		return fmt.Errorf("invalid i18n format '%s', must be 'json' or 'yaml'", format)
	}

	return nil
}

func isLocalesConfigValid(defaultLocale string, locales []Locale) error {
	localeSet := make(types.Set[string])

	for _, locale := range locales {
		if localeSet.Has(locale.Code) {
			return fmt.Errorf("duplicate locale code '%s'", locale.Code)
		}

		if err := isLocaleValid(locale); err != nil {
			return err
		}

		localeSet.Add(locale.Code)
	}

	if !localeSet.Has(defaultLocale) {
		return fmt.Errorf("default locale '%s' is not in the list of available locales", defaultLocale)
	}

	return nil
}

func isLocaleValid(locale Locale) error {
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
	return nil
}
