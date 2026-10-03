package logs

import (
	"errors"
	"solopg/app/shared/services/i19n"
)

func NewError(id string, data ...map[string]any) error {
	registerFromTemplate(Template{
		Type:    ERR,
		Message: i19n.Unlocalize(id, data...),
	})

	return errors.New(i19n.Localize(id, data...))
}

func NewCatalogError(id string, data ...map[string]any) error {
	registerFromTemplate(Template{
		Type:    ERR,
		Message: i19n.Unlocalize(id, data...),
	})

	return errors.New(i19n.Unlocalize(id, data...))
}

