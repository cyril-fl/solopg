package cmdserve

import (
	cmdserverunsession "solopg/app/cmdserve/service/runsession"
	"solopg/app/shared/services/logs"
	"solopg/app/shared/services/mango"
	"solopg/app/shared/services/mango/repository"
	"solopg/app/shared/services/process"
	"solopg/app/shared/services/process/initI19n"
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

	translation := initI19n.Process(config.Current.I18n, "")
	session := cmdserverunsession.Process(func() repository.Watchable[logs.Log] {
		repo := repository.LogSystem()
		repo.SetDb(db)
		return repo
	})

	processes := []process.Processable{
		translation,
		session,
	}
	
	return process.HandleProcess(processes)
}
