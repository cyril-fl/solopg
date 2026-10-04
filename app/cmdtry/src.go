package cmdtry

import (
	"fmt"
	"solopg/app/shared/services/logs"
	"solopg/app/shared/services/mango"
	"solopg/app/shared/services/mango/repository"
	"solopg/app/shared/services/process/initI19n"
	"solopg/config"

	"github.com/spf13/cobra"
)

const (
	DB_ON       = true
	LOCALIZE_ON = true
	LOG_ON      = true
)

// Actuellement enregistre un log en DB
func Try(cobra *cobra.Command, args []string) error {
	var db *mango.Mongo
	var err error

	if db, err = setDb(); db != nil {
		defer db.Disconnect()
	} else if err != nil {
		return err
	}

	setI19n()
	setLogRepository(db)

	s := logs.Success("success.try", map[string]any{
		"Subject": "try",
		"Args":    args,
	})

	i := logs.Info("info.try", map[string]any{
		"Subject": "try",
		"Args":    args,
	})

	w := logs.Warning("warning.try", map[string]any{
		"Subject": "try",
		"Args":    args,
	})

	fmt.Println(s)
	fmt.Println(i)
	fmt.Println(w)

	return nil
}

func setDb() (*mango.Mongo, error) {
	if !DB_ON {
		return nil, nil
	}
	return mango.Init()
}

func setLogRepository(db *mango.Mongo) {
	if !LOG_ON {
		return
	}

	if db == nil {
		panic(fmt.Errorf("database connection is nil, cannot set log repository"))
	}

	repo := repository.LogSystem()
	repo.SetDb(db)
	logs.Init(repo)
}

func setI19n() {
	if !LOCALIZE_ON {
		return
	}
	translation := initI19n.Process(config.Current.I18n, "")
	translation.Run()
}
