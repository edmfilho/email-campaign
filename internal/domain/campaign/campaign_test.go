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
	campaign := NewCampaign(name, content, contacts)

	assertions.Equal(campaign.ID, "1")
	assertions.Equal(campaign.Name, name)
	assertions.Equal(campaign.Content, content)
	assertions.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assertions := assert.New(t)

	campaign := NewCampaign(name, content, contacts)

	assertions.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assertions := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, contacts)

	assertions.NotNil(campaign.CreatedOn)
	assertions.Greater(campaign.CreatedOn, now)
}
