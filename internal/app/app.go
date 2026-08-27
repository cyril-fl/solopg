package app

import (
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/mongo"
	"solopg/internal/infrastructure/t"
	"solopg/internal/platform/log"
)

func Start() error {
	if err := t.Init(config.Current.I18n, ""); err != nil {
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

	log.ParseJson(config.Current)

	return nil
}
