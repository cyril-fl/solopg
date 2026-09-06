package t

import (
	"fmt"
	"path/filepath"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

var (
	bundle *i18n.Bundle
	Local  *i18n.Localizer
)

func Init(cfg Config, locale string) error {
	if err := cfg.Validate(); err != nil {
		return err
	}

	defaultLocale, err := cfg.getDefaultLocale()
	if err != nil {
		return err
	}

	if err := initBundle(cfg, defaultLocale); err != nil {
		return err
	}

	if err := initLocalizer(cfg, locale, defaultLocale); err != nil {
		return err
	}

	return nil
}

// -- Bundle -- //
func initBundle(cfg Config, defaultLocale *Locale) error {
	tag, err := parseTag(defaultLocale)
	if err != nil {
		return err
	}

	bundle = i18n.NewBundle(tag)
	bundle.RegisterUnmarshalFunc(string(FormatYAML), yaml.Unmarshal)

	if err := loadLocaleFiles(bundle, cfg); err != nil {
		return err
	}

	return nil
}

func loadLocaleFiles(bundle *i18n.Bundle, cfg Config) error {
	for _, locale := range cfg.Locales {
		path := filepath.Join(cfg.Dir, locale.File)
		if _, err := bundle.LoadMessageFile(path); err != nil {
			return fmt.Errorf("load locale %s from %s: %w", locale.Code, path, err)
		}
	}

	return nil
}

// -- Localizer -- //
func initLocalizer(cfg Config, lang string, defaultLocale *Locale) error {
	locale, err := getLocale(cfg, lang, defaultLocale)
	if err != nil {
		return err
	}

	tag, err := parseTag(locale)
	if err != nil {
		return err
	}

	defaultTag, err := parseTag(defaultLocale)
	if err != nil {
		return err
	}

	Local = i18n.NewLocalizer(bundle, tag.String(), defaultTag.String())
	return nil
}

func getLocale(cfg Config, lang string, defaultLocale *Locale) (*Locale, error) {
	if lang == "" {
		return defaultLocale, nil
	}

	locale, err := cfg.getLocaleByCode(lang)
	if err != nil {
		return nil, fmt.Errorf("locale '%s' not found: %w", lang, err)
	}

	return locale, nil

}

func parseTag(locale *Locale) (language.Tag, error) {
	tag, err := language.Parse(locale.ISO)
	if err != nil {
		return language.Tag{}, fmt.Errorf("invalid locale '%s': %w", locale.ISO, err)
	}

	return tag, nil
}
