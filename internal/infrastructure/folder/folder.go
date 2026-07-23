package folder

import (
	"os"
)

func GetContents(folderPath string) ([]string, error) {
	entries, err := os.ReadDir(folderPath)

	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		files = append(files, entry.Name())

	}

	return files, nil
}
