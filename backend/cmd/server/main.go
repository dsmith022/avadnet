package main

import (
	"avadnet/backend/internal/api"
	"avadnet/backend/internal/config"
	"errors"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := config.Load()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	srv := api.NewServer(cfg, logger)

	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
