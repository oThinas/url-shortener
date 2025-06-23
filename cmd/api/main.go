package main

import (
	"log/slog"
	"net/http"
	"time"
	"url-shortener/internal/api"
	"url-shortener/internal/store"

	"github.com/redis/go-redis/v9"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}

	slog.Info("all systems offline")
}

func run() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	store := store.NewStore(rdb)
	hadler := api.NewHandler(store)

	server := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  1 * time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      hadler,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
