package process

import (
	"fmt"
	"solopg/internal/app/tui"
	"solopg/internal/domain/campaign"
	"solopg/internal/domain/card/attributes"
	"solopg/internal/domain/card/characters"
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
		return nil, fmt.Errorf("incomplete character creation context")
	}

	player, err := characters.New(characters.Template{
		Name:   ctx.SelectedName,
		Race:   ctx.SelectedRace.GetName(),
		Class:  ctx.SelectedClass.GetName(),
		Rarity: attributes.A,
	})
	if err != nil {
		return nil, err
	}

	newCampaign := campaign.New(campaign.Template{
		Player:          player,
		CurrentLocation: ctx.SelectedLocation,
	})

	return newCampaign, nil
}
