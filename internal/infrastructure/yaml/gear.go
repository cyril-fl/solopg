package yaml

import "solopg/internal/domain/card/objects"

func GearFromFile(fileAddress string) (*objects.Gear, error) {
	var gearParams objects.GearTemplate
	if err := loadYAMLFromFile(fileAddress, &gearParams); err != nil {
		return nil, err
	}

	return objects.NewGear(gearParams)
}
