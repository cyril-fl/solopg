package filterlogs

import (
	"solopg/app/shared/services/factory/millfilterlog"
	"solopg/app/shared/services/logs"
	"solopg/app/shared/services/process"

	"github.com/spf13/pflag"
)

type run struct {
	process.Process
	cache

	flags *pflag.FlagSet
	list  []logs.Log
}

type cache struct {
	mill *millfilterlog.Mill
}

func Process(get func() *pflag.FlagSet, logs []logs.Log) *run {
	return &run{
		flags: get(),
		list:  logs,
	}
}

func (p *run) Run() {
	p.makeMill()
}

func (p *run) GetResult() []logs.Log {
	return p.mill.
		Tail().
		Filter().
		GetList()
}

// Methods
func (p *run) makeMill() {
	mill := millfilterlog.New(millfilterlog.Template{
		LogList: p.list,
		Flags:   p.flags,
	})

	p.mill = mill
}
