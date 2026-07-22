package main

import (
	"fmt"
	"os"
	"solopg/internal/app"
)

func main() {
	db, err := app.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer app.CloseDatabase(db)

	if err := app.Run(db); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
