package campaign

import (
	"campaign-project/internal/contract"
	internalerrors "campaign-project/internal/internalErrors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaignDTO) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", internalerrors.InternalServerError
	}

	return campaign.ID, nil
}
