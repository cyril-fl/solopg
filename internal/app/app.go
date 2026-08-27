package app

import (
	"solopg/internal/domain/card/effects"
	"solopg/internal/domain/gameplay"
	"solopg/internal/infrastructure/mongo"
)

func Start() error {
	db, err := mongo.Connect()
	if err != nil {
		return err
	}

	defer mongo.Disconnect(db)

	return runSession(db)
}

func Try() error {
	oracle, err := gameplay.GetOracleByID("stat_generation")
	if err != nil {
		return err
	}

	statsList := effects.ListStats()
	build := []effects.Modifier{}

	for _, stat := range statsList {
		roll, err := gameplay.RollOracle[int](oracle)
		if err != nil {
			return err
		}

		// jsonlog.JsonifiedLog(roll)

		build = append(build, effects.Modifier{
			Stat:  stat,
			Value: roll.Result,
		})
	}

	d := effects.BaseStats()
	d.ApplyModifiers(build)
	// jsonlog.JsonifiedLog(d)

	return nil
}
