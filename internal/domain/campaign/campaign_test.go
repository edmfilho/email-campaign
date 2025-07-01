package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "TESTE DE CONTEUDO"
	contacts = []string{"emailteste@email.com", "email2@email.com"}

	fake = faker.New()
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assertions := assert.New(t)

	campaign, err := NewCampaign(name, content, contacts)

	assertions.Nil(err)
	assertions.NotNil(campaign.ID)
	assertions.Equal(campaign.Name, name)
	assertions.Equal(campaign.Content, content)
	assertions.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assertions := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assertions.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assertions := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assertions.NotNil(campaign.CreatedOn)
	assertions.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assertions.Equal("name is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assertions.Equal("content is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts)

	assertions.Equal("content is required with max 1024", err.Error())
}

func Test_NewCampaign_MustValidateContactsLen(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assertions.Equal("contacts is required with min 1", err.Error())
}
func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email.com"})

	assertions.Equal("email is invalid", err.Error())
}
