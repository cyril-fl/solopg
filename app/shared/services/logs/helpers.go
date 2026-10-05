package logs

import (
	interfass "solopg/app/shared/types/interface"
)

func ConvertStringablesToList[T interfass.Stringable](list []T) []Log {
	logList := make([]Log, 0, len(list))

	for _, item := range list {
		value, ok := any(item).(Log)
		if !ok {
			Error("error.invalid:conversion", map[string]any{
				"Subject": "item",
				"Error":   "item is not a logs.Log",
			})
			continue
		}

		logList = append(logList, value)
	}
	return logList
}
