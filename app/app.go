package src

import (
	"solopg/app/services/mongo"
	"solopg/app/services/t"
	"solopg/app/utils/session"
	"solopg/config"
)

func Start() error {
	if err := t.Init(config.Current.I18n, ""); err != nil {
		return err
	}

	session.ClearTui()

	db := mongo.NewMongo()
	db.Connect()
	if db.HasErrors() {
		return db.GetErrors()
	}

	defer db.Disconnect()

	return session.RunSession(db)
}

func Try() error {

	if err := t.Init(config.Current.I18n, ""); err != nil {
		return err
	}

	return nil
}
