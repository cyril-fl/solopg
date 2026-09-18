package process

import (
	"solopg/internal/app/tui"
	"solopg/internal/domain/campaign"
	"solopg/internal/domain/card/attributes/rarity"
	"solopg/internal/domain/card/characters"
	"solopg/internal/domain/card/effects"
	"solopg/internal/domain/card/objects"
	"solopg/internal/infrastructure/t"
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
		Equipment: getArmorSetFromContext(ctx),
		Inventory: []objects.Object{},
		Wallet:    characters.Wallet{},
	})

	if err != nil {
		return nil, err
	}

	return player, nil
}

func getArmorSetFromContext(ctx *tui.Context) characters.ArmorSet {
	setName := ctx.SelectedClass.GetArmorSet()

	set := characters.FindArmorSetByName(setName)

	armorSet := characters.NewArmorSet(set)

	return armorSet
}

func getStatsFromContext(ctx *tui.Context) effects.Stats {
	modifiers := getModifiersFromContext(ctx)

	baseStats := effects.BaseStats()
	baseStats.ApplyModifiers(modifiers)

	return baseStats
}

func getModifiersFromContext(ctx *tui.Context) []effects.Modifier {
	raceBoost := ctx.SelectedRace.GetBonus()
	classBoost := ctx.SelectedClass.GetBonus()
	build := ctx.SelectedBuild

	modifiers := make([]effects.Modifier, 0, len(raceBoost)+len(classBoost)+len(build))
	modifiers = append(modifiers, raceBoost...)
	modifiers = append(modifiers, classBoost...)
	modifiers = append(modifiers, build...)

	return modifiers
}
