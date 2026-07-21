package types

type Description struct {
	Name        string
	Description string
}

func NewDescription(name string, description string) Description {
	return Description{
		Name:        name,
		Description: description,
	}
}