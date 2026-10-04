package cmdtry

import (
	"fmt"
	"os"
	"strings"
)

// Actuellement enregistre un log en DB
func Try() error {

	file := strings.TrimSpace(os.Getenv("CONFIG_FILE"))
	fmt.Println("Loading config from file:", file)


	return nil

}
