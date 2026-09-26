package t

import (
	"fmt"
	"path/filepath"
	"solopg/app/services/yaml"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	y "gopkg.in/yaml.v3"
)

var (
	bundle     *i18n.Bundle
	local      *i18n.Localizer
	localerror *i18n.Localizer
)

func Init(cfg Config, locale string) error {
	if err := cfg.Validate(); err != nil {
		return err
	}

	defaultLocale, err := cfg.getDefaultLocale(cfg.Default)
	if err != nil {
		return err
	}

	if err := initBundle(cfg, defaultLocale); err != nil {
		return err
	}

	localizer, err := newLocalizer(cfg, locale, defaultLocale)
	if err != nil {
		return err
	}
	local = localizer

	errorlocalizer, err := newLocalizer(cfg, "en", defaultLocale)
	if err != nil {
		return err
	}
	localerror = errorlocalizer

	return nil
}

// - Bundle - //
func initBundle(cfg Config, defaultLocale *Locale) error {
	tag, err := defaultLocale.parseTag()
	if err != nil {
		return err
	}

	bundle = i18n.NewBundle(tag)
	bundle.RegisterUnmarshalFunc(string(FormatYAML), y.Unmarshal)

	if err := registerLocaleFiles(bundle, cfg); err != nil {
		return err
	}

	return nil
}

func registerLocaleFiles(bundle *i18n.Bundle, cfg Config) error {
	files, err := loadLocaleFile(cfg)
	if err != nil {
		return err
	}

	for _, path := range files {
		if _, err := bundle.LoadMessageFile(path); err != nil {
			return fmt.Errorf("load message file '%s': %w", filepath.Base(path), err)
		}
	}

	return nil
}

func loadLocaleFile(cfg Config) ([]string, error) {
	files, err := yaml.GetFilesFromSource(cfg.Dir, true)
	if err != nil {
		return nil, err
	}

	assertedFiles := make([]string, 0)
	for _, path := range files {
		if !(filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml") {
			continue
		}

		name := filepath.Base(path)
		name = strings.TrimSuffix(name, filepath.Ext(name))
		if _, err := cfg.getLocaleByISO(name); err != nil {
			continue
		}

		assertedFiles = append(assertedFiles, path)
	}

	return assertedFiles, nil
}

// - Localizer - //
func newLocalizer(cfg Config, lang string, defaultLocale *Locale) (*i18n.Localizer, error) {
	locale, err := getLocale(cfg, lang, defaultLocale)
	if err != nil {
		return nil, err
	}

	tag, err := locale.parseTag()
	if err != nil {
		return nil, err
	}

	defaultTag, err := defaultLocale.parseTag()
	if err != nil {
		return nil, err
	}

	return i18n.NewLocalizer(bundle, tag.String(), defaultTag.String()), nil
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
