package main

import (
	"github.com/joho/godotenv"
	"solopg/cmd"
)

func main() {
	godotenv.Load()
	cmd.Execute()
}
"terminal.integrated.defaultProfile.linux": "bash"