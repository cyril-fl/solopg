package src

import (
	"solopg/app/services/logger"
	"solopg/app/services/mango"
	"solopg/app/services/mango/repository"
	initi19n "solopg/app/services/process/initI19n"
	initlogger "solopg/app/services/process/initLogger"
	"solopg/app/services/process/processkit"
	"solopg/app/services/process/runsession"
	"solopg/app/tui"
	"solopg/config"
)

func Start() error {
	var db *mango.Mongo
	var err error

	tui.Clear()

	if db, err = mango.Init(); err != nil {
		return err
	}
	defer db.Disconnect()




	/* --- NOTE Everything above this line is non loggable --- */

	logger := initlogger.Process(db)
	translation := initi19n.Process(config.Current.I18n, "")
	session := runsession.Process(db)

	processes := []processkit.Processable{
		logger,
		translation,
		session,
	}

	return processkit.HandleProcess(processes)
}

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
