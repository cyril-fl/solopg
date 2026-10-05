package yaml

import (
	"os"
	"solopg/app/shared/services/logs"

	yamlv3 "gopkg.in/yaml.v3"
)

func SaveToFile[T any](fileAddress string, value T) error {
	data, err := yamlv3.Marshal(value)
	if err != nil {
		// Ici
		return logs.RawError("marshal %s: %w", fileAddress, err)
	}

	if err := os.WriteFile(fileAddress, data, 0o644); err != nil {
		// Ici
		return logs.RawError("write %s: %w", fileAddress, err)
	}

	return nil
}

func SaveListToFile[T any](fileAddress string, values []T) error {
	return SaveToFile(fileAddress, values)
}
