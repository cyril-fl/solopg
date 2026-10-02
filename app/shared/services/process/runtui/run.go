package runtui

import "solopg/app/shared/services/process"

type runner struct {
	process.Process
	cache
}

type cache struct {
}

func Process() *runner {
	return &runner{

	}
}
func (i *runner) Run() {}

func (i *runner) GetResult() {}