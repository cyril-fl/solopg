package portal

import (
	"solopg/internal/domain/card/locations"
	"solopg/internal/domain/gameplay/dice"
	"solopg/internal/infrastructure/config"
	"solopg/internal/infrastructure/yaml"
)

/*
TODO
MEDIUM add a cache. a voir si pas deja fais dans Location !
*/
// - Configuration & caching - //
var fileConfigPath = config.Current.StructureFiles.Portal

var cachedConfig []locations.Location

func loadPresetLocations() ([]locations.Location, error) {
	folderName := fileConfigPath
	folderContents, err := yaml.GetFolderFiles(folderName)
	if err != nil {
		return nil, err
	}

	var locationsList []locations.Location
	for _, fileName := range folderContents {
		filePath := fileConfigPath + "/" + fileName

		location, err := locations.FromFile(filePath)
		if err != nil {
			return nil, err
		}
		locationsList = append(locationsList, *location)
	}

	return locationsList, nil
}

// - Portal - //
func Teleport() (*locations.Location, error) {
	list, err := loadPresetLocations()
	if err != nil {
		return nil, err
	}

	roll := dice.Roll(len(list))
	selectedLocation := list[roll-1]

	return &selectedLocation, nil
}
