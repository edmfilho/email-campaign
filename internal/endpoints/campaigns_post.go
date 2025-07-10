package endpoints

import (
	"campaign-project/internal/contract"
	internalerrors "campaign-project/internal/internalErrors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (any, int, error) {
	var request contract.NewCampaignDTO
	render.DecodeJSON(r.Body, &request)

	id, err := h.CampaignService.Create(request)

	if err != nil {
		if errors.Is(err, internalerrors.InternalServerError) {
			return nil, 500, internalerrors.InternalServerError
		} else {
			return nil, 400, err
		}
	}

	return map[string]string{"id": id}, 201, nil
}
