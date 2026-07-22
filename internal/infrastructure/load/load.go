package load

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
