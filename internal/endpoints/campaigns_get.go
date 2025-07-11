package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (any, int, error) {
	campaigns, err := h.CampaignService.FindAll()

	return campaigns, 200, err
}
