package t

import (
	"errors"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Translator struct {
	Local *i18n.Localizer
}

var translator *Translator

func Localize(id string, data ...map[string]any) string {
	return translator.Localize(id, data...)
}

func NewError(id string, data ...map[string]any) error {
	return translator.NewError(id, data...)
}

func (t *Translator) Localize(id string, data ...map[string]any) string {
	cfg := &i18n.LocalizeConfig{MessageID: id}
	if len(data) > 0 {
		cfg.TemplateData = data[0]
	}
	return t.Local.MustLocalize(cfg)
}

func (t *Translator) NewError(id string, data ...map[string]any) error {
	return errors.New(t.Localize(id, data...))
}
