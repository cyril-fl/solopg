package process

import (
	"solopg/internal/app/tui/forgeui"
	"solopg/internal/app/tui/portalui"
	"solopg/internal/domain/campaign"
)

func ResolveCampaign(selected *campaign.Campaign) (*campaign.Campaign, error) {
	if selected != nil {
		return selected, nil
	}

	var new *campaign.Campaign

	forgeUi := forgeui.New()
	newCharacter, err := forgeUi.CreateCharacter()

	if err != nil {
		return nil, err
	}

	portalUi := portalui.NewUi()
	location, err := portalUi.SelectLocation()

	if err != nil {
		return nil, err
	}

	new = campaign.New(campaign.Template{
		Player:          newCharacter,
		CurrentLocation: location,
	})

	return new, nil
}
