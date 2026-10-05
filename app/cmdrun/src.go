package cmdrun

import (
	cmdrunrunsession "solopg/app/cmdrun/services/process/runsession"
	"solopg/app/shared/services/mango"
	"solopg/app/shared/services/process"
	initI19n "solopg/app/shared/services/process/initI19n"
	initlogger "solopg/app/shared/services/process/initloggerepository"
	"solopg/app/shared/utils/cleanui"
	"solopg/config"
)

func RunCmdRun() error {
	var db *mango.Mongo
	var err error

	cleanui.Run()

	if db, err = mango.Init(); err != nil {
		return err
	}
	defer db.Disconnect()

	/* --- NOTE Everything above this line is non loggable --- */
	/*
		- STEP 1: Initialize translation bundle
		- STEP 2: Initialize logger repository
		- STEP 3: Run session
	*/

	logger := initlogger.Process(db)
	translation := initI19n.Process(config.Current.I18n, "")
	session := cmdrunrunsession.Process(db)

	processes := []process.Processable{
		logger,
		translation,
		session,
	}

	return process.HandleProcess(processes)
}
