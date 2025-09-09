package main

import (
	"FilmFinder/api"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		os.Exit(1)
	}

	slog.Info("all systems offline")
}

func run() error {
	db := make(map[string]string)
	handler := api.NewHandler(db)

	s := http.Server{
		ReadTimeout: 10 * time.Second,
		Addr:        ":8080",
		Handler:     handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
