package gameplay

import (
	"solopg/internal/domain/card/locations"
	"solopg/internal/infrastructure/yaml"
)

/*
	TODO add a cache.
	MEDIUM
*/
	const basePath = "data/template/locations"

func LoadPresetLocations() ([]locations.Location, error) {
	folderName := basePath
	folderContents, err := yaml.GetFolderFiles(folderName)
	if err != nil {
		return nil, err
	}

	var locationsList []locations.Location
	for _, fileName := range folderContents {
		filePath := basePath + "/" + fileName

		location, err := locations.FromFile(filePath)
		if err != nil {
			return nil, err
		}
		locationsList = append(locationsList, *location)
	}

	return locationsList, nil
}

func DrawLocations() (*locations.Location, error) {
	list, err := LoadPresetLocations()
	if err != nil {
		return nil, err
	}

	roll := roll(len(list))
	selectedLocation := list[roll-1]

	return &selectedLocation, nil
}
