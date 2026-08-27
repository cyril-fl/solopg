package app

import (
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/i18n"
	"solopg/internal/infrastructure/mongo"
	"solopg/internal/platform/jsonlog"
)

func Start() error {
	if err := i18n.Init(config.Current.I18n, ""); err != nil {
		return err
	}

	db, err := mongo.Connect()
	if err != nil {
		return err
	}

	defer mongo.Disconnect(db)

	return runSession(db)
}

func Try() error {

	jsonlog.JsonifiedLog(config.Current)

	return nil
}
