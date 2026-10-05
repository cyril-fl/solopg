package filterlogs

import (
	"solopg/app/shared/services/factory/millfilterlog"
	"solopg/app/shared/services/logs"
	"solopg/app/shared/services/process"

	"github.com/spf13/pflag"
)

type filter struct {
	process.Process
	cache

	flags *pflag.FlagSet
	list  []logs.Log
}

type cache struct {
	mill *millfilterlog.Mill
}

func Process(get func() *pflag.FlagSet, logs []logs.Log) *filter {
	return &filter{
		flags: get(),
		list:  logs,
	}
}

func (p *filter) Run() {
	p.makeMill()
}

func (p *filter) GetResult() []logs.Log {
	return p.mill.
		Tail().
		Filter().
		GetList()
}

// Methods
func (p *filter) makeMill() {
	mill := millfilterlog.New(millfilterlog.Template{
		LogList: p.list,
		Flags:   p.flags,
	})

	p.mill = mill
}
