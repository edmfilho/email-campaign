package campaign_test

import (
	"campaign-project/internal/contract"
	"campaign-project/internal/domain/campaign"
	internalmock "campaign-project/internal/internal-mock"
	internalerrors "campaign-project/internal/internalErrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	newCampaign = contract.NewCampaignDTO{
		Name:    "Teste Service",
		Content: "CONTEUDO",
		Emails:  []string{"teste@gmail.com", "testestestes@gmail.com"},
	}

	repository = new(internalmock.CampaignRepositoryMock)

	service = campaign.Service{}
)

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaignDTO{})

	assert.False(errors.Is(err, internalerrors.ErrInternalServerError))
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(nil)

	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Create", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
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

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(err, internalerrors.ErrInternalServerError))
}

func Test_GetById_ReturnCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
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

	campaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetByID", mock.Anything).Return(nil, errors.New("error to find campaign"))
	service.Repository = repositoryMock

	_, err := service.FindBy(campaign.ID)

	assert.Equal(internalerrors.ErrInternalServerError.Error(), err.Error())
}

func Test_Delete_ReturnNotFoundWhenCampaignDoesNotExist(t *testing.T) {
	assert := assert.New(t)

	IdInvalid := "invalid"

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetByID", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	service.Repository = repositoryMock

	err := service.Delete(IdInvalid)

	assert.Equal(err.Error(), gorm.ErrRecordNotFound.Error())
}

func Test_Delete_ReturnStatusInvalid_When_Campaign_Status_Is_Not_Peding(t *testing.T) {
	assert := assert.New(t)

	campaign := &campaign.Campaign{ID: "1", Status: campaign.Started}

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetByID", mock.Anything).Return(campaign, nil)

	service.Repository = repositoryMock

	err := service.Delete(campaign.ID)

	assert.Equal("invalid campaign status", err.Error())
}

func Test_Delete_ReturnInternalError_When_Fail(t *testing.T) {
	assert := assert.New(t)

	campaignFound := &campaign.Campaign{ID: "1", Status: campaign.Pending}

	repositoryMock := new(internalmock.CampaignRepositoryMock)

	repositoryMock.On("GetByID", mock.Anything).Return(campaignFound, nil)
	repositoryMock.On("Delete", mock.MatchedBy(
		func(campaign *campaign.Campaign) bool {
			return campaignFound == campaign
		}),
	).Return(errors.New("error to delete campaign"))

	service.Repository = repositoryMock

	err := service.Delete(campaignFound.ID)

	assert.Equal(internalerrors.ErrInternalServerError.Error(), err.Error())
}

func Test_Delete_ReturnNil_When_Delete_Has_Success(t *testing.T) {
	assert := assert.New(t)

	campaignFound, _ := campaign.NewCampaign("teste", "body 1", []string{"teste@email.com"})

	repositoryMock := new(internalmock.CampaignRepositoryMock)

	repositoryMock.On("GetByID", mock.Anything).Return(campaignFound, nil)
	repositoryMock.On("Delete", mock.MatchedBy(
		func(campaign *campaign.Campaign) bool {
			return campaignFound == campaign
		}),
	).Return(errors.New("error to delete campaign"))

	service.Repository = repositoryMock

	err := service.Delete(campaignFound.ID)

	assert.Equal(internalerrors.ErrInternalServerError.Error(), err.Error())
}
