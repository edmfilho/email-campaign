package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignGetByID(w http.ResponseWriter, r *http.Request) (any, int, error) {
	id := chi.URLParam(r, "id")

	campaign, err := h.CampaignService.FindBy(id)

	return campaign, 200, err
}
