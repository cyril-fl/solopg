package app

import (
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/mongo"
	"solopg/internal/platform/jsonlog"
)

func Start() error {
	db, err := mongo.Connect()
	if err != nil {
		return err
	}

	defer mongo.Disconnect(db)

	return runSession(db)
}

func Try() error {
	c := config.Load()

	jsonlog.JsonifiedLog(c)

	return nil
}
