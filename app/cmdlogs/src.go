package cmdlogs

import (
	"fmt"
	"solopg/app/cmdlogs/services/process/filterlogs"
	"solopg/app/shared/services/mango"
	"solopg/app/shared/services/mango/repository"
	"solopg/app/shared/services/process/initI19n"
	"solopg/config"

	"github.com/spf13/pflag"
)

func RunCmdLogs(flags func() *pflag.FlagSet) error {
	var db *mango.Mongo
	var err error

	if db, err = mango.Init(); err != nil {
		return err
	}
	defer db.Disconnect()

	/* --- NOTE Everything above this line is non loggable --- */

	// REFACTOR MEDIUM
	translation := initI19n.Process(config.Current.I18n, "")
	translation.Run()
	if translation.HasErr() {
		return translation.GetErr()
	}

	// Trasformer ça en process
	repo := repository.NewLogSystemRepos()
	repo.SetDb(db)
	logs, err := repo.Load(nil)
	if err != nil {
		return err
	}

	// touver le moyen de lier le prcees qui a pas encore run et logs
	filter := filterlogs.Process(flags, logs)
	filter.Run()
	if filter.HasErr() {
		return filter.GetErr()
	}

	for _, log := range filter.GetResult() {
		fmt.Println(log.String())
	}

	return nil
}
