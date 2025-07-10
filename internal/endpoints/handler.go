package endpoints

import "campaign-project/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}