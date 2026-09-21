package log

import (
	"encoding/json"
	"fmt"
	"os"
)

func ParseJson(data interface{}) {
	json, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(string(json))
}
