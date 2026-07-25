package yaml

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func loadYAMLFromFile(fileAddress string, target any) error {
	data, err := os.ReadFile(fileAddress)
	if err != nil {
		e := fmt.Errorf("read %s: %w", fileAddress, err)
		fmt.Println(e)
		return e
	}

	if err := yaml.Unmarshal(data, target); err != nil {
		e := fmt.Errorf("unmarshal %s: %w", fileAddress, err)
		fmt.Println(e)
		return e
	}

	return nil
}

type wantedEntries struct {
	directories bool
	files       bool
}

func getFolderEntries(folderPath string, wanted wantedEntries) ([]string, error) {
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	var names []string

	for _, entry := range entries {
		isDirectory := entry.IsDir()
		skipFile := !wanted.files && !isDirectory
		skipFolder := !wanted.directories && isDirectory

        if  skipFile {
            continue
        }

        if skipFolder {
            continue
        }

		names = append(names, entry.Name())
	}

	return names, nil
}

func GetFolderFiles(folderPath string) ([]string, error) {
	return getFolderEntries(folderPath, wantedEntries{directories: false, files: true})
}

func GetFolderDirectories(folderPath string) ([]string, error) {
	return getFolderEntries(folderPath, wantedEntries{directories: true, files: false})
}

func LoadFromFile[T any](fileAddress string) (*T, error) {
	var params T
	if err := loadYAMLFromFile(fileAddress, &params); err != nil {
		return nil, err
	}

	return &params, nil
}

func LoadListFromFile[T any](fileAddress string) ([]T, error) {
	var params []T
	if err := loadYAMLFromFile(fileAddress, &params); err != nil {
		return nil, err
	}

	return params, nil
}
