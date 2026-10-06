package generatebuildinfo

import (
	"os/exec"
	"solopg/app/cmdversion/types/buildinfo"
	"solopg/app/shared/services/process"
	"solopg/config"
	"strings"
)

var (
	Version = ""
	Commit  = ""
	Date    = ""
)
// TODO adapter cette ,ottation au autre process
type Generator = p

type p struct {
	process.Process
	cache
}

type cache struct {
	version string
	commit  string
	date    string
}

type shellcmd = string

const (
	GET_COMMIT  = "git rev-parse --short HEAD 2>/dev/null || echo none"
	GET_DATE    = "date -u +%Y-%m-%dT%H:%M:%SZ"
	GET_VERSION = "git describe --tags --always --dirty 2>/dev/null || echo dev"
)

func Process() *p {
	return &p{}
}
func (p *p) Run() {
	executeByEnv(executablenv{
		devfunc:  p.setDevEnv,
		prodfunc: p.setProdEnv,
	})
}

func (p *p) setDevEnv() {
	p.setVersion()
	p.setCommit()
	p.setDate()
}

func (p *p) setCommit() {
	p.cache.commit = p.executeCmd(GET_COMMIT)
}

func (p *p) setDate() {
	p.cache.date = p.executeCmd(GET_DATE)
}

func (p *p) setVersion() {
	p.cache.version = p.executeCmd(GET_VERSION)
}

func (p *p) executeCmd(cmd shellcmd) string {
	c := exec.Command("sh", "-c", cmd)
	out, err := c.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func (p *p) setProdEnv() {
	p.cache.version = Version
	p.cache.commit = Commit
	p.cache.date = Date
}

func (p *p) GetResult() buildinfo.BuildInfo {
	infos := buildinfo.New(buildinfo.InfoTemplate{
		Version: p.cache.version,
		Commit:  p.cache.commit,
		Date:    p.cache.date,
	})

	return infos
}

// helpers
type executablenv struct {
	devfunc  func()
	prodfunc func()
}

func executeByEnv(e executablenv) {
	if config.Current.IsDev() {
		e.devfunc()
	} else {
		e.prodfunc()
	}
}
