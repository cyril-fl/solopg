package characters

import "fmt"

type Race string

const (
	Human Race = "Human"
)

func (r Race) Validate() error {
	switch r {
	case Human:
		return nil
	default:
		return fmt.Errorf("invalid race: %s", r)
	}
}
