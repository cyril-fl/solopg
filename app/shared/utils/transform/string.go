package transform

import (
	"strings"
)

func Capitalize(s string) string {
	// if true {
	// 	return fmt.Sprintf("** %s **", s)
	// }
	return strings.ToUpper(s[:1]) + s[1:]
}

func Uppercase(s string) string {
	// if true {
	// 	return fmt.Sprintf("** %s **", s)
	// }
	return strings.ToUpper(s)
}

func CleanJoin(sep string, parts ...string) string {
	res := []string{}
	for _, p := range parts {
		if p != "" {
			res = append(res, p)
		}
	}

	return strings.Join(res, sep)
}
