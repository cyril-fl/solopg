package yaml

import "solopg/internal/domain/card/characters"

func CharacterFromFile(fileAddress string) (*characters.Character, error) {
	var characterParams characters.CharacterTemplate
	if err := loadYAMLFromFile(fileAddress, &characterParams); err != nil {
		return nil, err
	}

	return characters.NewCharacter(characterParams)
}
