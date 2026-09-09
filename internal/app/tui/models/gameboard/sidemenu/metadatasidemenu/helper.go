package metadatasidemenu

import (
	"fmt"
	"solopg/internal/app/game"
	"solopg/internal/domain/card/effects"
	"solopg/internal/infrastructure/t"
	"strings"
)

func GetCharacterInfo(content *strings.Builder, engine *game.Engine) {
	var characterName = t.Localize("character")

	isEngineNil := engine == nil || engine.State == nil

	if isEngineNil || engine.State.Player == nil {
		characterName += t.Localize("unknown_player")
	} else {
		characterName += engine.State.Player.Name
	}

	content.WriteString(characterName)
	content.WriteString("\n")
}

func GetLocationInfo(content *strings.Builder, engine *game.Engine) {
	var locationName = t.Localize("place")

	isEngineNil := engine == nil || engine.State == nil

	if isEngineNil || engine.State.CurrentLocation == nil {
		locationName += t.Localize("unknown_place")
	} else {
		locationName += engine.State.CurrentLocation.Name
	}

	content.WriteString(locationName)
	content.WriteString("\n")
}

func GetStatInfo(content *strings.Builder, engine *game.Engine) {
	content.WriteString(t.Localize("stats_upper") + "\n")

	if engine == nil || engine.State == nil || engine.State.Player == nil {
		content.WriteString(t.Localize("no_stats"))
	} else {
		for _, stat := range effects.ListStats() {
			key := "stat." + string(stat)
			label := t.Localize(key)
			fmt.Fprintf(content, "%-10s %d\n", label, engine.State.Player.Stats[stat])
		}
	}
}
