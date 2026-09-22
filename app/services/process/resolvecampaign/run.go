package resolvecampaign

import (
	"solopg/app/domain/campaign"
	"solopg/app/domain/card/attributes/rarity"
	"solopg/app/domain/card/attributes/stats"
	"solopg/app/domain/card/characters"
	"solopg/app/domain/card/characters/wallet"
	"solopg/app/domain/card/objects"
	"solopg/app/domain/card/objects/equipment"
	"solopg/app/services/t"
	"solopg/app/tui"
)

func ResolveCampaignFromContext(ctx *tui.Context) (*campaign.Campaign, error) {
	selectedCampaign := ctx.SelectedSave

	if selectedCampaign == nil {
		newCampaign, err := buildCampaignFromContext(ctx)
		if err != nil {
			return nil, err
		}

		selectedCampaign = newCampaign
	}

	return selectedCampaign, nil
}

func buildCampaignFromContext(ctx *tui.Context) (*campaign.Campaign, error) {
	isValidArgs := ctx.SelectedRace != nil && ctx.SelectedClass != nil && ctx.SelectedLocation != nil
	if !isValidArgs {
		return nil, t.NewError("error.incomplete_character_creation")
	}

	player, err := generateCharacter(ctx)
	if err != nil {
		return nil, err
	}

	newCampaign := campaign.New(campaign.Template{
		Player:          player,
		CurrentLocation: ctx.SelectedLocation,
	})

	return newCampaign, nil
}

func generateCharacter(ctx *tui.Context) (*characters.Character, error) {
	player, err := characters.New(characters.Template{
		Name:      ctx.SelectedName,
		Race:      ctx.SelectedRace.GetName(),
		Class:     ctx.SelectedClass.GetName(),
		Rarity:    rarity.Default(),
		Stats:     getStatsFromContext(ctx),
		Equipment: getEquipementFromContext(ctx),
		Inventory: []objects.Object{},
		Wallet:    wallet.Wallet{},
	})

	if err != nil {
		return nil, err
	}

	return player, nil
}

func getEquipementFromContext(ctx *tui.Context) equipment.Equipment {
	name := ctx.SelectedClass.GetEquipementName()
	set := equipment.FindEquipementByName(name)

	return equipment.NewSet(set)
}

func getStatsFromContext(ctx *tui.Context) stats.Stats {
	modifiers := getModifiersFromContext(ctx)

	baseStats := stats.GetBasic()
	baseStats.ApplyModifiers(modifiers)

	return baseStats
}

func getModifiersFromContext(ctx *tui.Context) []stats.Modifier {
	raceBoost := ctx.SelectedRace.GetBonus()
	classBoost := ctx.SelectedClass.GetBonus()
	build := ctx.SelectedBuild

	modifiers := make([]stats.Modifier, 0, len(raceBoost)+len(classBoost)+len(build))
	modifiers = append(modifiers, raceBoost...)
	modifiers = append(modifiers, classBoost...)
	modifiers = append(modifiers, build...)

	return modifiers
}
