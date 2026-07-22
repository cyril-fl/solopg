package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func JsonifiedLog(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(string(jsonData))
}