package src

import (
	"solopg/app/services/mango/repository"
	"solopg/app/tui"

	"solopg/app/services/mango"
	initi19n "solopg/app/services/process/initI19n"
	"solopg/app/services/process/processkit"
	"solopg/app/services/process/runsession"
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

	translation := initi19n.Process(config.Current.I18n, "")
	session := runsession.Process(db)

	processes := []processkit.Processable{
		translation,
		session,
	}

	return processkit.HandleProcess(processes)
}

func Try() error {

	// if err := i18n.Init(config.Current.I18n, ""); err != nil {
	// 	return err
	// }

	return repository.TestLog()
}
