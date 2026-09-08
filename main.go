package main

import (
	"solopg/cmd"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cmd.Execute()
}