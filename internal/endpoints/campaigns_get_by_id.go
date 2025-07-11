package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignGetByID(w http.ResponseWriter, r *http.Request) (any, int, error) {
	id := r.URL.Query().Get("id")

	campaign, err := h.CampaignService.FindBy(id)

	return campaign, 200, err
}
