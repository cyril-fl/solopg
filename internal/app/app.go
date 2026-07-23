package app

import (
	"fmt"
	"os"
	"solopg/internal/app/game"
	"solopg/internal/infrastructure/mongo"
)

func Start() {
	db, err := mongo.Connect()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer mongo.Disconnect(db)

	if err := game.Start(db); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
