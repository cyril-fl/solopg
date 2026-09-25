package src

import (
	"solopg/app/domain/gameplay/oracle"
	"solopg/app/services/mongo"
	"solopg/app/services/t"
	"solopg/app/utils/log"
	"solopg/app/utils/session"
	"solopg/config"
)

func Start() error {
	if err := t.Init(config.Current.I18n, ""); err != nil {
		return err
	}

	session.ClearTui()

	db := mongo.NewMongo()
	db.Connect()
	if db.HasErrors() {
		return db.GetErrors()
	}

	defer db.Disconnect()

	return session.RunSession(db)
}

func Try() error {

	oracle.Log()
	list := oracle.List()

	// for _, spark := range list {
	// 	fmt.Printf("Spark: %s \n", spark.Name())
	// 	newValues := strings.Join(spark.Values(), ", \n")
	// 	fmt.Println("Values: ", newValues)
	// }
	log.ParseJson(list)

	return nil
}
