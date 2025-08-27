package main

import (
	"campaign-project/internal/domain/campaign"
	"campaign-project/internal/endpoints"
	"campaign-project/internal/infra/database"
	"fmt"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.NewDB()
	campaignService := campaign.Service{
		Repository: &database.CampaignRepository{Db: db},
	}

	r.Use(middleware.Logger)
	handler := endpoints.Handler{
		CampaignService: campaignService,
	}

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/campaigns", func(r chi.Router) {
		r.Use(endpoints.Auth)

		r.Post("/", endpoints.HandlerError(handler.CampaignPost))
		r.Get("/", endpoints.HandlerError(handler.CampaignGet))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetByID))
		r.Delete("/{id}", endpoints.HandlerError(handler.CampaignDelete))
	})

	fmt.Println("Listening request...")
	http.ListenAndServe(":3000", r)
}
