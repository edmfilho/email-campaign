package database

import (
	"campaign-project/internal/domain/campaign"
	"fmt"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.campaigns = append(c.campaigns, *campaign)
	return nil
}

func (c *CampaignRepository) Get() []campaign.Campaign {
	return c.campaigns
}

func (c *CampaignRepository) GetByID(id string) (campaign.Campaign, error) {
	for _, v := range c.campaigns {
		if id == v.ID {
			return v, nil
		}
	}

	return campaign.Campaign{}, fmt.Errorf("campaign with id %s not founded", id)
}
