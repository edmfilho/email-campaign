package endpoints

import (
	internalerrors "campaign-project/internal/internalErrors"
	"errors"
	"net/http"
	"slices"

	"github.com/go-chi/render"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (any, int, error)

func HandlerError(endpointfunc EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := endpointfunc(w, r)

		if err != nil {
			if errors.Is(err, internalerrors.ErrInternalServerError) {
				render.Status(r, 500)
			} else {
				render.Status(r, 400)
			}

			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		if !slices.Contains([]int{200, 201, 202}, status) {
			render.Status(r, status)

			render.JSON(w, r, map[string]string{"error": "resource not found"})
			return
		}

		render.Status(r, status)
		if obj != nil {
			render.JSON(w, r, obj)
		}
	})
}
