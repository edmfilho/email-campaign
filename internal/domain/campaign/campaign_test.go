package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCampaign(t *testing.T) {
	assertions := assert.New(t)
	name := "Campaign X"
	content := "TESTE"
	contacts := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contacts)

	assertions.Equal(campaign.ID, "1")
	assertions.Equal(campaign.Name, name)
	assertions.Equal(campaign.Content, content)
	assertions.Equal(len(campaign.Contacts), len(contacts))
}
