package cmdtry

import (
	"solopg/app/shared/services/logs"
	"solopg/app/shared/services/mango"
	"solopg/app/shared/services/mango/repository"
)

func Try() error {
	log := logs.New(logs.Template{
		Type:    logs.INFO,
		Message: "This is a test log entry",
	})

	db := mango.New()
	db.Connect()
	if db.HasErrors() {
		return db.GetErrors()
	}
	defer db.Disconnect()

	repo := repository.NewLogSystemRepos()
	repo.SetDb(db)
	logs.Init(repo)

	return logs.TestLog(log)
}
