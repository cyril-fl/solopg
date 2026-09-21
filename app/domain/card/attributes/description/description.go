package description

type Description struct {
	Name        string
	Description string
}

func New(name string, description string) Description {
	return Description{
		Name:        name,
		Description: description,
	}
}
