package i18n

import (
	"fmt"
	"path/filepath"

	"encoding/json"

	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

var (
	bundle *goi18n.Bundle
	T      *goi18n.Localizer
)

func Init(cfg Config, requested string) error {
	if err := cfg.Validate(); err != nil {
		return err
	}

	defaultLocale, err := cfg.DefaultLocale()
	if err != nil {
		return err
	}

	defaultTag, err := language.Parse(defaultLocale.ISO)
	if err != nil {
		return fmt.Errorf("invalid default locale '%s': %w", defaultLocale.ISO, err)
	}

	bundle = goi18n.NewBundle(defaultTag)
	bundle.RegisterUnmarshalFunc(string(FormatJSON), json.Unmarshal)
	bundle.RegisterUnmarshalFunc(string(FormatYAML), yaml.Unmarshal)

	for _, locale := range cfg.Locales {
		path := filepath.Join(cfg.Dir, locale.File)
		if _, err := bundle.LoadMessageFile(path); err != nil {
			return fmt.Errorf("load locale %s from %s: %w", locale.Code, path, err)
		}
	}

	selectedCode := requested
	if selectedCode == "" {
		selectedCode = defaultLocale.Code
	}

	selectedLocale, err := cfg.LocaleByCode(selectedCode)
	if err != nil {
		return err
	}

	selectedTag, err := language.Parse(selectedLocale.ISO)
	if err != nil {
		return fmt.Errorf("invalid locale '%s': %w", selectedLocale.ISO, err)
	}

	T = goi18n.NewLocalizer(bundle, selectedTag.String(), defaultTag.String())
	return nil
}
