package t

import (
	"errors"
	// "solopg/internal/infrastructure/t"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Localize(id string, data ...map[string]any) string {
	cfg := &i18n.LocalizeConfig{MessageID: id}
	if len(data) > 0 {
		cfg.TemplateData = data[0]
	}

	localized, err := Local.Localize(cfg)
	if err != nil {
		return id
	}

	return localized
}

func NewError(id string, data ...map[string]any) error {
	return errors.New(Localize(id, data...))
}
