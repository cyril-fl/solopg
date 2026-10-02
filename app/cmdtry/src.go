package cmdtry

import (
	"solopg/app/cmdrun/services/logger"
	"solopg/app/shared/services/mango"
	"solopg/app/shared/services/mango/repository"
)

func Try() error {
	log := logger.New(logger.Template{
		Type:    logger.INFO,
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
	logger.Init(repo)

	return logger.TestLog(log)
}
