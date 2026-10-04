package debug

import (
	"fmt"
	"solopg/app/shared/utils/transform"
)

func Json(data interface{}) {
	json := transform.ParseJson(data)
	fmt.Println(string(json))
}