package yaml

import "solopg/internal/domain/card/locations"

func LocationFromFile(fileAddress string) (*locations.Location, error) {
	var locationParams locations.LocationTemplate
	if err := loadYAMLFromFile(fileAddress, &locationParams); err != nil {
		return nil, err
	}

	return locations.NewLocation(locationParams)
}
