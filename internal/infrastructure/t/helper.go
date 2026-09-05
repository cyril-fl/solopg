package t

import (
	"errors"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Localize(id string) string {
	return Local.MustLocalize(&i18n.LocalizeConfig{MessageID: id})
}

func NewError(config *i18n.LocalizeConfig) error {
	return errors.New(Local.MustLocalize(config))
}
