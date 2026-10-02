package process

import "errors"

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
	errs []error
}

// Methods
func NewErr() *Process {
	return &Process{}
}

func (p *Process) GetErr() error {
	return errors.Join(p.errs...)
}

func (p *Process) SetErr(err error) {
	if err != nil {
		p.errs = append(p.errs, err)
	}
}

func (p *Process) HasErr() bool {
	return len(p.errs) > 0
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