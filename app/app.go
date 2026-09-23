package src

import (
	"solopg/app/domain/card/characters"
	"solopg/app/services/mongo"
	"solopg/app/services/t"
	"solopg/app/utils/log"
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

	// oracle.Log()
	list := characters.List()

	log.ParseJson(list)

	return nil
}
