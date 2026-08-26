package gameplay

import (
	"solopg/internal/domain/card/locations"
	"solopg/internal/infrastructure/yaml"
)

func LoadPresetLocations() ([]locations.Location, error) {
	folderName := "data/template/locations"
	folderContents, err := yaml.GetFolderFiles(folderName)
	if err != nil {
		return nil, err
	}

	var locationsList []locations.Location
	for _, fileName := range folderContents {
		filePath := "data/template/locations/" + fileName

		location, err := locations.FromFile(filePath)
		if err != nil {
			return nil, err
		}
		locationsList = append(locationsList, *location)
	}

	return locationsList, nil
}

// TODO add a cache.
func DrawLocations() (*locations.Location, error) {
	list, err := LoadPresetLocations()
	if err != nil {
		return nil, err
	}

	roll := roll(len(list))
	selectedLocation := list[roll-1]

	return &selectedLocation, nil
}
