package api

import (
	"errors"
	"log/slog"
	"net/http"
	"url-shortener/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

type getShortenedURLResponse struct {
	FullURL string `json:"full_url"`
}

func handleGetShortenedURL(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		if code == "" {
			sendJSON(w, apiResponse{Error: "code is required"}, http.StatusBadRequest)
			return
		}

		url, err := store.GetFullURL(r.Context(), code)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				sendJSON(w, apiResponse{Error: "code not found"}, http.StatusBadRequest)
				return
			}

			slog.Error("failed to get code", "error", err)
			sendJSON(w, apiResponse{Error: "something went wrong"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, apiResponse{Data: getShortenedURLResponse{FullURL: url}}, http.StatusOK)
	}
}
