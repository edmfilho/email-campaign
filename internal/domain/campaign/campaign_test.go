package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "TESTE"
	contacts = []string{"email1@e.com", "email2@e.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assertions := assert.New(t)
	campaign, _ := NewCampaign(name, content, contacts)

	assertions.Equal(campaign.ID, "1")
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

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assertions.Equal("name is required", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assertions.Equal("content is required", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assertions.Equal("contacts is required", err.Error())
}
