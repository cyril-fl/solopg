package mango

type Collection string

func (c Collection) String() string {
	return string(c)
}