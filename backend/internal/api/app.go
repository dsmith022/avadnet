package api

import (
	"avadnet/backend/internal/config"
	"log/slog"
)

type application struct {
	config *config.Config
	logger *slog.Logger
}

type envelope map[string]any
