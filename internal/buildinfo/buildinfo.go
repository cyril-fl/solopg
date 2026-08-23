package buildinfo

import (
	"fmt"
	"runtime/debug"
)

var (
	AppName = "solopg"
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

func init() {
	fillFromBuildInfo()
}

func fillFromBuildInfo() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	commitFromBuildInfo := false

	isDev :=Version == "dev" && info.Main.Version != "" && info.Main.Version != "(devel)" 

	if isDev {
		Version = info.Main.Version
	}

	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.revision":
			if Commit == "none" && s.Value != "" {
				if len(s.Value) > 7 {
					Commit = s.Value[:7]
				} else {
					Commit = s.Value
				}
				commitFromBuildInfo = true
			}
		case "vcs.time":
			if Date == "unknown" && s.Value != "" {
				Date = s.Value
			}
		case "vcs.modified":
			if s.Value == "true" && commitFromBuildInfo {
				Commit += "-dirty"
			}
		}
	}
}

func String() string {
	return fmt.Sprintf("%s %s (%s, %s)", AppName, Version, Commit, Date)
}
