package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (any, int, error) {
	return h.CampaignService.Repository.Get(), 200, nil
}
