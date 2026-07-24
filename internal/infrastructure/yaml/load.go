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

func GetFolderContents(folderPath string) ([]string, error) {
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
