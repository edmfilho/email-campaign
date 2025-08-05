package campaign

import (
	"campaign-project/internal/contract"
	internalerrors "campaign-project/internal/internalErrors"
	"errors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaignDTO) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Create(campaign)

	if err != nil {
		return "", internalerrors.ErrInternalServerError
	}

	return campaign.ID, nil
}

func (s *Service) FindAll() ([]contract.CampaignResponse, error) {
	campaigns, err := s.Repository.Get()

	if err != nil {
		return nil, internalerrors.ErrInternalServerError
	}

	var response []contract.CampaignResponse

	for _, v := range campaigns {
		response = append(response, contract.CampaignResponse{
			ID:      v.ID,
			Name:    v.Name,
			Content: v.Content,
			Status:  v.Status,
		})
	}

	return response, nil
}

func (s *Service) FindBy(id string) (*contract.CampaignResponse, error) {
	campaign, err := s.Repository.GetByID(id)

	if err != nil {
		return nil, internalerrors.ProcessError(err)
	}

	return &contract.CampaignResponse{
		ID:             campaign.ID,
		Status:         campaign.Status,
		Name:           campaign.Name,
		Content:        campaign.Content,
		AmountContacts: len(campaign.Contacts),
	}, nil
}

func (s *Service) Delete(id string) error {
	campaign, err := s.Repository.GetByID(id)

	if err != nil {
		return internalerrors.ProcessError(err)
	}

	if campaign.Status != Pending {
		return errors.New("invalid campaign status")
	}

	campaign.Delete()

	err = s.Repository.Delete(campaign)

	if err != nil {
		return internalerrors.ErrInternalServerError
	}

	return nil
}
