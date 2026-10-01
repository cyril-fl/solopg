package i19n

import (
	"errors"
	// "solopg/app/services/t"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Localize(id string, data ...map[string]any) string {
	return lookupMessage(cache.local, id, data...)
}

func lookupErrorMessage(id string, data ...map[string]any) string {
	return lookupMessage(cache.localerror, id, data...)
}

func lookupMessage(t *i18n.Localizer, id string, data ...map[string]any) string {
	// if true {
	// 	return fmt.Sprintf("{{%s}}", id)
	// }

	cfg := &i18n.LocalizeConfig{MessageID: id}
	if len(data) > 0 {
		cfg.TemplateData = data[0]
	}

	localized, err := t.Localize(cfg)
	if err != nil {
		return id
	}

	return localized
}

func NewError(id string, data ...map[string]any) error {
	return errors.New(Localize(id, data...))
}
func NewCatalogError(id string, data ...map[string]any) error {
	return errors.New(lookupErrorMessage(id, data...))
}
