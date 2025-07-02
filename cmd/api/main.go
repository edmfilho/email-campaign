package main

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)


	println("Iniciando Server...")
	http.ListenAndServe(":3000", r)
}
