package process

import "solopg/app/shared/types/primitive"

// Interface
type Processable interface {
	Run()

	GetErr() error
	SetErr(err error)
	HasErr() bool
}

type ProcessableResult[T any] interface {
	Processable

	GetResult() T
}

type Process struct {
	primitive.Fallible
}

// Methods
func NewErr() *Process {
	return &Process{}
}

// Helper
func HandleProcess(processes []Processable) error {
	for _, p := range processes {
		p.Run()

		if p.HasErr() {
			return p.GetErr()
		}
	}

	return nil
}

/*
NOTE Snippet

type p struct {
	process.Process
	cache
}

type cache struct {
}

func Process() *p {
	return &p{

	}
}

func (i *p) Run() {}

func (i *p) GetResult() {
	logs.SilentWarning("error:not_implemented", map[string]any{
		"Function": "GetResult",
		"Subject":  "",
	})
}
*/
