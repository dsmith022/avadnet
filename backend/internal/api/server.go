package api

import (
	"avadnet/backend/internal/config"
	"log/slog"
	"net/http"
	"time"
)

func NewServer(cfg *config.Config, logger *slog.Logger) *http.Server {
	app := &application{
		config: cfg,
		logger: logger,
	}

	return &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           app.routes(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
}
