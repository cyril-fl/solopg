package src

import (
	"solopg/app/services/i18n"
	"solopg/app/services/mango/repository"

	"solopg/app/services/mango"
	"solopg/app/utils/session"
	"solopg/config"
)

func Start() error {
	if err := i18n.Init(config.Current.I18n, ""); err != nil {
		return err
	}

	session.ClearTui()

	db := mango.New()
	db.Connect()
	if db.HasErrors() {
		return db.GetErrors()
	}

	defer db.Disconnect()

	return session.RunSession(db)
}

func Try() error {

	// if err := i18n.Init(config.Current.I18n, ""); err != nil {
	// 	return err
	// }

	return repository.TestLog()
}
