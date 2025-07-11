package campaign

import (
	"campaign-project/internal/contract"
	internalerrors "campaign-project/internal/internalErrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Get() []Campaign {
	return nil
}

func (r *repositoryMock) GetByID(id string) (*Campaign, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Campaign), nil
}

var (
	newCampaign = contract.NewCampaignDTO{
		Name:    "Teste Service",
		Content: "CONTEUDO",
		Emails:  []string{"teste@gmail.com", "testestestes@gmail.com"},
	}

	repository = new(repositoryMock)

	service = Service{}
)

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaignDTO{})

	assert.False(errors.Is(err, internalerrors.InternalServerError))
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)

	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service.Repository = repositoryMock

	service.Create(newCampaign)

	repository.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(err, internalerrors.InternalServerError))
}

func Test_GetById_ReturnCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("GetByID", mock.Anything).Return(campaign, nil)

	repositoryMock.On("GetByID", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).Return(campaign, nil)

	service.Repository = repositoryMock

	campaignReturned, _ := service.FindBy(campaign.ID)

	assert.Equal(campaign.ID, campaignReturned.ID)
	assert.Equal(campaign.Name, campaignReturned.Name)
	assert.Equal(campaign.Content, campaignReturned.Content)
	assert.Equal(campaign.Status, campaignReturned.Status)
}

func Test_GetById_ReturnError(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("GetByID", mock.Anything).Return(nil, errors.New("error to find campaign"))
	service.Repository = repositoryMock

	_, err := service.FindBy(campaign.ID)

	assert.Equal(internalerrors.InternalServerError.Error(), err.Error())
}
