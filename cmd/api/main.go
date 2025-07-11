package main

import (
	"campaign-project/internal/domain/campaign"
	"campaign-project/internal/endpoints"
	"campaign-project/internal/infra/database"

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

	handler := endpoints.Handler{
		CampaignService: campaignService,
	}

	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaigns", endpoints.HandlerError(handler.CampaignGet))
	r.Get("/campaign/{id}", endpoints.HandlerError(handler.CampaignGetByID))

	println("Iniciando Server...")
	http.ListenAndServe(":3000", r)
}
