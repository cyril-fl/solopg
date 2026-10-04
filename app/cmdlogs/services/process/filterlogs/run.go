package filterlogs

import (
	"fmt"
	"slices"
	"solopg/app/shared/services/logs"
	"solopg/app/shared/services/process"

	"github.com/spf13/pflag"
)

type filter struct {
	process.Process
	cache

	options *pflag.FlagSet
	list    []logs.Entry
}

type cache struct {
	tail   int
	filter string

	err error
}

func Process(get func() *pflag.FlagSet, logs []logs.Entry) *filter {
	return &filter{
		options: get(),
		list:    logs,
	}
}

func (p *filter) Run() {
	p.setFlags()

	p.assertFilter()

	p.filterLogs()
	p.tailLogs()
}

func (p *filter) GetResult() []logs.Entry {
	return p.list
}

// Methods
func (p *filter) setFlags() {
	p.tail, p.err = p.options.GetInt("tail")
	p.filter, p.err = p.options.GetString("filter")

	p.checkErr()
}

func (p *filter) assertFilter() {
	if p.HasErr() {
		return
	}

	if p.filter == "" {
		return
	}

	if !slices.Contains(logs.Kinds, logs.Kind(p.filter)) {
		// i18N
		p.err = fmt.Errorf("invalid log level: %s", p.filter)
	}

	p.checkErr()
}

func (p *filter) filterLogs() {
	if p.HasErr() {
		return
	}

	if p.filter == "" {
		return
	}

	newlist := []logs.Entry{}
	for _, log := range p.list {
		if log.Type != logs.Kind(p.filter) {
			continue
		}

		newlist = append(newlist, log)
	}

	p.list = newlist
}

func (p *filter) tailLogs() {
	if p.HasErr() {
		return
	}

	if p.tail <= 0 {
		return
	}

	if p.tail > len(p.list) {
		return
	}

	if p.tail < len(p.list) {
		p.list = p.list[len(p.list)-p.tail:]
	}
}

func (p *filter) checkErr() {
	if p.err != nil {
		p.SetErr(p.err)
	}
}
