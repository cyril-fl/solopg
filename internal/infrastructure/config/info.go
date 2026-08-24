package config

import (
	"fmt"
	"runtime/debug"
)

var (
	Version             = "dev"
	Commit              = "none"
	Date                = "unknown"
	commitFromBuildInfo = false
)

func init() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	setVersion(*info)

	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.revision":
			setCommit(s)
		case "vcs.time":
			setTime(s)
		case "vcs.modified":
			setModified(s)
		}
	}
}

func setVersion(i debug.BuildInfo) {
	if Version == "dev" && i.Main.Version != "" && i.Main.Version != "(devel)" {
		Version = i.Main.Version
	}
}

func setCommit(s debug.BuildSetting) {
	if Commit == "none" && s.Value != "" {
		if len(s.Value) > 7 {
			Commit = s.Value[:7]
		} else {
			Commit = s.Value
		}
		commitFromBuildInfo = true
	}
}

func setTime(s debug.BuildSetting) {
	if Date == "unknown" && s.Value != "" {
		Date = s.Value
	}
}

func setModified(s debug.BuildSetting) {
	if s.Value == "true" && commitFromBuildInfo {
		Commit += "-dirty"
	}
}

func String() string {
	return fmt.Sprintf("%s %s (%s, %s)", Current.Name, Version, Commit, Date)
}
