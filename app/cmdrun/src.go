package cmdrun

import (
	"solopg/app/cmdrun/services/process/runsession"
	cmdruntui "solopg/app/cmdrun/tui"
	"solopg/app/shared/services/mango"
	"solopg/app/shared/services/process"
	initI19n "solopg/app/shared/services/process/initI19n"
	initlogger "solopg/app/shared/services/process/initloggerepository"
	"solopg/config"
)

func Start() error {
	var db *mango.Mongo
	var err error

	cmdruntui.Clear()

	if db, err = mango.Init(); err != nil {
		return err
	}
	defer db.Disconnect()

	/* --- NOTE Everything above this line is non loggable --- */

	logger := initlogger.Process(db)
	translation := initI19n.Process(config.Current.I18n, "")
	session := runsession.Process(db)

	processes := []process.Processable{
		logger,
		translation,
		session,
	}

	return process.HandleProcess(processes)
}
