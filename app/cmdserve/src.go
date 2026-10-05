package cmdserve

import (
	cmdserverunsession "solopg/app/cmdserve/service/runsession"
	"solopg/app/shared/services/logs"
	"solopg/app/shared/services/mango"
	"solopg/app/shared/services/mango/repository"
	"solopg/app/shared/services/process"
	"solopg/app/shared/services/process/initI19n"
	"solopg/app/shared/services/process/initloggerepository"
	"solopg/config"

	"github.com/spf13/pflag"
)

func RunCmdServe(flags func() *pflag.FlagSet) error {
	var db *mango.Mongo
	var err error

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

	logger := initloggerepository.Process(db)
	translation := initI19n.Process(config.Current.I18n, "")
	session := cmdserverunsession.Process(cmdserverunsession.ProcessTemplate[logs.Log]{
		GetRepository: func() repository.Watchable[logs.Log] {
			repo := repository.LogSystem()
			repo.SetDb(db)
			return repo
		},
		GetFlags: flags,
	})

	processes := []process.Processable{
		logger,
		translation,
		session,
	}

	return process.HandleProcess(processes)
}
