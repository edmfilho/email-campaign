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

func (s *Service) FindBy(id string) (*contract.CampaignResponse, error) {
	campaign, err := s.Repository.GetByID(id)

	if err != nil {
		return nil, internalerrors.InternalServerError
	}

	return &contract.CampaignResponse{
		ID:      campaign.ID,
		Status:  campaign.Status,
		Name:    campaign.Name,
		Content: campaign.Content,
	}, nil
}
