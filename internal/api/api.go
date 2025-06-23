package api

import (
	"net/http"
	"url-shortener/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(store store.Store) http.Handler {
	routes := chi.NewMux()
	routes.Use(middleware.Recoverer)
	routes.Use(middleware.RequestID)
	routes.Use(middleware.Logger)

	routes.Route("/api", func(r chi.Router) {
		r.Post("/shorten", handleShortenURL(store))
		r.Get("/{code}", handleGetShortenedURL(store))
	})

	return routes
}

type apiResponse struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}
