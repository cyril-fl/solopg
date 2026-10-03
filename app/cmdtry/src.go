package cmdtry

import (
	"solopg/app/shared/services/logs"
	"solopg/app/shared/services/mango"
	"solopg/app/shared/services/mango/repository"
	initI19n "solopg/app/shared/services/process/initI19n"
	"solopg/config"
)

// Actuellement enregistre un log en DB
func Try() error {
	db := mango.New()
	db.Connect()
	if db.HasErrors() {
		return db.GetErrors()
	}
	defer db.Disconnect()

	repo := repository.LogSystem()
	repo.SetDb(db)
	logs.Init(repo)

	translation := initI19n.Process(config.Current.I18n, "")
	translation.Run()
	if translation.HasErr() {
		return translation.GetErr()
	}

	err := logs.NewError("error.required", map[string]any{
		"Subject":  "race",
		"Property": 1,
	})

	println(err.Error())

	return nil

}
