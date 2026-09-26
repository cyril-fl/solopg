package metadatamenu

import (
	"fmt"
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/services/game"
	"solopg/app/services/i18n"
	"solopg/app/utils/transform"
	"strings"
)

// TODO LOW refactor ca surtour la maniere de if else
func GetCharacterInfo(content *strings.Builder, engine *game.Engine) {
	var characterName = transform.Uppercase(i18n.Localize("character"))

	isEngineNil := engine == nil || engine.State == nil

	if isEngineNil || engine.State.Player == nil {
		fmt.Fprintf(content, "%s", characterName)
	} else {
		fmt.Fprintf(content, "%s: %s", characterName, i18n.Localize(engine.State.Player.Name))
	}

	content.WriteString(characterName)
	content.WriteString("\n")
}

func GetLocationInfo(content *strings.Builder, engine *game.Engine) {
	var locationName = transform.Uppercase(i18n.Localize("location"))

	isEngineNil := engine == nil || engine.State == nil

	if isEngineNil || engine.State.CurrentLocation == nil {
		fmt.Fprintf(content, "%s %s", locationName, i18n.Localize("unknown_place"))
	} else {
		fmt.Fprintf(content, "%s %s", locationName, i18n.Localize(engine.State.CurrentLocation.Name))
	}
	content.WriteString("\n")
}

func GetStatInfo(content *strings.Builder, engine *game.Engine) {
	content.WriteString(transform.Uppercase(i18n.Localize("stats")))
	content.WriteString("\n")

	if engine == nil || engine.State == nil || engine.State.Player == nil {
		content.WriteString(i18n.Localize("no_stats"))
	} else {
		for _, stat := range stats.List() {
			key := "stat." + string(stat)
			label := i18n.Localize(key)
			fmt.Fprintf(content, "%-10s %d\n", label, engine.State.Player.Stats[stat])
		}
	}
}
